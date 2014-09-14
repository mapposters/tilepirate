package main

import (
	"fmt"
	"image"
	"image/png"
	"os"
)

const (
	tileDir      = "output"
	TileFileName = "tile%v_%vx%v.png"
)

func readTile(z, x, y uint) (image.Image, error) {
	file, err := os.Open(fmt.Sprintf(tileDir+"/"+TileFileName, z, x, y))
	if err != nil {
		return nil, err
	}
	return png.Decode(file)
}
