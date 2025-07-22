package main

import (
	"log"
	"os"

	cronlib "github.com/robfig/cron/v3"
)

func registerJobs(c *cronlib.Cron, app *App) error {
	// 注册所有定时任务
	_, err := c.AddJob(app.Config.Cron.Spec, app.HelloWorldJob)
	if err != nil {
		return err
	}
	// 可以在这里继续注册更多 job
	return nil
}

func main() {
	app, err := createApp()
	if err != nil {
		log.Fatalf("create app failed: %v", err)
	}
	if !app.Config.Cron.Enabled {
		log.Println("cronjob disabled by config")
		os.Exit(0)
	}
	c := cronlib.New()
	if err := registerJobs(c, app); err != nil {
		log.Fatalf("register jobs failed: %v", err)
	}
	log.Printf("cronjob started, spec: %s", app.Config.Cron.Spec)
	c.Run()
}
