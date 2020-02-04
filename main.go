package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"log"
	"os"
)

var (
	blue = color.RGBA{0, 0, 255, 0xff}
	green = color.RGBA{0,252,0, 0xff}
	red = color.RGBA{255,0,0, 0xff}
	yellow = color.RGBA{255,255,0, 0xff}
)

func AddPhtoFrame() {
	//TODO:user the res/gophergala.jpg to generate a image and write to res/m.jpg which is similar to like the logo in the README.md

	// 1.首先要读取gophergala.jpg获取其数据
	gala := "/Users/hwy/Code/Go/learn-Go-the-hard-way/res/gophergala.jpg"
	f, err := os.Open(gala)
	defer f.Close()
	if err != nil {
		log.Fatal("read error :", err)
		return
	}
	galaImage, err := jpeg.Decode(f)
	if err != nil {
		log.Fatal("decode error :", err)
		return
	}
	// 2.建立新的image
	x_end, y_end := 799, 799
	newImage := image.NewNRGBA(image.Rect(0,0,x_end,y_end))
	// 3.截取gala图片中的人物部分
	//icon的大小 := image.Rect(127, 137, 672, 662)
	icon_sx, icon_sy, icon_ex, icon_ey := 127, 137, 672, 662
	r := image.Rect(icon_sx, icon_sy, icon_ex, icon_ey)
	draw.Draw(newImage, r, galaImage, image.Point{icon_sx, icon_sy}, draw.Src)
	// 4. 给周围填充颜色
	// x轴两边各给127，在一边宽的颜色给110窄的颜色给17 蓝、绿
	// y轴两边各给137，在一边宽的颜色给120窄的颜色给17 红、黄
	x_long, x_short, y_long, y_short := 110, 17, 120, 17
	// x 轴填充 蓝色 绿色
	setColor(newImage, 0, 0, x_long, y_end, blue)
	setColor(newImage, x_long, 0, icon_sx, y_end, green)
	setColor(newImage, icon_ex, 0, icon_ex + x_short, y_end, blue)
	setColor(newImage,  icon_ex + x_short, 0,  x_end, y_end, green)
	// y 轴填充 红色 黄色
	setColor(newImage, icon_sx, 0, icon_ex, y_long, yellow)
	setColor(newImage, icon_sx, y_long, icon_ex, icon_sy, red)
	setColor(newImage, icon_sx, icon_ey, icon_ex, icon_ey+y_short, yellow)
	setColor(newImage, icon_sx, icon_ey+y_short, icon_ex, y_end, red)
	out, _ := os.Create("./res/myicon.jpg")
	defer out.Close()
	jpeg.Encode(out, newImage, &jpeg.Options{100})
}

// 给src的区域（x0,y0）~(x1,y1)填充 c 颜色
func setColor(src *image.NRGBA, x0, y0, x1, y1 int, c color.RGBA) {
	for i := y0; i <= y1; i++ {
		for j := x0; j <= x1; j++ {
			src.Set(j, i, c)
		}
	}
}

func main() {
	AddPhtoFrame()
//	println(`This final exercise,let's add a photo frame for gala logo!
//You should use image package to generate a new iamge from the gala log(which is stored in res/gophergala.jpg,and makes it like the res/m.jpg.
//Now edit main.go to complete 'AddPhtoFrame' function,this task has no test,enjoy your trip!`)
}
