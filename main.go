package main

import (
	"flag"
	"fmt"
	"image"
	"image/draw"
	"os"
	"runtime"
	_ "image/jpeg"    //register JPEG decoder
	_ "image/png"
)

var (
	source = flag.String("source", "source.jpg", "-source source.jpg")
	compare = flag.String("compare", "compare.jpg", "-compare compare.jpg")

)

func init() {
	flag.Parse()
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	img1 := loadImage(*source)
	if img1 == nil {
		return
	}
	i1b := calculateImageBrightness(img1) / (img1.Bounds().Size().X + img1.Bounds().Size().Y)
	fmt.Println("Source Image Brightness: ")
	fmt.Println(i1b)

	img2 := loadImage(*compare)
	if img2 == nil {
		return
	}

	i2b := calculateImageBrightness(img2) / (img2.Bounds().Size().X + img2.Bounds().Size().Y)

	fmt.Println("Compare Image Brightness: ")
	fmt.Println(i2b)

	if i1b > i2b {
		fmt.Println("Source image is brightest")
	}
	if i2b > i1b {
		fmt.Println("Compare image is brightest")
	}

}

func calculateImageBrightness(img image.Image) int {
	bounds1 := img.Bounds()

	resultImg := image.NewRGBA(image.Rect(
		bounds1.Min.X,
		bounds1.Min.Y,
		bounds1.Max.X,
		bounds1.Max.Y,
	))
	draw.Draw(resultImg, resultImg.Bounds(), img, image.Point{0, 0}, draw.Src)
	totalBrightness := 0
	for x := bounds1.Min.X; x < bounds1.Max.X; x++ {
		for y := bounds1.Min.Y; y < bounds1.Max.Y; y++ {
			r, g, b, _ := img.At(x, y).RGBA()
			totalBrightness += int(r) + int(g) + int(b)
		}
	}
	return totalBrightness
}


func loadImage(fileName string) image.Image {
	f, err := os.Open(fileName)
	if err != nil {
		// Handle error
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		fmt.Println("Unable to load or read " + fileName + " " + err.Error())
		return nil
	}
	return img
}
