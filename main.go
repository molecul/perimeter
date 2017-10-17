package main

import (
	"github.com/v0rt3x/perimeter"
	"github.com/v0rt3x/perimeter/config"
)

func main() {
	perimeter_config := config.NewPerimeterConfig("perimeter.yml")
	perimeter_config.Read()
	perimeter.Run(perimeter_config)
}
