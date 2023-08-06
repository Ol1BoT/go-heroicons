package main

import (
	"bytes"
	"embed"
	"fmt"
	"os"
	"path"
	"strings"
	"text/template"
)

//go:embed optimized/*
var FS embed.FS

type IconData struct {
	Name string
	Svg  string
}

func main() {

	// twentyPath := path.Join("optimized", "20")
	twentyFourPathSolid := path.Join("optimized", "24", "solid")
	twentyFourPathOutline := path.Join("optimized", "24", "outline")

	twentyFourDirSolid, err := os.ReadDir(twentyFourPathSolid)
	if err != nil {
		panic(err)
	}

	twentyFourDirOutline, err := os.ReadDir(twentyFourPathOutline)
	if err != nil {
		panic(err)
	}

	var iconsData []IconData

	for _, file := range twentyFourDirSolid {
		if file.IsDir() {
			continue
		}

		splitName := strings.Split(strings.TrimSuffix(file.Name(), ".svg"), "-")

		var newName string
		for _, part := range splitName {
			newName += strings.ToUpper(string(part[0])) + string(part[1:]) + "Solid"
		}

		fullFileDir := fmt.Sprintf("%s/%s", twentyFourPathSolid, file.Name())

		bt, err := os.ReadFile(fullFileDir)
		if err != nil {
			panic(err)
		}

		iconsData = append(iconsData, IconData{
			Name: newName,
			Svg:  string(bt),
		})
	}

	for _, file := range twentyFourDirOutline {
		if file.IsDir() {
			continue
		}

		splitName := strings.Split(strings.TrimSuffix(file.Name(), ".svg"), "-")

		var newName string
		for _, part := range splitName {
			newName += strings.ToUpper(string(part[0])) + string(part[1:]) + "Outline"
		}

		fullFileDir := fmt.Sprintf("%s/%s", twentyFourPathSolid, file.Name())

		bt, err := os.ReadFile(fullFileDir)
		if err != nil {
			panic(err)
		}

		iconsData = append(iconsData, IconData{
			Name: newName,
			Svg:  string(bt),
		})
	}

	var buf bytes.Buffer

	tpl, err := template.ParseFiles("heroicons.go.tmpl")
	tpl.Execute(&buf, iconsData)

	err = os.WriteFile("../heroicons.go", buf.Bytes(), 0777)
	if err != nil {
		panic(err)
	}
}
