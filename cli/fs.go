package cliq

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	logq "github.com/dx-luna/go-utils/log"
)

func Check(e error, desc string) {
	if e != nil {
		logq.Error(desc)
		panic(e)
	}
}
func ReadFile(dir string, dirname string, result any) error {
	content, err := os.ReadFile(dir)
	Check(err, "failed ReadFile : "+dirname)

	logq.Success("success open file ", dirname)
	err = json.Unmarshal(content, &result)
	Check(err, "failed unmarshal .json file : "+dirname)

	return nil
}

func WriteFile(dir string, data interface{}) {

	file, err := json.MarshalIndent(data, "", " ")
	Check(err, "error json.MarshalIndent")

	err = ioutil.WriteFile(dir, file, 0644)
	Check(err, "error writeFile")
}
func ReadDir(dir string) []string {

	fileNames := make([]string, 0)
	outputDirRead, err := os.Open(dir)
	Check(err, "Failed readDir")

	outputDirFiles, _ := outputDirRead.ReadDir(0)
	for outputIndex := range outputDirFiles {
		outputFileHere := outputDirFiles[outputIndex]

		// Get name of file.
		outputNameHere := outputFileHere.Name()

		// Print name.
		fmt.Println(outputNameHere)
		fileNames = append(fileNames, outputNameHere)
	}
	return fileNames
}
func Resolve(fname string) string {
	path, err := filepath.Abs(fname)
	Check(err, "no file on "+fname)

	return path
}
