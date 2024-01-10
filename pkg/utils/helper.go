package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

func Debug(obj any) { //any == interface{}
	raw, _ := json.MarshalIndent(obj, "", "\t")
	fmt.Print(string(raw))
}

func LocalTime() time.Time {
	loc, _ := time.LoadLocation("Asia/Bangkok")
	return time.Now().In(loc)
}

func ConvertStringToTime(t string) time.Time {
	layout := "2006-01-02 15:04:05.999 -0700 MST"
	result, err := time.Parse(layout, t)
	if err != nil {
		log.Printf("error: %v", err)
	}

	return result
}
