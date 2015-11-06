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
 "fmt"
	"io/ioutil"
	"os"
	"text/template"
)

/*
*Dockerfile struct that gets populated by user input
*
*/
type Dockerfile struct {
    Add []string
    Cmd []string
    Copy []string
    Entrypoint string
    Env map[string]string
    Expose []string
    From string
    Label map[string]string
    Maintainer string
    Run []string
    User int
    Volume []string
    Workdir string
}

/*
*Generic Error handling in go
*/
func check(e error) {
    if e != nil {
        panic(e)
    }
}

/*
*Function that reads in user filepath from command line
*/
func readInput() string {
    fmt.Print("Enter desired Dockerfile output filepath: ")
    var input string
    fmt.Scanln(&input)
    fmt.Println("Output dockerfile to: " + input)
    return input
}

/*
*Returns whether or not the given file or directory exists
*/
func exists(path string) (bool) {
    _, err := os.Stat(path)
    if err == nil { return true }
    if os.IsNotExist(err) { return false }
    return true
}

/*
*This function takes []byte as a parameter and writes to Dockerfile
*/
func writeToFile(templateBytes []byte) {
 
    s := readInput()
    
    if exists(s) { 
       ioutil.WriteFile(s + "Dockerfile", templateBytes, os.ModePerm)
       fmt.Print("Dockerfile Generated")
    } else {
       fmt.Print("Filepath did not previously exist...creating now")
       os.Mkdir(s, os.ModePerm)
       ioutil.WriteFile(s + "Dockerfile", templateBytes, os.ModePerm)
    }
 
}

func main() {

    example := &Dockerfile{
        []string{"/test1","/test2"}, 
        []string{"echo", "This is a test" }, 
        []string{"src1", "src2", "src3", "dest"},
        "/home", 
        map[string]string{"key":"value", "anotherKey":"anotherValue"},
        []string{"testExpose"}, 
        "testFrom",
        map[string]string{"key":"value", "anotherKey":"anotherValue"},
        "testMaintainer", 
        []string{"testRun"}, 
        999, 
        []string{"testVolume"}, 
        "/testWorkdir"}
    
    //Reading in template file to begin transformation
    content, _ := ioutil.ReadFile("../templates/DockerFile.tmpl")
	   tmpl, err := template.New("DockerFile").Parse(string(content))
	   check(err)
	   
	   //Executes template and stores into []byte
	   var templateOutput bytes.Buffer
    tmpl.Execute(&templateOutput, example)
    s := templateOutput.Bytes() 
	   
	   writeToFile(s)
    
}


