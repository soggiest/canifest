
package dockerfile

import (
	"fmt"
	"io/ioutil"
	"os"
)

/*
 * Dockerfile struct that gets populated by user input
 */

type Dockerfile struct {
	Add        []string
	Cmd        []string
	Copy       []string
	Entrypoint string
	Env        map[string]string
	Expose     []string
	From       string
	Label      map[string]string
	Maintainer string
	Run        []string
	User       int
	Volume     []string
	Workdir    string
}


func NewDockerfile() Dockerfile {
	return Dockerfile{}
}


/*
 * Generic Error handling in go
 */

func Check(e error) {
	if e != nil {
		panic(e)
	}
}


/*
 * Function that reads in user filepath from command line
 */

func readInput() string {
	fmt.Print("Enter desired Dockerfile output filepath: ")
	var input string
	fmt.Scanln(&input)
	fmt.Println("Output dockerfile to: " + input)
	return input
}


/*
 * Returns whether or not the given file or directory exists
 */

func exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}


/*
 *This function takes []byte as a parameter and writes to Dockerfile
 */

func WriteToFile(templateBytes []byte) {

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

