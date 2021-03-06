package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		log.Panic(err)
	}

	var t time.Time
	t = time.Date(2021, time.March, 14, 2, 0, 0, 0, loc)
	fmt.Printf("%s\t%s\n", t.Format(time.RFC3339), t.UTC().Format(time.RFC3339))

	t = time.Date(2021, time.November, 7, 1, 0, 0, 0, loc)
	fmt.Printf("%s\t%s\n", t.Format(time.RFC3339), t.UTC().Format(time.RFC3339))
}
