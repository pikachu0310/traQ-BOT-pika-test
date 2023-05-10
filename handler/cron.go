package handler

import (
	"fmt"
	"log"
	"time"

	"github.com/robfig/cron/v3"

	"example-bot/commands"
)

const trendWChannelID = "5b7b8143-7c0d-4ade-8658-3a8d8ce4dd83"

func Cron() {
	c := cron.New()

	t := time.Now()
	midnight := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
	after := midnight
	before := after.AddDate(0, 0, 1)

	c.AddFunc("39 16 * * *", func() {
		commands.StampsDay(fmt.Sprintf(":w: %d/%d/%d %d/%d/%d 10", after.Year(), after.Month(), after.Day(), before.Year(), before.Month(), before.Day()), trendWChannelID)
	})

	p := cron.NewParser(cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow)
	s, err := p.Parse("50 23 * * *")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(s.Next(time.Now()))
	c.Start()
}
