package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/arran4/golang-ical"
)

func main() {
	f, err := os.Open("dates.txt")
	if err != nil {
		log.Fatalf("err opening file: %v", err)
	}
	defer f.Close()

	cal := ics.NewCalendar()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		err := addEventFromLine(cal, line)
		if err != nil {
			log.Fatal("err handling line: %v", err)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("err scanning file:", err)
	}

	fmt.Println(cal.Serialize())
}

func addEventFromLine(cal *ics.Calendar, line string) error {
	d, err := time.Parse("2006-01-02", line)
	if err != nil {
		return fmt.Errorf("err parsing date: %w", err)
	}

	event := cal.AddEvent(fmt.Sprintf("paydate-%v@blachniet.com", line))
	event.SetCreatedTime(time.Now())
	event.SetAllDayStartAt(d)
	event.SetAllDayEndAt(d)
	event.SetSummary("💰 Pay Day")
	return nil
}
