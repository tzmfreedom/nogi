package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {

	members := [][]string{}
	doc, _ := goquery.NewDocumentFromResponse(GetPage("http://www.nogizaka46.com/member/"))
	doc.Find(".unit a").Each(func(i int, s *goquery.Selection) {
		member := []string{}
		member = append(member, s.Find(".main").Text())
		member = append(member, s.Find(".sub").Text())

		href, _ := s.Attr("href")
		doc, _ := goquery.NewDocumentFromResponse(GetPage("http://www.nogizaka46.com/member/" + href))
		doc.Find("#profile dd").Each(func(i int, s *goquery.Selection) {
			member = append(member, s.Text())
		})
		members = append(members, member)
	})
	for _, member := range members {
		fmt.Println("[[members]]")
		fmt.Println("name = \"" + member[0] + "\"")
		fmt.Println("hiragana = \"" + member[1] + "\"")
		fmt.Println("birthday = \"" + member[2] + "\"")
		fmt.Println("blood_type = \"" + member[3] + "\"")
		fmt.Println("constellation = \"" + member[4] + "\"")
		fmt.Println("height = \"" + member[5] + "\"")
		fmt.Println("")
	}
}

func GetPage(url string) *http.Response {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.1)")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	return resp
}
