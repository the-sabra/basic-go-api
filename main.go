package main

import (
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func init() {
	_ = godotenv.Load(".env")
	log.SetLevel(log.DebugLevel)
}

func main() {
	s := NewServer(":" + os.Getenv("PORT"))

	if err := s.ListenAndServe(); err != nil {
		log.Errorln(err)
	}
}

