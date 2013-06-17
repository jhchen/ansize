package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	_ "image/gif"
	_ "image/png"
	_ "image/jpeg"
	"math/rand"
	"os"
	"strconv"
	"github.com/nfnt/resize"
	"time"
)

const (
	ANSI_BASIC_BASE 	int 		= 16
	ANSI_COLOR_SPACE 	uint32 	= 6
	ANSI_FOREGROUND 	string 	= "38"
	ANSI_RESET 				string 	= "\x1b[0m"
	CHARACTERS 				string 	= "01"
	DEFAULT_WIDTH			int 		= 100
	PROPORTION				float32 = 0.46
	RGBA_COLOR_SPACE 	uint32 	= 1 << 16
)


func toAnsiCode(c color.Color) (string) {
	r, g, b, _ := c.RGBA()
	code := int(ANSI_BASIC_BASE + toAnsiSpace(r) * 36 + toAnsiSpace(g) * 6 + toAnsiSpace(b))
	if code == ANSI_BASIC_BASE {
		return ANSI_RESET
	}
	return "\033[" + ANSI_FOREGROUND + ";5;" + strconv.Itoa(code) + "m"
}

func toAnsiSpace(val uint32) (int) {
	return int(float32(ANSI_COLOR_SPACE) * (float32(val) / float32(RGBA_COLOR_SPACE)))
}

func writeAnsiImage(img image.Image, file *os.File, width int) {
	m := resize.Resize(uint(width), uint(float32(width) * PROPORTION), img, resize.Lanczos3)
	var current, previous string
	bounds := m.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x:= bounds.Min.X; x < bounds.Max.X; x++ {
			current = toAnsiCode(m.At(x, y))
			if (current != previous) {
				fmt.Print(current)
				file.WriteString(current)
			}
			if (ANSI_RESET != current) {
				char := string(CHARACTERS[rand.Int()%len(CHARACTERS)])
				fmt.Print(char)
				file.WriteString(char)
			} else {
				fmt.Print(" ")
				file.WriteString(" ")
			}
		}
		fmt.Print("\n")
		file.WriteString("\n")
	}
	fmt.Print(ANSI_RESET)
	file.WriteString(ANSI_RESET)
}


func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	if len(os.Args) < 3 {
		fmt.Println("Usage: ./ansize <image> <output> [width]>")
		return
	}
	imageName, outputName := os.Args[1], os.Args[2]
	var width int = DEFAULT_WIDTH
	if len(os.Args) >= 4 {
		var err error
		width, err = strconv.Atoi(os.Args[3])
		if err != nil {
			fmt.Println("Invalid width " + os.Args[3] + ". Please enter an integer.")
			return
		}
	}
	imageFile, err := os.Open(imageName)
	if err != nil {
		fmt.Println("Could not open image " + imageName)
		return
  }
  outFile, err := os.Create(outputName)
	if err != nil {
		fmt.Println("Could not open " + outputName + " for writing")
		return
  }
  defer imageFile.Close()
  defer outFile.Close()
  imageReader := bufio.NewReader(imageFile)
  img, _, err := image.Decode(imageReader)
  if err != nil {
  	fmt.Println("Could not decode image")
  	return
  }

  writeAnsiImage(img, outFile, width)
}
