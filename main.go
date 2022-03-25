package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var vimModeline = "# vim:fileencoding=utf-8:ft=kitty"

func main() {
	var vim bool
	flag.BoolVar(&vim, "vim", false, "prepare the output for use in nvim.")
	flag.Parse()

	filename := flag.Arg(0)
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

	writeKittySessionFile(state, ProgramArguments{vim})
}

func writeKittySessionFile(state []OSWindow, opts ProgramArguments) {
	// Don't try to write if the file will be empty
	if len(state) == 0 {
		log.Fatal("The kitty session file is empty!")
	}

	file, err := os.Create("session.conf")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()
	buffer := bufio.NewWriter(file)
	lines := []string{}
	for _, window := range state {
		for _, t := range window.Tabs {
			lines = append(lines, fmt.Sprintf("new_tab %s", t.Title)+"\n")
			for _, w := range t.Windows {
				lines = append(lines, fmt.Sprintf("cd %s", w.Cwd)+"\n")
				lines = append(lines, "launch "+strings.Join(w.Cmdline, " ")+"\n")
			}
			lines = append(lines, "\n")
		}
	}
	if opts.Vim {
		lines = append(lines, vimModeline)
	}
	for _, line := range lines {
		_, err := buffer.WriteString(line)
		if err != nil {
			log.Fatal(err.Error())
		}
	}
	if err := buffer.Flush(); err != nil {
		log.Fatal(err.Error())
	}
}
