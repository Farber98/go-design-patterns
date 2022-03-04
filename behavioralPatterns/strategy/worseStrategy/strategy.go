package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"log"
	"os"
)

/*
STRATEGY:
- Uses different algorithms to achieve some specific functionality. Algorithms are hidden behind an interface and they must be interchangeable.
- All algorithms achieve the same functionality in a different way.

OBJECTIVE:
- Provide a few algorithms to achieve some specific functionality.
- All types achieve the same functionality in a different way but the client of the strategy isn't affected

EXAMPLE: Rendering images or text.
- Two strategies: console and file.
- Key featuer is that the caller doesnt know how the underlying library is working and he just knows the information available on the defined strategy.
- The StrategyInterface will hide ConsoleStrategy and FileStrategy complexity.

ACCEPTANCE CRITERIA:
- Provide a way to show to the user an object (a square) in text or image.
- The user must choose between image or text when launching the app.
- The app must be able to add more visualization strategies (audio, for example)
- If the user selects text, the word Square must be printed in the console.
- If the user selects image, an image of a white square on a black background will be printed on a file.

*/

type PrintStrategy interface {
	Print() error
}

type ConsoleStrategy struct {
}

func (c *ConsoleStrategy) Print() error {
	fmt.Println("Square")
	return nil
}

type FileStrategy struct {
	DestinationFilePath string
}

func (f *FileStrategy) Print() error {
	width := 800
	height := 600

	bgColor := image.Uniform{color.RGBA{R: 70, G: 70, B: 70, A: 0}}
	origin := image.Point{0, 0}
	quality := &jpeg.Options{Quality: 75}

	bgRectangle := image.NewRGBA(image.Rectangle{
		Min: origin,
		Max: image.Point{X: width, Y: height},
	})

	draw.Draw(bgRectangle, bgRectangle.Bounds(), &bgColor, origin, draw.Src)

	squareWidth := 200
	squareHeight := 200
	squareColor := image.Uniform{color.RGBA{R: 255, G: 0, B: 0, A: 1}}
	square := image.Rect(0, 0, squareWidth, squareHeight)
	square = square.Add(image.Point{
		X: (width / 2) - (squareWidth / 2),
		Y: (height / 2) - (squareHeight / 2),
	})
	squareImg := image.NewRGBA(square)

	draw.Draw(bgRectangle, squareImg.Bounds(), &squareColor, origin, draw.Src)

	w, err := os.Create(f.DestinationFilePath)
	if err != nil {
		return fmt.Errorf("Error opening image")
	}
	defer w.Close()

	if err = jpeg.Encode(w, bgRectangle, quality); err != nil {
		return fmt.Errorf("Error writing image to disk")
	}

	return nil
}

var output = flag.String("output", "console", "The output to use between 'console' and 'image' file")

func main() {
	flag.Parse()
	var activeStartegy PrintStrategy
	switch *output {
	case "console":
		activeStartegy = &ConsoleStrategy{}
	case "image":
		activeStartegy = &FileStrategy{"/tmp/image.jpg"}
	default:
		activeStartegy = &ConsoleStrategy{}
	}
	err := activeStartegy.Print()
	if err != nil {
		log.Fatal(err)
	}
}
