package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

var vimModeline = "# vim:fileencoding=utf-8:ft=kitty"

func main() {
	args := ProgramArgs{}
	flag.BoolVar(&args.Vim, "vim", false, "prepare the output for use in nvim.")
	flag.StringVar(&args.Filename, "filename", "session.conf", "prepare the output for use in nvim.")
	flag.StringVar(&args.Source, "source", "", "kitty session json file to restore from")
	flag.Parse()

	var bytes []byte
	if len(args.Source) > 0 {
		bytes = readSessionFromFile(args.Source)
	} else {
		bytes = readSessionFromKitty()
	}
	var state []OSWindow
	err := json.Unmarshal(bytes, &state)
	if err != nil {
		log.Fatal(err.Error())
	}
	writeSessionFile(state, args)
}

func readSessionFromKitty() []byte {
	rsp, err := exec.Command("kitty", "@", "ls").Output()
	if err != nil {
		log.Fatal(err.Error())
	}
	return rsp
}

func readSessionFromFile(filename string) []byte {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err.Error())
	}
	return bytes
}

// writeSessionFile Creates a session file based on kitty's current state
func writeSessionFile(state []OSWindow, args ProgramArgs) {
	if len(state) == 0 {
		log.Fatal("The kitty session file is empty!")
	}

	file, err := os.Create(args.Filename)
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
	if args.Vim {
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
