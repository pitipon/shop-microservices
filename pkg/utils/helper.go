package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

func Debug(object any) {
	raw, _ := json.MarshalIndent(object, "", "\t")
	fmt.Println(string(raw))

}

func LocalTime() time.Time {
	loc, _ := time.LoadLocation("Asia/Bangkok")
	return time.Now().In(loc)
}

func ConvertStringTimeToTime(t string) time.Time {
	layout := "2006-01-02 15:04:05.999 -0700 MST"
	result, err := time.Parse(layout, t)
	if err != nil {
		log.Printf("Error: Parse time failed: %s", err.Error())
	}
	return result
}
