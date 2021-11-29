package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"sort"

	"github.com/icza/mjpeg"
	"github.com/jtguibas/cinema"
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

	matches, err := filepath.Glob("*.jpg")
	checkErr(err)
	sort.Strings(matches)

	fmt.Println("Imagens encontradas:", matches)
	for _, name := range matches {
		data, err := ioutil.ReadFile(name)
		checkErr(err)
		checkErr(aw.AddFrame(data))
	}

	checkErr(aw.Close())
	fmt.Printf("%s Baixou com sucesso\n", outName)

	v, err := cinema.Load("file.mp4")
	if err != nil {
		fmt.Println(err)
	}

	v.SetStart(0)
	v.SetEnd(10)
	v.Render("file_modificar.mp4")

	fmt.Println(v.Filepath())
	fmt.Println(v.Start())
	fmt.Println(v.End())
	fmt.Println(v.Width())
	fmt.Println(v.Height())
	fmt.Println(v.Duration())

}
