package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

import (
	"github.com/pkg/errors"
)

import (
	"github.com/sandhillgeo/go-gpkg/gpkg"
)

var GO_GPKG_VERSION = "0.0.1"

func printUsage() {
	fmt.Println("Usage: gpkg -output_uri OUTPUT_URI [-related_tables] [-version] [-help]")
}

func main() {

	start := time.Now()

	var output_uri string
	//var verbose bool
	var related_tables bool
	var version bool
	var help bool

	flag.StringVar(&output_uri, "output_uri", "", "The output uri of the GeoPackage.")

	flag.BoolVar(&related_tables, "related_tables", false, "Includes tables for the Related Tables Extension.")

	flag.BoolVar(&version, "version", false, "Prints version to stdout.")
	flag.BoolVar(&help, "help", false, "Print help.")

	flag.Parse()

	if help {
		printUsage()
		fmt.Println("Options:")
		flag.PrintDefaults()
		os.Exit(0)
	} else if len(os.Args) == 1 {
		fmt.Println("Error: Provided no arguments.")
		fmt.Println("Run \"gpkg -help\" for more information.")
		os.Exit(0)
	} else if len(os.Args) == 2 && os.Args[1] == "help" {
		printUsage()
		fmt.Println("Options:")
		flag.PrintDefaults()
		os.Exit(0)
	}

	if version {
		fmt.Println(GO_GPKG_VERSION)
		os.Exit(0)
	}

	if len(output_uri) == 0 {
		fmt.Println("Error: Provided no -output_uri.")
		fmt.Println("Run \"gpkg -help\" for more information.")
		os.Exit(1)
	}

	g := gpkg.New(output_uri)

	err := g.Init()
	if err != nil {
		fmt.Println(errors.Wrap(err, "Error initializing GeoPackage"))
		os.Exit(1)
	}

	err = g.AutoMigrate()
	if err != nil {
		fmt.Println(errors.Wrap(err, "Error auto migrating GeoPackage"))
		os.Exit(1)
	}

	if related_tables {
		err = g.AutoMigrateRelatedTables()
		if err != nil {
			fmt.Println(errors.Wrap(err, "Error auto migrating tables for the Related Tables Extension"))
			os.Exit(1)
		}
	}

	err = g.Close()
	if err != nil {
		fmt.Println(errors.Wrap(err, "Error closing GeoPackage"))
		os.Exit(1)
	}

	elapsed := time.Since(start)
	fmt.Println("Done in " + elapsed.String())
}
