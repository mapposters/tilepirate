package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
)

const (
	TILEFOLDER           = "output" // dir path where tile files will be looked up
	TILE_SIZE            = 512      // image tile size (px width and height)
	RESULT_FILE          = "result.png"
	TILE_FILENAME_FORMAT = "tile_%v_%v.png"
)

var tileRange = image.Rectangle{
	Min: image.Point{X: 655, Y: 1581},
	Max: image.Point{X: 675, Y: 1591},
}

func readTile(x, y int) (image.Image, error) {
	file, err := os.Open(fmt.Sprintf(TILEFOLDER+"/"+TILE_FILENAME_FORMAT, x, y))
	if err != nil {
		return nil, err
	}
	return png.Decode(file)
}

func main() {
	m := image.NewRGBA(image.Rect(0, 0, tileRange.Dx()*TILE_SIZE, tileRange.Dy()*TILE_SIZE))
	fmt.Println("MOTHERFUCKER!!!")
	res, err := os.Create(RESULT_FILE) // For read access.
	if err != nil {
		log.Fatal(err)
	}
	defer res.Close()

	blue := color.RGBA{0, 0, 255, 255}
	draw.Draw(m, m.Bounds(), &image.Uniform{blue}, image.ZP, draw.Src)
	img, err := readTile(656, 1582)
	if err != nil {
		fmt.Println("no luck reading this tile:")
		fmt.Println(err)
		return
	}
	draw.Draw(m, m.Bounds(), img, image.ZP, draw.Src)
	png.Encode(res, m)
}
