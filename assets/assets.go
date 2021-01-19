package assets

import (
	"image"
	_ "image/png"
	"os"
	"path/filepath"
)

var images = []string{
	"map/default.png",
	"sprite_sheet.png",
	"screen/signin.png",
}

type Assets struct {
	imageMap map[string]image.Image
}

func (a Assets) GetImage(imageName string) image.Image {
	return a.imageMap[imageName]
}

func LoadAssets(assetRootDir string) (Assets, error) {
	imageMap := make(map[string]image.Image)

	for _, imgPath := range images {
		file, err := os.Open(filepath.Join(assetRootDir, imgPath))
		if err != nil {
			return Assets{}, err
		}

		img, _, err := image.Decode(file)
		if err != nil {
			_ = file.Close()
			return Assets{}, err
		}
		_ = file.Close()
		imageMap[imgPath] = img
	}

	return Assets{imageMap: imageMap}, nil
}
