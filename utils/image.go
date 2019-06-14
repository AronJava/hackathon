package utils

import (
	"encoding/base64"
	"fmt"
	"log"
	"math/rand"
	"os"
)

var enc = base64.StdEncoding

func GenerateImage(img string, userid string) string {
	randNum := rand.Intn(100000)

	path := ""
	if img[11] == 'j' {
		img = img[23:]
		name := img[10:25]
		path = fmt.Sprintf("/img/%s/%s%d.jpg", userid, name, randNum)
	} else if img[11] == 'p' {
		img = img[22:]
		name := img[10:25]
		path = fmt.Sprintf("/img/%s/%s%d.png", userid, name, randNum)
	} else if img[11] == 'g' {
		img = img[22:]
		name := img[10:25]
		path = fmt.Sprintf("/img/%s/%s%d.png", userid, name, randNum)
	} else {
		log.Println("不支持的文件类型")
		return ""
	}

	data, err := enc.DecodeString(img)
	if err != nil {
		log.Printf("decodestring failed. err:%v", err)
		return ""
	}

	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_EXCL, os.ModePerm)
	if err != nil {
		log.Printf("openfile failed. err:%v", err)
		return ""
	}
	defer f.Close()
	f.Write(data)

	return path
}
