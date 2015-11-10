/*
TODO
    Need to be able to define order
    Which of these can happen multiple times?
    Which of these are REQUIRED vs OPTIONAL?
    Define relationships between members - see https://docs.docker.com/reference/builder/
*/

/* Onbuild
   ONBUILD is sort of a nested Dockerfile, but without ONBUILD, FROM, or MAINTAINER instructions
   -- Not sure how inheritence works but we could have a base onbuild or something that we inherit from
*/

package main

import (
	"bytes"
	"dockerfile"
	"io/ioutil"
	"text/template"
)

func main() {

	example := dockerfile.Dockerfile{
		[]string{"/test1", "/test2"},
		[]string{"echo", "This is a test"},
		[]string{"src1", "src2", "src3", "dest"},
		"/home",
		map[string]string{"key": "value", "anotherKey": "anotherValue"},
		[]string{"testExpose"},
		"testFrom",
		map[string]string{"key": "value", "anotherKey": "anotherValue"},
		"testMaintainer",
		[]string{"testRun"},
		999,
		[]string{"testVolume"},
		"/testWorkdir"}

	//Reading in template file to begin transformation
	content, _ := ioutil.ReadFile("../templates/Dockerfile.tmpl")
	tmpl, err := template.New("Dockerfile").Parse(string(content))
	dockerfile.Check(err)

	//Executes template and stores into []byte
	var templateOutput bytes.Buffer
	tmpl.Execute(&templateOutput, example)
	s := templateOutput.Bytes()

	dockerfile.WriteToFile(s)

}
