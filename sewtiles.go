package main

import (
  "fmt"
  "image/png"
  "os"
)

const (
  TileSize     = 1024 // image tile size (px width and height)
  resultFolder = "results"
)

func main() {
  fmt.Printf("loading configs..\n")
  err := readConfig()
  if err != nil {
    return
  }
  for styleName, style := range styles {
    for _, area := range areas {
      res, err := os.Create(resultFolder + "/" + area.Url + "(in " + styleName + ").png") // For read access.
      if err != nil {
        fmt.Println(err)
        return
      }
      defer res.Close()

      m, err := traverseTiles(style, area)
      if err != nil {
        fmt.Println("Alas, no luck when traversing maptiles: ", err)
        return
      }
      png.Encode(res, m)
    }
  }
}
