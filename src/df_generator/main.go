package main

import (
	"io/ioutil"
	"log"
	"os"
	"text/template"
)

type DockerFile struct {
	From       string
	Maintainer string
	Label      string
}

func main() {
	df := DockerFile{"Fedora", "Who Knows <whothere@redhat.com>", "Description=\"This image is just an example...\""}

	content, _ := ioutil.ReadFile("templates/DockerFile.tmpl")
	tmpl, err := template.New("DockerFile").Parse(string(content))
	if err != nil {
		log.Panic(err)
	}

	tmpl.Execute(os.Stdout, df)
}
