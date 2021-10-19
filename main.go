package main

import (
	"embed"
	"flag"
	"io"
	"os"
	"path"
	"text/template"
)

var (
	//go:embed templates/*
	templates embed.FS

	authorFlag  string
	licenseFlag string
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

func init() {
	flag.StringVar(&authorFlag, "author", "yuekcc", "setup project author, default: yuekcc")
	flag.StringVar(&licenseFlag, "license", "MIT", "setup project license, default: MIT")
}

func main() {
	flag.Parse()

	meta := Meta{
		ProjectName: path.Base(pwd()),
		Author:      authorFlag,
		License:     licenseFlag,
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
