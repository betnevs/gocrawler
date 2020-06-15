package engine

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v7"
	"github.com/jinzhu/gorm"

	"github.com/betNevS/gocrawler/models"

	"github.com/betNevS/gocrawler/fetcher"
)

const layout = "2006-01-02 15:04"

type Scheduler interface {
	Submit(Request)
	WorkerChan() chan Request
	ReadyNotifier
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
	Db          *gorm.DB
	RdClient    *redis.Client
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)
	e.Scheduler.Run()
	for i := 0; i < e.WorkerCount; i++ {
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}
	for _, seed := range seeds {
		if e.isDuplicate(seed.Url) {
			log.Printf("Duplicate url : %s", seed.Url)
			continue
		}
		e.Scheduler.Submit(seed)
	}
	houseCount := 0
	for {
		if houseCount != 0 && houseCount%30 == 0 {
			time.Sleep(6 * time.Second)
		}
		result := <-out
		for _, item := range result.Items {
			if h, ok := item.(models.House); ok {
				log.Printf("Got item #%d: %v\n", houseCount, item)
				// 转换模型入库
				houseItem := &models.HouseItem{}
				houseItem.Title = h.Title
				// 转换时间
				tmpT := "2020-" + h.DateTime
				t, _ := time.ParseInLocation(layout, tmpT, time.Local)
				houseItem.PublishTime = t
				// 转换图片
				imgs, _ := json.Marshal(h.Images)
				houseItem.Images = string(imgs)
				houseItem.Url = h.Url
				e.Db.Create(houseItem)
				// 写入重复url的md5
				e.setUrlRecord(h.Url)
				houseCount++
			}
		}
		for _, request := range result.Requests {
			if e.isDuplicate(request.Url) {
				log.Printf("Duplicate url : %s", request.Url)
				continue
			}
			e.Scheduler.Submit(request)
		}
	}
}

func (e *ConcurrentEngine) isDuplicate(url string) bool {
	md5url := Md5Str(url)
	_, err := e.RdClient.Get(md5url).Result()
	if err == nil {
		return true
	} else {
		return false
	}
}

func (e *ConcurrentEngine) setUrlRecord(url string) {
	md5url := Md5Str(url)
	e.RdClient.Set(md5url, 1, 0)
}

func createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {
	go func() {
		for {
			ready.WorkerReady(in)
			req := <-in
			result, err := worker(req)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

func worker(request Request) (ParseResult, error) {
	log.Printf("Use url： %s", request.Url)
	contents, err := fetcher.Fetch(request.Url)
	//fmt.Println(string(contents))
	if err != nil {
		return ParseResult{}, nil
	}
	return request.ParserFunc(contents), nil
}

func Md5Str(str string) string {
	data := []byte(str)
	md5data := md5.Sum(data)
	return fmt.Sprintf("%x", md5data)
}
