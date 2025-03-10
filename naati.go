package naati

import (
	"encoding/json"
	"io"
	"os"
	"text/template"
)

type Config struct {
	Input     string
	Output    string
	Templates *template.Template
}

func PrepareDialgoueFile(conf Config) error {
	jsonFile, err := os.Open(conf.Input)
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	var dialogues []Dialogue
	err = json.Unmarshal(byteValue, &dialogues)
	if err != nil {
		return err
	}

	// open output file
	output, err := os.Create(conf.Output)
	if err != nil {
		return err
	}

	err = conf.Templates.ExecuteTemplate(output, "layout.tmpl", dialogues)
	if err != nil {
		return err
	}

	return nil
}
