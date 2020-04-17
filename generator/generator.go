package generator

import (
	"beer_web_static_builder/model"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

const imgHTMLPath = "img"
const staticDirName = "static"
const defaultIndexTemplateName = "index.html"
const defaultPageTemplateName = "page.html"

var funcMap = template.FuncMap{
	"inc": func(i int) int {
		return i + 1
	},
}

type Generator interface {
	Generate(brands model.Brands) error
}

func NewHTMLGenerator(dataPath, templatesPath, resultPath string) Generator {
	return htmlGenerator{
		dataPath: dataPath,
		templatesPath: templatesPath,
		indexTemplateName: defaultIndexTemplateName,
		pageTemplateName:  defaultPageTemplateName,
		resultPath:        resultPath,
	}
}

func copyFile(sourceFile, destinationFile string) error {
	input, err := ioutil.ReadFile(sourceFile)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if err := createDirectoryIfNotExists(path.Dir(destinationFile)); err != nil {
		return err
	}
	err = ioutil.WriteFile(destinationFile, input, 0644)
	if err != nil {
		fmt.Println("Error creating", destinationFile)
		fmt.Println(err)
		return err
	}
	return nil
}

func createDirectoryIfNotExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err = os.MkdirAll(path, 0755); err != nil {
			return err
		}
	}
	return nil
}


type htmlGenerator struct {
	dataPath string
	templatesPath string
	indexTemplateName string
	pageTemplateName string
	resultPath string
}

func (hg htmlGenerator) getIndexTemplatePath() string {
	return path.Join(hg.templatesPath, hg.indexTemplateName)
}

func (hg htmlGenerator) getPageTemplatePath() string {
	return path.Join(hg.templatesPath, hg.pageTemplateName)
}

func (hg htmlGenerator) generateIndex(brands model.Brands) error {
	tmpl := template.Must(template.New(hg.indexTemplateName).Funcs(funcMap).ParseFiles(hg.getIndexTemplatePath()))

	f, err := os.Create(path.Join(hg.resultPath, defaultIndexTemplateName))
	if err != nil {
		return err
	}
	if err = tmpl.Execute(f, brands); err != nil {
		return err
	}
	if err = f.Close(); err != nil {
		return err
	}
	return nil
}

func (hg htmlGenerator) generatePage(brand HTMLBrand) error {
	tmpl := template.Must(template.New(hg.pageTemplateName).Funcs(funcMap).ParseFiles(hg.getPageTemplatePath()))
	f, err := os.Create(path.Join(hg.resultPath, brand.ID + ".html"))
	if err != nil {
		return err
	}
	if err = tmpl.Execute(f, brand); err != nil {
		return err
	}
	if err = f.Close(); err != nil {
		return err
	}
	return nil
}

func (hg htmlGenerator) copyImages(brand model.Brand) error {
	logoPath := path.Join(hg.resultPath, imgHTMLPath, brand.Logo)
	err := copyFile(path.Join(hg.dataPath, brand.Logo), logoPath)
	if err != nil {
		return err
	}
	nameImagePath := path.Join(hg.resultPath, imgHTMLPath, brand.NameImage)
	err = copyFile(path.Join(hg.dataPath, brand.NameImage), nameImagePath)
	if err != nil {
		return err
	}
	for _, glass := range brand.Glasses {
		glassPath := path.Join(hg.resultPath, imgHTMLPath, glass.Photo)
		err := copyFile(path.Join(hg.dataPath, glass.Photo), glassPath)
		if err != nil {
			return err
		}
	}
	return nil
}

func (hg htmlGenerator) copyStaticFiles() error {
	staticPath := path.Join(hg.templatesPath, staticDirName)
	dataPathLen := len(hg.templatesPath)
	err := filepath.Walk(staticPath, func(p string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		return copyFile(p, path.Join(hg.resultPath, p[dataPathLen:]))
	})
	return err
}

func (hg htmlGenerator) copyResources(brands model.Brands) error {
	for _, b := range brands.Brands {
		if err := hg.copyImages(b); err != nil {
			return err
		}
	}
	return hg.copyStaticFiles()
}

func (hg htmlGenerator) Generate(brands model.Brands) error {
	if err := createDirectoryIfNotExists(hg.resultPath); err != nil {
		return err
	}
	err := hg.generateIndex(brands)
	if err != nil {
		return err
	}
	for _, b := range brands.Brands {
		if err = hg.generatePage(BuildHTMLBrand(b)); err != nil {
			return err
		}
	}
	return hg.copyResources(brands)
}

