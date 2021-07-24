package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/sclevine/agouti"
)

const (
	instagramUrl  = "https://www.instagram.com/[targetAccount]/?hl=ja"
	replaceString = "[targetAccount]"

	imageDirectoryPath = "C:\\Users\\kawaguchi_takumi\\workspace\\instascraperStorage\\"
	timeToStringFormat = "20060102150405"
)

func main() {
	fmt.Print("input account name: ")
	var inputAccount string
	fmt.Scan(&inputAccount)

	url := "https://www.instagram.com/" + inputAccount + "/?hl=ja"
	directoryName := time.Now().Format(timeToStringFormat)
	directoryPath := imageDirectoryPath + directoryName
	if err := os.Mkdir(directoryPath, 0777); err != nil {
		fmt.Println(err)
	}
	driver := agouti.ChromeDriver()

	if err := driver.Start(); err != nil {
		fmt.Printf("failed to start driver: %v", err)
	}
	defer driver.Stop()

	page, err := driver.NewPage(agouti.Browser("chrome"))
	if err != nil {
		fmt.Printf("failed to open page: %v", err)
	}

	err = page.Navigate(url)
	if err != nil {
		fmt.Printf("failed to navigate: %v", err)
	}

	contentsDom, err := page.HTML()
	if err != nil {
		fmt.Printf("failed to get html: %v", err)
	}

	reader := strings.NewReader(contentsDom)
	contents, _ := goquery.NewDocumentFromReader(reader)

	contents.Find(".KL4Bh > img").Each(func(i int, s *goquery.Selection) {
		if i <= 3 {
			img, _ := s.Attr("src")
			res, err := http.Get(img)
			if err != nil {
				panic(err)
			}
			defer res.Body.Close()

			fileName := directoryPath + "\\" + inputAccount + strconv.Itoa(i) + ".jpg"
			file, err := os.Create(fileName)
			if err != nil {
				panic(err)
			}
			defer file.Close()

			io.Copy(file, res.Body)
			fmt.Println("download file: ", i)
			time.Sleep(1 * time.Second)
		}
	})
}
