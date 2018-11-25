package gowechat

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/jpeg"

	"github.com/golang/glog"
)

// 展示登陆二维码
func (w *WeChat) showQrcode() error {
	if w.uuid == "" {
		err := errors.New("haven't get uuid")
		glog.Error("gen qrcode fail", err)
		return err
	}

	uri := qrcodeURI + "/" + w.uuid + "?t=webwx&_=" + timestamp()

	resp, err := w.get(uri)
	if err != nil {
		return err
	}
	img, err := jpeg.Decode(bytes.NewReader([]byte(resp)))
	if err != nil {
		return err
	}

	c, _, err := image.DecodeConfig(bytes.NewReader(resp))
	if err != nil {
		return err
	}
	fmt.Printf("\n")
	for x := 0; x < c.Width; x += 7 {
		for y := 0; y < c.Height; y += 7 {
			r32, g32, b32, _ := img.At(x, y).RGBA()
			r, g, b := int(r32>>8), int(g32>>8), int(b32>>8)
			if (r+g+b)/3 > 180 {
				fmt.Print("\033[47;30m  \033[0m")
			} else {
				fmt.Printf("\033[40;41m  \033[0m")
			}
		}
		fmt.Printf("\n")
	}
	return nil
}
