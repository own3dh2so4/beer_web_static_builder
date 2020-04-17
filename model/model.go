package model

type Glass struct {
	Name          string `yaml:"name"`
	ID            string `yaml:"id"`
	Photo         string `yaml:"photo"`
	BoughtCity    string `yaml:"bought_city"`
	BoughtCountry string `yaml:"bought_country"`
	Got           string `yaml:"got"`
	GotFrom       string `yaml:"got_from"`
	Map           string `yaml:"map"`
}

type Brand struct {
	Name           string  `yaml:"name"`
	ID             string  `yaml:"id"`
	FromCity       string  `yaml:"from_city"`
	FromCountry    string  `yaml:"from_country"`
	Website        string  `yaml:"website"`
	Map            string  `yaml:"map"`
	Logo           string  `yaml:"logo"`
	BackgroundLogo bool    `yaml:"background_logo"`
	NameImage      string  `yaml:"name_image"`
	Glasses        []Glass `yaml:"glasses"`
}

type Brands struct {
	Brands []Brand
}
