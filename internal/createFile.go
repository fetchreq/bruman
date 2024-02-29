package internal

import (
	"fmt"
	"log"
	"os"
	"path"
	"text/template"
)

func CreateHtml(bruFile BruFile) {

	tmplFile := "./internal/bru.tmpl"
	name := path.Base(tmplFile)
	tmpl, err := template.New(name).ParseFiles(tmplFile)
	fmt.Println(bruFile.Summary.TotalRequests)
	if err != nil {
		log.Fatal(err)
	}
	var f *os.File
	f, err = os.Create("bru.html")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(f, bruFile)
	if err != nil {
		panic(err)
	}
	err = f.Close()
	if err != nil {
		panic(err)
	}
}
