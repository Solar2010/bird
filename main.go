package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strconv"
)

func main()  {
	url := "https://movie.douban.com/top250?start="

	var urls []string
	var newUrl string

	fmt.Printf("%v", urls)

	for i := 0; i < 10; i ++ {
		start := i * 25
		newUrl = url + strconv.Itoa(start)

		urls = getTopList(newUrl)

		for _, url := range urls {
			getMovie(url)
		}
	}
}



func getMovie(url string)  {
	fmt.Println(url)

	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("err")
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		panic(err)
	}

	doc.Find("#content h1").Each(func(i int, s *goquery.Selection) {
		fmt.Println("name:" + s.ChildrenFiltered(`[property="v:itemreviewed"]`).Text())

		fmt.Println("year:" + s.ChildrenFiltered(`.year`).Text())
	})

	director := ""

	doc.Find("#info span:nth-child(1) span.attrs").Each(func(i int, s *goquery.Selection) {
		director += s.Text()
	})

	fmt.Println("导演：" + director)

	pl := ""

	doc.Find("#info span:nth-child(3) span.attrs").Each(func(i int, s *goquery.Selection) {
		pl += s.Text()
	})

	fmt.Println("编剧：" + pl)

	charctor := ""

	doc.Find("#info span.actor span.attrs").Each(func(i int, s *goquery.Selection) {
		charctor += s.Text()
	})

	fmt.Println("主演：" + charctor)

	typeStr := ""

	doc.Find("#info > span:nth-child(8)").Each(func(i int, s *goquery.Selection) {
		typeStr += s.Text()
	})

	fmt.Println("类型：" + typeStr )
}

func getTopList(url string) []string  {
	var urls []string
	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	if resp.StatusCode == 200 {
		println("err")
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		panic(err)
	}

	doc.Find("#content div div.article ol li div div.info div.hd a").Each(func(i int, s *goquery.Selection) {
		fmt.Printf("%v", s)
		href, _ := s.Attr("href")
		urls = append(urls, href)
	})

	return urls
}