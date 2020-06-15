package parser

import (
	"regexp"

	"github.com/betNevS/gocrawler/models"

	"github.com/betNevS/gocrawler/engine"
)

//var authorRe = regexp.MustCompile(`<span class="from">[\s\S]+<a href="([^<>]+)">([^<>]+)</a></span>`)
//var dateRe = regexp.MustCompile(`<span class="color-green" style="display:inline-block">([^<>]+)</span>`)
var imgRe = regexp.MustCompile(`<img src="([^<>]+)" width="500"/>`)

func House(datetime string, title string, houseurl string, contents []byte) engine.ParseResult {
	house := models.House{}
	house.Title = title
	house.DateTime = datetime
	house.Url = houseurl
	//authorMatch := authorRe.FindSubmatch(contents)
	//house.Author = string(authorMatch[2])
	//dateMatch := dateRe.FindSubmatch(contents)
	//house.DateTime = string(dateMatch[1])
	ImgMatches := imgRe.FindAllSubmatch(contents, -1)
	for _, img := range ImgMatches {
		house.Images = append(house.Images, string(img[1]))
	}
	result := engine.ParseResult{}
	result.Items = []interface{}{house}
	return result
}
