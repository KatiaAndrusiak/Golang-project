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

	anim := gif.GIF{}

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
		var palette = []color.Color{colornames.Green, color.Black, color.White,
			colornames.Red, colornames.Grey, colornames.Blue, colornames.Brown}

		palettedImg := image.NewPaletted(img.Bounds(), palette)
		draw.Draw(palettedImg, palettedImg.Bounds(), img, image.Point{}, draw.Src)

		anim.Image = append(anim.Image, palettedImg)
		anim.Delay = append(anim.Delay, 100)

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
	err = gif.EncodeAll(outfile, &anim)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("GIF created successfully!")
}
