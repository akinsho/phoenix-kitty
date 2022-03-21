package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	filename := os.Args[1]
	fmt.Println(filename)

	if filename == "" {
		log.Fatal("No file path was passed in!")
	}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err.Error())
	}

	var state []OSWindow
	err = json.Unmarshal(bytes, &state)
	if err != nil {
		log.Fatal(err.Error())
	}

	writeKittySessionFile(state)
}

func writeKittySessionFile(state []OSWindow) {
	file, err := os.Create("session.conf")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()
	buffer := bufio.NewWriter(file)
	for _, window := range state {
		for _, t := range window.Tabs {
			buffer.WriteString(fmt.Sprintf("new_tab %s", t.Title) + "\n")
			fmt.Println(t.Windows)
			for _, w := range t.Windows {
				buffer.WriteString(fmt.Sprintf("cd %s", w.Cwd) + "\n")
				buffer.WriteString("launch " + strings.Join(w.Cmdline, " ") + "\n")
			}
			buffer.WriteString("\n")
		}
	}
	if err := buffer.Flush(); err != nil {
		log.Fatal(err.Error())
	}
}
