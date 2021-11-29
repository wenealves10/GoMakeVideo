package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"sort"

	"github.com/icza/mjpeg"
)

func main() {
	checkErr := func(err error) {
		if err != nil {
			panic(err)
		}
	}

	outName := "file.mp4"
	aw, err := mjpeg.New(outName, 1336, 700, 10)
	checkErr(err)

	// Create a movie from images:
	matches, err := filepath.Glob("*.jpg")
	checkErr(err)
	sort.Strings(matches)

	fmt.Println("Found images:", matches)
	for _, name := range matches {
		data, err := ioutil.ReadFile(name)
		checkErr(err)
		checkErr(aw.AddFrame(data))
	}

	checkErr(aw.Close())
	fmt.Printf("%s was written successfully.\n", outName)

}
