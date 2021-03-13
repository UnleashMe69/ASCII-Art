package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"log"
	"math"
	"os"
)

func main() {

	reader, err := os.Open("ascii-pineapple.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	m, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully loaded image!")

	bounds := m.Bounds()

	imageWidth := bounds.Max.X - bounds.Min.X
	imageHeight := bounds.Max.Y - bounds.Min.Y

	fmt.Printf("Image size: %d x %d\n", imageWidth, imageHeight)

	var pixelMatrix [][]pixel
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		var row []pixel
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			row = append(row, rgbToPixel(m.At(x, y).RGBA()))
		}
		pixelMatrix = append(pixelMatrix, row)
	}

	fmt.Println("\nSuccessfully constructed pixel matrix!")
	fmt.Printf("Pixel matrix size: %d x %d\n", len(pixelMatrix[0]), len(pixelMatrix))
	fmt.Println("Iterating through pixel contents:")

	var pixelBrightness [][]float64
	for y := 0; y < len(pixelMatrix); y++ {
		var row []float64
		for x := 0; x < len(pixelMatrix[0]); x++ {
			row = append(row, float64((pixelMatrix[y][x].R+pixelMatrix[y][x].G+pixelMatrix[y][x].B))/3.0)
		}
		pixelBrightness = append(pixelBrightness, row)
	}

	fmt.Println("\nSuccessfully constructed brightness matrix!")
	fmt.Printf("Brightness matrix size: %d x %d\n", len(pixelBrightness[0]), len(pixelBrightness))
	fmt.Println("Iterating through pixel brightnesses:")

	const asciiChar = "`^\",:;Il!i~+_-?][}{1)(|\\/tfjrxnuvczXYUJCLQ0OZmwqpdbkhao*#MW&8%B@$"

	var pixelASCII [][]string
	for y := 0; y < len(pixelBrightness); y++ {
		var row []string
		for x := 0; x < len(pixelBrightness[0]); x++ {
			row = append(row, string(asciiChar[int(math.Round(pixelBrightness[y][x]*float64(len(asciiChar))/255.0))]))
		}
		pixelASCII = append(pixelASCII, row)
	}

	fmt.Println("\nSuccessfully constructed ASCII matrix!")
	fmt.Printf("ASCII matrix size: %d x %d\n", len(pixelASCII[0]), len(pixelASCII))
	fmt.Println("Iterating through pixel ASCII characters:")

	for k := range pixelASCII {
		fmt.Println(pixelASCII[k])
	}

}

func rgbToPixel(r uint32, g uint32, b uint32, a uint32) pixel {
	return pixel{int(r >> 8), int(g >> 8), int(b >> 8)}

}

type pixel struct {
	R int
	G int
	B int
}
