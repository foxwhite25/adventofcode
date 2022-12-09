package main

import (
	"image"
	"image/color"
	"image/gif"
	"os"
)

func saveToImages(coords []Data) {
	width := 0
	height := 0
	origin := Coordinate{0, 0}
	for _, data := range coords {
		for _, coord := range data.coords {
			if coord.X > width {
				width = coord.X
			}
			if coord.Y > height {
				height = coord.Y
			}
			if coord.X < origin.X {
				origin.X = coord.X
			}
			if coord.Y < origin.Y {
				origin.Y = coord.Y
			}
		}
	}
	var palette = []color.Color{
		color.RGBA{A: 0xff}, color.RGBA{B: 0xff, A: 0xff},
		color.RGBA{G: 0xff, A: 0xff}, color.RGBA{G: 0xff, B: 0xff, A: 0xff},
		color.RGBA{R: 0xff, A: 0xff}, color.RGBA{R: 0xff, B: 0xff, A: 0xff},
		color.RGBA{R: 0xff, G: 0xff, A: 0xff}, color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff},
	}
	var images []*image.Paletted
	for _, c := range coords {
		img := image.NewPaletted(image.Rect(0, 0, (width-origin.X+1)*4, (height-origin.Y+1)*4), palette)
		for _, coordinate := range c.grid {
			img.Set((coordinate.X-origin.X)*4, (coordinate.Y-origin.Y)*4, palette[4])
			img.Set((coordinate.X-origin.X)*4+1, (coordinate.Y-origin.Y)*4, palette[4])
			img.Set((coordinate.X-origin.X)*4, (coordinate.Y-origin.Y)*4+1, palette[4])
			img.Set((coordinate.X-origin.X)*4+1, (coordinate.Y-origin.Y)*4+1, palette[4])
		}

		for _, coord := range c.coords {
			img.Set((coord.X-origin.X)*4, (coord.Y-origin.Y)*4, image.White)
			img.Set((coord.X-origin.X)*4+1, (coord.Y-origin.Y)*4, image.White)
			img.Set((coord.X-origin.X)*4, (coord.Y-origin.Y)*4+1, image.White)
			img.Set((coord.X-origin.X)*4+1, (coord.Y-origin.Y)*4+1, image.White)
		}
		images = append(images, img)
		//f, err := os.Create(fmt.Sprintf("./2022/9/img/%d.png", i))
		//if err != nil {
		//	panic(err)
		//}
		//png.Encode(f, img)
	}

	f, err := os.Create("./2022/9/img/animation.gif")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	err = gif.EncodeAll(f, &gif.GIF{
		Image: images,
		Delay: make([]int, len(images)),
	})
}
