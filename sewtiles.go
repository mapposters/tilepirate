package main

import (
	"fmt"
	"image/png"
	"os"
)

const (
	TileSize     = 512 // image tile size (px width and height)
	resultFolder = "results"
)

func main() {
	err := readConfig()
	if err != nil {
		return
	}

	for _, area := range areas {

		res, err := os.Create(resultFolder + "/" + area.Url + ".png") // For read access.
		if err != nil {
			fmt.Println(err)
			return
		}
		defer res.Close()

		m, err := traverseTiles(area)
		if err != nil {
			fmt.Println("Alas, no luck when traversing maptiles: ", err)
			return
		}
		png.Encode(res, m)
	}
}
