package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	filename := os.Args[1]
	fmt.Println(filename)

	if filename == "" {
		fmt.Println("No file path was passed in!")
		os.Exit(1)
	}

	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Unable to open the path: %s", filename)
		os.Exit(1)
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("Unable to read contents of the file: %s", filename)
		os.Exit(1)
	}
	
	var state []OSWindow
	err = json.Unmarshal(bytes, &state)
	if err != nil {
		fmt.Printf("Unable to unmarshal the session file %s: %s", filename, err.Error())
		os.Exit(1)
	}
	fmt.Println(state)
}
