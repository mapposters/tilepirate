package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
)

const (
	Margins = 200 // width of image's white margins
	Frame   = 10
)

func traverseTiles(tileRange image.Rectangle) (image.Image, error) {
	m := image.NewRGBA(image.Rect(-Margins, -Margins, tileRange.Dx()*TileSize+Margins, tileRange.Dy()*TileSize+Margins))

	whiteColor := color.RGBA{255, 255, 255, 255}
	draw.Draw(m, m.Bounds(), &image.Uniform{whiteColor}, image.ZP, draw.Src)

	frameColor := color.RGBA{50, 50, 50, 255}
	draw.Draw(m, image.Rect(-Frame, -Frame, tileRange.Dx()*TileSize+Frame, tileRange.Dy()*TileSize+Frame), &image.Uniform{frameColor}, image.ZP, draw.Src)

	position := image.Rectangle{image.ZP, image.Point{X: TileSize, Y: TileSize}}
	for y := tileRange.Min.Y; y < tileRange.Max.Y; y++ {
		for x := tileRange.Min.X; x < tileRange.Max.X; x++ {
			img, err := readTile(x, y)
			if err != nil {
				fmt.Println("no luck reading this tile:")
				fmt.Println(err)
				return nil, err
			}
			fmt.Printf("...adding %vx%v \n", x, y)
			draw.Draw(m, position, img, image.ZP, draw.Src)
			position = position.Add(image.Point{X: TileSize, Y: 0})
		}
		position = image.Rectangle{image.Point{X: 0, Y: (y - tileRange.Min.Y + 1) * TileSize}, image.Point{X: TileSize, Y: (y - tileRange.Min.Y + 2) * TileSize}}
	}
	fmt.Println("Great success!")
	return m, nil
}
