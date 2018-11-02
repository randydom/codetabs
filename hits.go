package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	db "./_db"
)

func addHit(service, mode, ip, quest string) {
	layout := "2006-01-02 15:04:05"
	now := time.Now().Format(layout)
	text := fmt.Sprintf("INFO %s %s %s %s\n", now, ip, service, quest)

	// save hit to DB
	if mode == "production" {
		db.MYDB.InsertHit(service, strings.Split(now, " ")[0])
	} else {
		// fmt.Println(`TESTING ... DO NOT DB SAVE`)
	}

	var hitslog = c.App.HitsLog
	hits, err := os.OpenFile(hitslog, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Printf("ERROR opening hits file %s\n", err)
		hits.Close()
		return
	}
	_, err = hits.WriteString(text)
	if err != nil {
		log.Printf("ERROR logging new hit %s\n", err)
		hits.Close()
		return
	}
	defer hits.Close()
}