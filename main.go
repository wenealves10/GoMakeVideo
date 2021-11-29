package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"sort"

	"github.com/icza/mjpeg"
	"github.com/sikang99/cinema"
)

func main() {
	checkErr := func(err error) {
		if err != nil {
			panic(err)
		}
	}

	outName := "file.mp4"
	aw, err := mjpeg.New(outName, 1336, 700, 1)
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

	// loading the test video
	// fmt.Println("Downloading Test Video...")
	// video_url := "https://media.w3.org/2010/05/sintel/trailer.mp4"
	// if err := DownloadFile("test.mp4", video_url); err != nil {
	// 	panic(err)
	// }

	// // initializing the test video as a cinema video object
	v, err := cinema.MakeVideo("file.mp4")
	if err != nil {
		fmt.Println(err)
	}

	// testing all setters
	// v.Trim(0, 10)
	v.SetSize(400, 400)
	v.Render("test_output.mp4")

	// testing all getters
	fmt.Println(v.Filepath())
	fmt.Println(v.Start())
	fmt.Println(v.End())
	fmt.Println(v.Width())
	fmt.Println(v.Height())
	fmt.Println(v.Duration())

}

// func DownloadFile(filepath string, url string) error {
// 	// check if the file exist already
// 	finfo, err := os.Stat(filepath)
// 	if err != nil {
// 		return err
// 	}
// 	log.Println(finfo.Name(), finfo.Size())

// 	// Get the data
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		return err
// 	}
// 	defer resp.Body.Close()

// 	// Create the file
// 	out, err := os.Create(filepath)
// 	if err != nil {
// 		return err
// 	}
// 	defer out.Close()

// 	// Write the body to file
// 	_, err = io.Copy(out, resp.Body)
// 	return err
// }
