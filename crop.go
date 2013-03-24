package main

import (
	"github.com/hawx/img/crop"
	"github.com/hawx/img/utils"
	"github.com/hawx/hadfield"
)

var cmdCrop = &hadfield.Command{
	Usage: "crop [options]",
	Short: "crop an image",
Long: `
  Crop takes an image fron STDIN, and returns a cropped version to STDOUT. By
  default it will use the largest size possible; that is, if the image is wider
  than tall it will use the height, and vica versa.

    --square               # Crop to a square (default)
    --circle               # Crop to a circle

    --size <pixels>        # Size to crop to (default: largest possible)

    --centre               # Centre the image
    --top                  # Centre the image to the top of the frame
    --top-right            # Centre the image to the top-right of the frame
    --right                # Centre the image to the right of the frame
    --bottom-right         # Centre the image to the bottom-right of the frame
    --bottom               # Centre the image to the bottom of the frame
    --bottom-left          # Centre the image to the bottom-left of the frame
    --left                 # Centre the image to the left of the frame
    --top-left             # Centre the image to the top-left of the frame
`,
}

var cropSquare, cropCircle bool
var cropSize int
var cropCentre, cropTop, cropTopRight, cropRight, cropBottomRight, cropBottom,
    cropBottomLeft, cropLeft, cropTopLeft bool

func init() {
	cmdCrop.Run = runCrop

	cmdCrop.Flag.BoolVar(&cropSquare,      "square",       false, "")
	cmdCrop.Flag.BoolVar(&cropCircle,      "circle",       false, "")

	cmdCrop.Flag.IntVar(&cropSize,         "size",         -1,    "")

	cmdCrop.Flag.BoolVar(&cropCentre,      "centre",       false, "")
	cmdCrop.Flag.BoolVar(&cropTop,         "top",          false, "")
	cmdCrop.Flag.BoolVar(&cropTopRight,    "top-right",    false, "")
	cmdCrop.Flag.BoolVar(&cropRight,       "right",        false, "")
	cmdCrop.Flag.BoolVar(&cropBottomRight, "bottom-right", false, "")
	cmdCrop.Flag.BoolVar(&cropBottom,      "bottom",       false, "")
	cmdCrop.Flag.BoolVar(&cropBottomLeft,  "bottom-left",  false, "")
	cmdCrop.Flag.BoolVar(&cropLeft,        "left",         false, "")
	cmdCrop.Flag.BoolVar(&cropTopLeft,     "top-left",     false, "")
}

func runCrop(cmd *hadfield.Command, args []string) {
	i := utils.ReadStdin()

	direction := utils.Centre

	if      cropTop           { direction = utils.Top
	} else if cropTopRight    { direction = utils.TopRight
	} else if cropRight       { direction = utils.Right
	} else if cropBottomRight { direction = utils.BottomRight
	} else if cropBottom      { direction = utils.Bottom
	} else if cropBottomLeft  { direction = utils.BottomLeft
	} else if cropLeft        { direction = utils.Left
	} else if cropTopLeft     { direction = utils.TopLeft }

	if cropCircle {
		i = crop.Circle(i, cropSize, direction)
	} else {
		i = crop.Square(i, cropSize, direction)
	}

	utils.WriteStdout(i)
}