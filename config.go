package main

import (
	"encoding/json"
	"fmt"
	"image"
	"io/ioutil"
)

type area struct {
	Name      string
	Url       string
	Z         uint
	TileRange image.Rectangle
}

var areas map[string]area

const (
	configPath = "areas.json"
)

func readConfig() error {
	content, err := ioutil.ReadFile(configPath)
	if err != nil {
		fmt.Println(err)
		return err
	}
	json.Unmarshal(content, &areas)
	return nil
}
