package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

const (
	instagramUrl  = "https://www.instagram.com/[targetAccount]/?hl=ja"
	targetAccount = ""
	replaceString = "[targetAccount]"

	imageDirectoryPath = ""
	timeToStringFormat = "20060102150405"
)

func main() {
	// doc, err := goquery.NewDocument("https://tabelog.com/")
	// if err != nil {
	// 	panic("Failed to get html.")
	// }
	// contents := doc.Find("h2.rsttop-heading1.rsttop-search__title")
	// contents.Each(func(i int, s *goquery.Selection) {
	// 	fmt.Print(strings.TrimSpace(s.Text()))
	// })

	// 保存先ディレクトリ作成
	directoryName := time.Now().Format(timeToStringFormat)
	directoryPath := imageDirectoryPath + directoryName
	if err := os.Mkdir(directoryPath, 0777); err != nil {
		fmt.Println(err)
	}

	// 画像取得
	imageUrl := "https://scontent-nrt1-1.cdninstagram.com/v/t51.2885-15/sh0.08/e35/s640x640/218653095_183383480434911_6078081123550727209_n.jpg?_nc_ht=scontent-nrt1-1.cdninstagram.com&_nc_cat=1&_nc_ohc=vNlDdILSMasAX-YdwVJ&edm=ABfd0MgBAAAA&ccb=7-4&oh=7042feb4a119a5d644f2b7682a6641c1&oe=60FD2641&_nc_sid=7bff83" // TBD

	// アカウントが存在しない場合には処理を中断する

	// 画像ダウンロード
	response, err := http.Get(imageUrl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	file, err := os.Create(directoryPath + "\\" + targetAccount + ".jpg")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	io.Copy(file, response.Body)
}
