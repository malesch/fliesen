package main

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"os"
)

func init() {
	image.RegisterFormat("gif", "gif", gif.Decode, gif.DecodeConfig)
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
	image.RegisterFormat("jpg", "jpg", jpeg.Decode, jpeg.DecodeConfig)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Missing image argument")
		os.Exit(1)
	}
	fileName := os.Args[1]
	fmt.Printf("Processing image '%s'\n", fileName)
	reader, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	img, _, err := image.Decode(reader)
	bounds := img.Bounds()
	fmt.Printf("Dimensions: width=%d, height=%d\n", (bounds.Max.X - bounds.Min.X), (bounds.Max.Y - bounds.Min.Y))

	// Frequency map
	pixelFrequency := make(map[string]int)

	var color string
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			color = fmt.Sprintf("%.2X%.2X%.2X", uint8(r), uint8(g), uint8(b))

			_, ok := pixelFrequency[color]
			if ok == true {
				pixelFrequency[color]++
			} else {
				pixelFrequency[color] = 1
			}
		}
	}

	fmt.Printf("\nPixel color frequencies:\n")
	total := 0
	for c, f := range pixelFrequency {
		total = total + f
		fmt.Println(c, ": ", f)
	}
	fmt.Printf("\nTotal number tiles: %d\n", total)
}
