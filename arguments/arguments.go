package arguments

import "flag"

const defaultDataPath  = "data/"
const defaultDataOutput = "output/"
const defaultTemplatePath = "templates/"

var DataPath string
var OutputPath string
var TemplatePath string

func init() {
	flag.StringVar(&DataPath, "data", defaultDataPath, "Path where the data is allocated")
	flag.StringVar(&OutputPath, "output", defaultDataOutput, "Path where the result will be allocated")
	flag.StringVar(&TemplatePath, "template-path", defaultTemplatePath,
		"Path where the templates are allocated")
	flag.Parse()
}
