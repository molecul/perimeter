package main

import (
	"v0rt3x/perimeter"
	"v0rt3x/perimeter/config"
)

func main() {
	perimeter_config := config.PerimeterConfig{ConfigPath: "perimeter.yml"}
	perimeter_config.Read()

	perimeter.Run(&perimeter_config)
}
