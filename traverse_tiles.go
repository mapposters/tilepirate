package main

import (
	"fmt"
	"image"
	"image/draw"
)

func traverseTiles(tileRange image.Rectangle) (image.Image, error) {
	m := image.NewRGBA(image.Rect(0, 0, tileRange.Dx()*TileSize, tileRange.Dy()*TileSize))
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
