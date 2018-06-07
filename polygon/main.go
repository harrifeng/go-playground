package main

import (
	"encoding/json"
	"fmt"
	"os"

	geo "github.com/kellydunn/golang-geo"
)

func main() {
	// latitude, longtitude
	point := geo.NewPoint(4.9402900, 114.9480600)
	demo, err := polygonFromFile("./data/demo.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(demo.Contains(point))
	os.Exit(0)
}

type testPoints struct {
	Points []*geo.Point
}

func polygonFromFile(filename string) (*geo.Polygon, error) {
	p := &geo.Polygon{}
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	points := new(testPoints)
	jsonParser := json.NewDecoder(file)
	if err = jsonParser.Decode(&points); err != nil {
		return nil, err
	}

	for _, point := range points.Points {
		p.Add(point)
	}

	return p, nil
}
