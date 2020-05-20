package main

import (
	"bufio"
	"encoding/base64"
	"image"
	"image/png"
	"io"
	"os"
)

//"golang.org/x/tour/pic"

func interesting1(x, y int) int {
	return (x + y) / 2
}

func interesting2(x, y int) int {
	return x * y
}

func interesting3(x, y int) int {
	return x ^ y
}

// Pic is for Drawing
func Pic(dx, dy int) [][]uint8 {
	result := make([][]uint8, dy)
	for i := 0; i < len(result); i++ {
		result[i] = make([]uint8, dx)
		for index := range result[i] {
			result[i][index] = uint8(interesting3(i, index))
		}
	}
	return result
}

func main() {
	Show(Pic)
}

func Show(f func(dx, dy int) [][]uint8) {
	const (
		dx = 256
		dy = 256
	)
	data := f(dx, dy)
	m := image.NewNRGBA(image.Rect(0, 0, dx, dy))
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			v := data[y][x]
			i := y*m.Stride + x*4
			m.Pix[i] = v
			m.Pix[i+1] = v
			m.Pix[i+2] = 255
			m.Pix[i+3] = 255
		}
	}
	ShowImage(m)
}

func ShowImage(m image.Image) {
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	io.WriteString(w, "IMAGE:")
	b64 := base64.NewEncoder(base64.StdEncoding, w)
	err := (&png.Encoder{CompressionLevel: png.BestCompression}).Encode(b64, m)
	if err != nil {
		panic(err)
	}
	b64.Close()
	io.WriteString(w, "\n")
}
