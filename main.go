package main

import (
	"embed"
	"io"
	"os"
	"path"
	"text/template"
)

var (
	//go:embed templates/*
	templates embed.FS
)

type Meta struct {
	ProjectName string
	Author      string
	License     string
}

func pwd() string {
	wd, _ := os.Getwd()
	return wd
}

func parseTemplate(tplName string, meta Meta) {
	fp, err := templates.Open("templates/" + tplName)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	content, err := io.ReadAll(fp)
	if err != nil {
		panic(err)
	}

	parser := template.Must(template.New(tplName).Parse(string(content)))
	output, err := os.OpenFile(tplName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer output.Close()

	err = parser.Execute(output, meta)
	if err != nil {
		panic(err)
	}
}

func main() {
	meta := Meta{
		ProjectName: path.Base(pwd()),
		Author:      "yuekcc",
		License:     "MIT",
	}

	templateList := []string{
		".editorconfig",
		".gitignore",
		"package.json",
	}

	for _, name := range templateList {
		parseTemplate(name, meta)
	}
}
