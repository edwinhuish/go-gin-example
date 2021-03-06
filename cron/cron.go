package main

import (
	"log"
	"time"

	"github.com/robfig/cron"

	"github.com/edwinhuish/go-gin-example/models"
)

func main() {
	log.Println("Starting...")

	c := cron.New()
	c.AddFunc("* * * * * *", func() {
		log.Println("Run models.CleanAllTag...")
		models.CleanAllTag()
	})
	c.AddFunc("* * * * * *", func() {
		log.Println("Run models.CleanAllArticle...")
		models.CleanAllArticle()
	})

	c.Start()

	t1 := time.NewTimer(time.Second * 10)
	for {
		<-t1.C
		t1.Reset(time.Second * 10)
	}
}
