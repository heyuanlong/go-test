package main

import (
	"fmt"
	"os"
	"image"
	"errors"
	"io/ioutil"
	"strings"
	"flag"
	"image/jpeg"
	"image/png"
	"image/draw"
)


func main() {
	dir := flag.String("d", "./tmp/", "d")
	l := flag.Int("l", 0, "l")
	r := flag.Int("r", 0, "r")
	t := flag.Int("t", 0, "t")
	b := flag.Int("b", 0, "b")
	flag.Parse()

	list := make([]string,0)
	listDir(&list,".bmp","./")
	listDir(&list,".png","./")
	listDir(&list,".jpg","./")
	listDir(&list,".jpeg","./")
	os.Mkdir(*dir,os.ModePerm)

	for _, f:= range list {
		fmt.Println(f)
		img , filetype,err := LoadImage(f)
		if err != nil {
			fmt.Println(err)
		}else{
			dst, _ := DealImage(img,*l,*r,*t,*b)
			SaveImage(*dir+f,dst,filetype )
		}
		fmt.Println()
	}

}
func listDir(list *[]string,suffix string,dir string) {
	dirList,err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println("read dir error")
		return
	}
	suffix = strings.ToUpper(suffix)
	for _,v := range dirList {
		if v.IsDir() {
			continue
		}
		if strings.HasSuffix(strings.ToUpper(v.Name()),suffix){
			*list = append(*list,v.Name())
		}
	}
}


func LoadImage(path string) (img image.Image, filetype string, err error) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	img, filetype, err = image.Decode(file)
	return
}
func SaveImage(path string, img *image.RGBA, filetype string) (err error) {
	file, err := os.Create(path)
	if err != nil {
		return
	}
	defer file.Close()
	if filetype == "png" {
		err = png.Encode(file, img)
	} else if filetype == "jpeg" {
		err = jpeg.Encode(file, img, nil)
	} else {
		err = errors.New("载图像后不能识别的类型")
	}
	return
}

func DealImage(img image.Image,left int,right int,top int,bottom int) ( *image.RGBA, error){
	fmt.Println("left:",left)
	fmt.Println("right:",right)
	fmt.Println("top:",top)
	fmt.Println("bottom:",bottom)
	x0 := 0
	y0 := 0
	x1 := img.Bounds().Max.X + left + right
	y1 := img.Bounds().Max.Y + top + bottom
	dst := image.NewRGBA(image.Rect(x0,y0, x1 , y1 ) )

	draw.Draw(dst,dst.Bounds(),img, image.Pt(0 - left,0 - top),draw.Src)
	return dst,nil
}