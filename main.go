package main

import (
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetLevel(log.DebugLevel)
}

func main() {
	s := NewServer()

	if err := s.ListenAndServe(); err != nil {
		log.Errorln(err)
	}
}
