package model

import (
	"log"
	"time"
)

func NewSeed(uid string, seed int) {
	err := RDB.SetNX(uid, seed, 30*time.Minute).Err()
	if err != nil {
		log.Println(err)
	}
}
