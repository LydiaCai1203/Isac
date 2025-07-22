package cronjob

import (
	"log"

	"framework/repo"
)

type HelloWorldJob struct {
	UserRepo *repo.UserRepo
}

func NewHelloWorldJob(userRepo *repo.UserRepo) *HelloWorldJob {
	return &HelloWorldJob{UserRepo: userRepo}
}

func (j *HelloWorldJob) Run() {
	count := j.UserRepo.CountUsers()
	log.Printf("Helloworld, user count: %d", count)
}
