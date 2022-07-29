package main

import (
	"fmt"
	"log"

	"github.com/eaigner/jet"
	"github.com/lib/pq"
)

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	fmt.Println("Hello World")
	psqlURL := "postgres://ebddjabf:u2RaNK1qawGzXy-xatygoQAI94cdkD2q@fanny.db.elephantsql.com/ebddjabf"
	pgUrl, err := pq.ParseURL(psqlURL)
	logFatal(err)

	db, err := jet.Open("postgres", pgUrl)
	logFatal(err)

	var shows []*struct {
		Id       int
		Showid   string
		Showjson string
	}

	err = db.Query("SELECT * FROM shows").Rows(&shows)
	logFatal(err)

	for _, show := range shows {
		log.Printf("\nid: %v,\nshowid: %s,\nshowjson: %s\n\n",
			show.Id,
			show.Showid,
			show.Showjson)
	}
}
