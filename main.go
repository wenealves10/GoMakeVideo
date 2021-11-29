package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"sort"
	"time"

	"github.com/icza/mjpeg"
	cinema "github.com/jtguibas/cinema"
)

func main() {
	checkErr := func(err error) {
		if err != nil {
			panic(err)
		}
	}

	outName := "file.mp4"
	aw, err := mjpeg.New(outName, 144, 108, 60)
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

	video, err := cinema.Load("file.mp4")
	checkErr(err)

	video.Trim(10*time.Second, 20*time.Second) // trim video from 10 to 20 seconds
	video.SetStart(1 * time.Second)            // trim first second of the video
	video.SetEnd(9 * time.Second)              // keep only up to 9 seconds
	video.SetSize(400, 300)                    // resize video to 400x300
	video.Crop(0, 0, 200, 200)                 // crop rectangle top-left (0,0) with size 200x200
	video.SetSize(400, 400)                    // resize cropped 200x200 video to a 400x400
	video.SetFPS(48)                           // set the output framerate to 48 frames per second
	video.Render("test_output.mp4")            // note format conversion by file extension

	// you can also generate the command line instead of applying it directly
	fmt.Println("FFMPEG Command", video.CommandLine("test_output.mp4"))
}
