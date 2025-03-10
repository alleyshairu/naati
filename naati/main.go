package main

import (
	"embed"
	"log"
	"naati"
	"text/template"

	"github.com/alexflint/go-arg"
)

//go:embed templates/*
var templates embed.FS

func main() {

	templates, err := template.New("").Funcs(naati.GetFuncsMap()).ParseFS(templates, "templates/*.tmpl")
	if err != nil {
		log.Fatal(err)
		return
	}

	var args struct {
		Input  string `arg:"required"`
		Output string `arg:"required"`
	}

	arg.MustParse(&args)

	config := naati.Config{
		Input:     args.Input,
		Output:    args.Output,
		Templates: templates,
	}

	err = naati.PrepareDialgoueFile(config)
	if err != nil {
		log.Fatalf("%v", err)
		return
	}
}
