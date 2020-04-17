package repository

import (
	"beer_web_static_builder/model"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
)


type Repository interface {
	GetBrands() (model.Brands, error)
}

func NewFileRepository(dataPath string) Repository {
	return fileRepository{path: dataPath}
}

type fileRepository struct {
	path string
}

func (f fileRepository) GetBrands() (model.Brands, error) {

	var brands []model.Brand

	err := filepath.Walk(f.path, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) == ".yaml" {
			yamlFile, err := ioutil.ReadFile(path)
			if err != nil {
				fmt.Printf("Error reading YAML file: %s\n", err)
				return err
			}

			brand := model.Brand{BackgroundLogo: false}
			err = yaml.Unmarshal(yamlFile, &brand)
			if err != nil {
				fmt.Printf("Error parsing YAML file: %s\n", err)
				return err
			}
			brands = append(brands, brand)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("walk error [%v]\n", err)
		return model.Brands{}, err
	}

	return model.Brands{Brands: brands}, nil
}
