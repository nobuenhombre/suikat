package imgs

import (
	"image"
	"os"
)

func GetDimension(imagePath string) (width int, height int, err error) {
	file, err := os.Open(imagePath)
	if err != nil {
		return 0, 0, err
	}

	image, _, err := image.DecodeConfig(file)
	if err != nil {
		return 0, 0, err
	}

	return image.Width, image.Height, nil
}
