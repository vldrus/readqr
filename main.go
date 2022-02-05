package main

import (
	"bytes"
	"fmt"
	"image"
	"os"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	_ "golang.org/x/image/bmp"
	_ "golang.org/x/image/tiff"
	_ "golang.org/x/image/vp8"
	_ "golang.org/x/image/vp8l"
	_ "golang.org/x/image/webp"

	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Printf("Usage: %s image1.png image2.jpg ...\n", args[0])
		return
	}

	for i := 1; i < len(args); i++ {
		path := args[i]

		data, err := os.ReadFile(path)
		if err != nil {
			fmt.Printf("%s: Cannot read file: %v\n", path, err)
			continue
		}

		img, _, err := image.Decode(bytes.NewReader(data))
		if err != nil {
			fmt.Printf("%s: Cannot decode image: %v\n", path, err)
			continue
		}

		bmp, _ := gozxing.NewBinaryBitmapFromImage(img)

		res, err := qrcode.NewQRCodeReader().Decode(bmp, nil)
		if err != nil {
			fmt.Printf("%s: Cannot read QR Code from image: %v\n", path, err)
			continue
		}

		fmt.Printf("%s: %s\n", path, res.String())
	}
}
