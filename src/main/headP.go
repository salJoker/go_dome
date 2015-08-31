package main

import(
	"os"
	"image/color"
	"image/png"
	"github.com/issue9/identicon"
)

func main(){
//	img, _ := identicon.Make(128, color.NRGBA{},color.NRGBA{}, []byte("huangchao387@sina.com"))
	ii, _ := identicon.New(128, color.YCbCr{}, color.Gray{}, color.Gray16{}, color.NRGBA{})
	img := ii.Make([]byte("abcdefg"))
	fi, _ := os.Create("resources/u1.png")
	png.Encode(fi, img)
	fi.Close()
}