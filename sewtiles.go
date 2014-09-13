package main

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
)

const (
	TILEFOLDER           = "output" // dir path where tile files will be looked up
	TileSize             = 512      // image tile size (px width and height)
	RESULT_FILE          = "result.png"
	TILE_FILENAME_FORMAT = "tile_%v_%v.png"
)

func readTile(x, y int) (image.Image, error) {
	file, err := os.Open(fmt.Sprintf(TILEFOLDER+"/"+TILE_FILENAME_FORMAT, x, y))
	if err != nil {
		return nil, err
	}
	return png.Decode(file)
}

func main() {
	res, err := os.Create(RESULT_FILE) // For read access.
	if err != nil {
		log.Fatal(err)
	}
	defer res.Close()

	var tileRange = image.Rectangle{
		Min: image.Point{X: 655, Y: 1581},
		Max: image.Point{X: 660, Y: 1591},
	}
	m, err := traverseTiles(tileRange)
	if err != nil {
		fmt.Println("Alas, no luck when traversing maptiles: ", err)
		return
	}
	png.Encode(res, m)
}
