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

type style struct {
	Mapid string
}

var areas map[string]area
var styles map[string]style

const (
	configPath = "areas.json"
	stylesPath = "styles.json"
)

func readConfig() error {
	// areas config
	content, err := ioutil.ReadFile(configPath)
	if err != nil {
		fmt.Println(err)
		return err
	}
	json.Unmarshal(content, &areas)

	// styles config
	content, err = ioutil.ReadFile(stylesPath)
	if err != nil {
		fmt.Println(err)
		return err
	}
	json.Unmarshal(content, &styles)

	return nil
}
