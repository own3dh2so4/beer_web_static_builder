package generator

import (
	"beer_web_static_builder/model"
	"html/template"
)

type HTMLGlass struct {
	Name          string
	ID            string
	Photo         string
	BoughtCity    string
	BoughtCountry string
	Got           string
	GotFrom       string
	Map           string
	FunctionName  template.JS
}

type HTMLBrand struct {
	Name           string
	ID             string
	FromCity       string
	FromCountry    string
	Website        string
	Map            string
	Logo           string
	BackgroundLogo bool
	NameImage      string
	Glasses        []HTMLGlass
}

func BuildHTMLBrand(brand model.Brand) HTMLBrand {
	return HTMLBrand{
		Name:           brand.Name,
		ID:             brand.ID,
		FromCity:       brand.FromCity,
		FromCountry:    brand.FromCountry,
		Website:        brand.Website,
		Map:            brand.Map,
		Logo:           brand.Logo,
		BackgroundLogo: brand.BackgroundLogo,
		NameImage:      brand.NameImage,
		Glasses:        buildHTMLGlasses(brand.Glasses),
	}
}

func buildHTMLGlasses(glasses []model.Glass) []HTMLGlass {
	var HTMLGlasses []HTMLGlass
	for _, glass := range glasses {
		HTMLGlasses = append(HTMLGlasses, BuildHTMLGlass(glass))
	}
	return HTMLGlasses
}

func BuildHTMLGlass(glass model.Glass) HTMLGlass {
	return HTMLGlass{
		Name:          glass.Name,
		ID:            glass.ID,
		Photo:         glass.Photo,
		BoughtCity:    glass.BoughtCity,
		BoughtCountry: glass.BoughtCountry,
		Got:           glass.Got,
		GotFrom:       glass.GotFrom,
		Map:           glass.Map,
		FunctionName:  template.JS(glass.ID),
	}
}
