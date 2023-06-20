package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"sort"

	"golang.org/x/image/colornames"
)

func main() {

	outputFile := os.Args[1] + ".gif"
	if len(outputFile) == 0 {
		outputFile = "output.gif"
	}

	filepaths, err := filepath.Glob("*.png")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	sort.Strings(filepaths)

	var images []*image.Paletted
	var delays []int

	var palette = []color.Color{colornames.Green, color.Black, color.White,
		colornames.Red, colornames.Grey, colornames.Blue, colornames.Brown}
	images = append(images, image.NewPaletted(image.Rect(0, 0, 550, 550), palette))
	delays = append(delays, 0)
	for i := range filepaths {
		file, err := os.Open(fmt.Sprintf("%d.gv.png", i))
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		defer file.Close()

		img, err := png.Decode(file)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		palettedImg := image.NewPaletted(image.Rect(0, 0, 550, 550), palette)
		draw.Draw(palettedImg, palettedImg.Bounds(), img, image.Point{}, draw.Src)
		images = append(images, palettedImg)
		delays = append(delays, 100)
		e := os.Remove(file.Name())
		if e != nil {
			log.Fatal(e)
		}
	}

	// Create the output GIF file
	outfile, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer outfile.Close()

	// Encode the GIF and write it to the output file
	err = gif.EncodeAll(outfile, &gif.GIF{Image: images, Delay: delays})
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("GIF created successfully!")
}
