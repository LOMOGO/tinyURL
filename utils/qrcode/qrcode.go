package qrcode

import (
	qrcode "github.com/skip2/go-qrcode"
	"image/color"
	"strconv"
	"time"
	"tinyURL/global"
	"tinyURL/utils"
)

//用于生成二维码，返回值为图片名称以及出错时的err
func GenerateQrcode(url string) (string, error) {
	filename := utils.MakeSHA1(url+strconv.Itoa(int(time.Now().UnixNano()))) + ".png"
	filepath := global.QRCODECONF.FilePath + filename
	err := qrcode.WriteColorFile(url, qrcode.Highest, global.QRCODECONF.ImgHeight, color.White, color.RGBA{R: 30, G: 144, B: 255, A: 255}, filepath)
	if err != nil {
		return "", err
	}
	return filename, nil
}

//获取二维码的保存位置
func GetQRCodeFullPath() string {
	return global.QRCODECONF.FilePath
}
