package main

import (
	"fmt"
	"log"
	"os"

	"github.com/swaggo/swag"
	"github.com/swaggo/swag/gen"
	"github.com/urfave/cli"
)

const (
	searchDirFlag        = "dir"
	generalInfoFlag      = "generalInfo"
	propertyStrategyFlag = "propertyStrategy"
	outputFlag           = "output"
	parseVendorFlag      = "parseVendor"
	parseDependencyFlag  = "parseDependency"
	markdownFilesFlag    = "markdownFiles"
)

var initFlags = []cli.Flag{
	cli.StringFlag{
		Name:  generalInfoFlag + ", g",
		Value: "main.go",
		Usage: "Go file path in which 'swagger general API Info' is written",
	},
	cli.StringFlag{
		Name:  searchDirFlag + ", d",
		Value: "./",
		Usage: "Directory you want to parse",
	},
	cli.StringFlag{
		Name:  propertyStrategyFlag + ", p",
		Value: "camelcase",
		Usage: "Property Naming Strategy like snakecase,camelcase,pascalcase",
	},
	cli.StringFlag{
		Name:  outputFlag + ", o",
		Value: "./docs",
		Usage: "Output directory for all the generated files(swagger.json, swagger.yaml and doc.go)",
	},
	cli.BoolFlag{
		Name:  parseVendorFlag,
		Usage: "Parse go files in 'vendor' folder, disabled by default",
	},
	cli.BoolFlag{
		Name:  parseDependencyFlag,
		Usage: "Parse go files in outside dependency folder, disabled by default",
	},
	cli.StringFlag{
		Name:  markdownFilesFlag + ", md",
		Value: "",
		Usage: "Parse folder containing markdown files to use as description, disabled by default",
	},
}

func initAction(c *cli.Context) error {
	strategy := c.String(propertyStrategyFlag)

	switch strategy {
	case swag.CamelCase, swag.SnakeCase, swag.PascalCase:
	default:
		return fmt.Errorf("not supported %s propertyStrategy", strategy)
	}

	return gen.New().Build(&gen.Config{
		SearchDir:          c.String(searchDirFlag),
		MainAPIFile:        c.String(generalInfoFlag),
		PropNamingStrategy: strategy,
		OutputDir:          c.String(outputFlag),
		ParseVendor:        c.Bool(parseVendorFlag),
		ParseDependency:    c.Bool(parseDependencyFlag),
		MarkdownFilesDir:   c.String(markdownFilesFlag),
	})
}

func main() {
	app := cli.NewApp()
	app.Version = swag.Version
	app.Usage = "Automatically generate RESTful API documentation with Swagger 2.0 for Go."
	app.Commands = []cli.Command{
		{
			Name:    "init",
			Aliases: []string{"i"},
			Usage:   "Create docs.go",
			Action:  initAction,
			Flags:   initFlags,
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
