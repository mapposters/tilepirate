package main

import (
	"fmt"
	"image"
	"image/png"
	"os"
)

const (
	tileDir      = "tiles"
	TileFileName = "tile%v-%v_%vx%v.png"
)

func readTile(mapid string, z, x, y uint) (image.Image, error) {
	file, err := os.Open(fmt.Sprintf(tileDir+"/"+TileFileName, mapid, z, x, y))
	if err != nil {
		return nil, err
	}
	return png.Decode(file)
}
