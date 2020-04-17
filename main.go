package main

import (
	"beer_web_static_builder/arguments"
	"beer_web_static_builder/generator"
	"beer_web_static_builder/repository"
)


func main() {
	brands, err := repository.NewFileRepository(arguments.DataPath).GetBrands()
	if err != nil {
		panic(err)
	}
	err = generator.NewHTMLGenerator(arguments.DataPath, arguments.TemplatePath, arguments.OutputPath).Generate(brands)
	if err != nil {
		panic(err)
	}
	//fmt.Printf("%v", brands)
}
