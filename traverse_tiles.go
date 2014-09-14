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

func traverseTiles(Area area) (image.Image, error) {
	m := image.NewRGBA(image.Rect(-Margins, -Margins, Area.TileRange.Dx()*TileSize+Margins, Area.TileRange.Dy()*TileSize+Margins))

	whiteColor := color.RGBA{255, 255, 255, 255}
	draw.Draw(m, m.Bounds(), &image.Uniform{whiteColor}, image.ZP, draw.Src)

	frameColor := color.RGBA{50, 50, 50, 255}
	draw.Draw(m, image.Rect(-Frame, -Frame, Area.TileRange.Dx()*TileSize+Frame, Area.TileRange.Dy()*TileSize+Frame), &image.Uniform{frameColor}, image.ZP, draw.Src)

	position := image.Rectangle{image.ZP, image.Point{X: TileSize, Y: TileSize}}
	for y := Area.TileRange.Min.Y; y < Area.TileRange.Max.Y; y++ {
		for x := Area.TileRange.Min.X; x < Area.TileRange.Max.X; x++ {
			img, err := readTile(Area.Z, uint(x), uint(y))
			if err != nil {
				fmt.Printf("no luck reading this tile: z:%v, %vx %v \n", Area.Z, x, y)
				fmt.Println(err)
				break
			}
			fmt.Printf("...adding %vx%v \n", x, y)
			draw.Draw(m, position, img, image.ZP, draw.Src)
			position = position.Add(image.Point{X: TileSize, Y: 0})
		}
		position = image.Rectangle{image.Point{X: 0, Y: (y - Area.TileRange.Min.Y + 1) * TileSize}, image.Point{X: TileSize, Y: (y - Area.TileRange.Min.Y + 2) * TileSize}}
	}
	fmt.Println("Great success!")
	return m, nil
}
