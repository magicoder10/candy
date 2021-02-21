package assets

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"path/filepath"
	"strings"

	"candy/audio"
)

var imageExtensions = map[string]struct{}{"jpg": {}, "png": {}}
var audioExtensions = map[string]struct{}{"mp3": {}, "wav": {}}

type Assets struct {
  imageMap map[string]image.Image
	audioMap map[string]audio.Audio
}

func (a Assets) GetImage(imageName string) image.Image {
  return a.imageMap[convertPath(imageName)]
}

func (a Assets) GetAudio(audioName string) audio.Audio {
  return a.audioMap[convertPath(audioName)]
}

func convertPath(path string) string {
  parts := strings.Split(path, "/")
  return filepath.Join(parts...)
}

func LoadAssets(assetRootDir string) (Assets, error) {
	imageAssets, err := loadImages(assetRootDir)
	if err != nil {
		return Assets{}, err
	}
	audioAssets, err := loadAudios(assetRootDir)
	return Assets{
		imageMap: imageAssets,
		audioMap: audioAssets,
	}, err
}

func loadAudios(assetRootDir string) (map[string]audio.Audio, error) {
	audioAssets := make(map[string]audio.Audio)
	err := lisFiles(assetRootDir, func(path string, ext string, rel string) error {
		if _, ok := audioExtensions[ext]; !ok {
			return nil
		}
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		bp, err := audio.NewAudio(file, ext)
		if err != nil {
			return err
		}
		audioAssets[rel] = bp
		return nil
	})
	return audioAssets, err
}

func loadImages(assetRootDir string) (map[string]image.Image, error) {
	imageAssets := make(map[string]image.Image)
	err := lisFiles(assetRootDir, func(path string, ext string, rel string) error {
		if _, ok := imageExtensions[ext]; !ok {
			return nil
		}
		file, err := os.Open(path)
		defer file.Close()
		if err != nil {
			return err
		}
		img, _, err := image.Decode(file)
		if err != nil {
			return err
		}
		imageAssets[rel] = img
		return nil
	})
	return imageAssets, err
}

func lisFiles(rootDir string, processFile func(path string, ext string, rel string) error) error {
	return filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		parts := strings.Split(info.Name(), ".")
		if len(parts) == 0 {
			return nil
		}
		ext := parts[len(parts)-1]
		rel, err := filepath.Rel(rootDir, path)
		return processFile(path, ext, rel)
	})
}
