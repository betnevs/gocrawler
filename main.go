package main

import (
	"fmt"
	"log"
	"time"

	"github.com/betNevS/gocrawler/douban/parser"
	"github.com/betNevS/gocrawler/engine"
	"github.com/betNevS/gocrawler/scheduler"
	"github.com/go-redis/redis/v7"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	startUrl := "https://www.douban.com/group/gz_rent/discussion?start=0"
	rd := redis.NewClient(&redis.Options{Addr: "127.0.0.1:6379", Password: "", DB: 0})
	db, err := gorm.Open("mysql", "root:666666@tcp(127.0.0.1:3306)/crawler?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("Open mysql failed : ", err)
		log.Fatal("Open mysql failed : ", err)
	}
	db.DB().SetMaxOpenConns(100)
	db.DB().SetMaxIdleConns(20)
	db.DB().SetConnMaxLifetime(4 * time.Hour)
	db.SingularTable(true)
	e := &engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 10,
		RdClient:    rd,
		Db:          db,
	}
	e.Run(engine.Request{Url: startUrl, ParserFunc: parser.HouseList})
}
