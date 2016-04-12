package console

import (
	"encoding/json"
	"log"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func JSON(input ...interface{}) {
	for _, item := range input {
		data, err := json.MarshalIndent(item, "", "  ")
		if err != nil {
			continue
		}
		log.Println(string(data))
	}
}

func Print(v ...interface{}) {
	log.Print(v...)
}

func Println(v ...interface{}) {
	log.Println(v...)
}

func Printf(format string, v ...interface{}) {
	log.Printf(format, v...)
}

func Panic(v ...interface{}) {
	log.Fatal(v...)
}
