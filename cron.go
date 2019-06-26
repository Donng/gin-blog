package main

import (
	"gin-blog/models"
	"github.com/robfig/cron"
	"log"
	"time"
)

func main()  {
	log.Println("Starting...")

	// 创建 cron json runner
	c := cron.New()

	// 加入 schedule 队列
	c.AddFunc("* * * * * *", func() {
		log.Println("Run models.CleanAllTag...")
		models.CleanAllTag()
	})
	c.AddFunc("* * * * * *", func() {
		log.Println("Run models.CleanAllArticle...")
		models.CleanAllArticle()
	})

	// 开启调度程序
	c.Start()

	t1 := time.NewTimer(time.Second * 10)
	for {
		select {
		case <-t1.C:
			t1.Reset(time.Second * 10)
		}
	}
}