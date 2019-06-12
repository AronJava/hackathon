package main

import "github.com/skip2/go-qrcode"

func main() {
	qrcode.WriteFile("http://188.131.146.10:8089/", qrcode.Medium, 256, "./qr/baidu_qrcode.png")
}
