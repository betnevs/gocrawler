package parser

import (
	"regexp"

	"github.com/betNevS/gocrawler/engine"
)

var HouseListRe = regexp.MustCompile(`<a href="([^<>]+)" title="([^<>]+)" class="">[\s\S]*?<td nowrap="nowrap" class="time">([^<>]+)</td>`)
var nexListRe = regexp.MustCompile(`<a href="([^<>]+)" >后页&gt;</a>`)

func HouseList(contents []byte) engine.ParseResult {
	matches := HouseListRe.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		title := string(m[2])
		datetime := string(m[3])
		houseurl := string(m[1])
		result.Items = append(result.Items, title)
		result.Requests = append(result.Requests, engine.Request{Url: string(m[1]), ParserFunc: func(contents []byte) engine.ParseResult {
			return House(datetime, title, houseurl, contents)
		}})
	}
	nextMatch := nexListRe.FindSubmatch(contents)
	result.Requests = append(result.Requests, engine.Request{Url: string(nextMatch[1]), ParserFunc: HouseList})
	return result
}
