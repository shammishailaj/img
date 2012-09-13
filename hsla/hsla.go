// Package hsla implements the HSLA colour model.
package hsla

import (
	"../utils"
	"image"
	"image/color"
	"math"
)

// HSLA represents a colour with hue, saturation, lightness and alpha channels.
// All values are represented as floating point numbers with hue taking a value
// between 0 and 360, and the rest between 0 and 1.
type HSLA struct {
	H, S, L, A float64
}

func (col HSLA) RGBA() (red, green, blue, alpha uint32) {
	h := col.H; s := col.S; l := col.L; a := col.A

	h = math.Mod(h, 360)

	c := (1.0 - math.Abs(2.0 * l - 1.0)) * s

	hdash := h / 60.0
	x := c * (1 - math.Abs(math.Mod(hdash,2) - 1))

	var r, g, b float64
	if h == 0 {
		r = 0; g = 0; b = 0
	} else if hdash < 1 {
		r = c; g = x; b = 0
	} else if hdash < 2 {
		r = x; g = c; b = 0
	} else if hdash < 3 {
		r = 0; g = c; b = x
	} else if hdash < 4 {
		r = 0; g = x; b = c
	} else if hdash < 5 {
		r = x; g = 0; b = c
	} else if hdash < 6 {
		r = c; g = 0; b = x
	}

	m := l - 0.5 * c

	red   = uint32(utils.Truncatef((r + m) * a * 255)) << 8
	green = uint32(utils.Truncatef((g + m) * a * 255)) << 8
	blue  = uint32(utils.Truncatef((b + m) * a * 255)) << 8
	alpha = uint32(a * 255) << 8

	return
}

var HSLAModel color.Model = color.ModelFunc(hslaModel)

func hslaModel(c color.Color) color.Color {
	if _, ok := c.(HSLA); ok {
		return c
	}

	r, g, b, a := utils.RatioRGBA(c)

	maxi   := utils.Maxf(r, g, b)
	mini   := utils.Minf(r, g, b)
	chroma := maxi - mini

	// Work out hue
	hdash  := 0.0
	if chroma == 0 {
		hdash = 0
	}else if maxi == r {
		hdash = math.Mod((g - b) / chroma, 6)
	} else if maxi == g {
		hdash = (b - r) / chroma + 2.0
	} else if maxi == b {
		hdash = (r - g) / chroma + 4.0
	}

	hue := hdash * 60

	if chroma == 0 {
		hue = 0
	}

	// Work out lightness
	lightness := 0.5 * (maxi + mini)

	// Work out saturation
	saturation := 0.0
	if chroma != 0 {
		saturation = chroma / (1 - math.Abs(2 * lightness - 1))
	}

	return HSLA{hue, saturation, lightness, a}
}


// Hue shifts the hue of the Image by the amount given.
func Hue(img image.Image, amount float64) image.Image {
	f := func(c color.Color) color.Color {
		h := HSLAModel.Convert(c).(HSLA)
		h.H += amount
		return h
	}

	return utils.EachPixel(img, f)
}

// Saturation adjusts the saturation of the Image by the amount given.
func Saturation(img image.Image, amount float64) image.Image {
	f := func(c color.Color) color.Color {
		h := HSLAModel.Convert(c).(HSLA)
		h.S += amount
		if h.S > 1 { h.S = 1 }
		return h
	}

	return utils.EachPixel(img, f)
}

// Lightness adjusts the lightness of the Image by the amount given.
func Lightness(img image.Image, amount float64) image.Image {
	f := func(c color.Color) color.Color {
		h := HSLAModel.Convert(c).(HSLA)
		h.L += amount
		if h.L > 1 { h.L = 1 }
		return h
	}

	return utils.EachPixel(img, f)
}