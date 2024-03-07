package main

import (
	"fmt"
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

func Pic(dx,dy int) [][]uint8 {
	picture := make([][]uint8, dy)

	for y := 0; y < dy; y++ {
		row := make([]uint8, dx)
		for  x := 0; x < dx; x++ {
			row[x] = uint8((x + y) / 2)
		}
		picture[y] = row
	}
	return picture
}

func main() {
	pic.Show(Pic)
}
