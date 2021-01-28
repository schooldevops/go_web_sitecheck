package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Noti struct {
	SystemName string //	시스템 이름
	Title      string //	제목
	Contents   string //	내용
	Period     bool   // 	기간 유무
	StartTime  string //	시작시간
	EndTime    string //	종료시간
}

func Index(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))

	err := godotenv.Load("envfile")
	if err != nil {
		log.Fatalf("Error loading envfile. %s", err)
	}

	isShowPeriod, _ := strconv.ParseBool(os.Getenv("PERIOD"))

	data := Noti{
		SystemName: os.Getenv("SITE_NAME"),
		Title:      os.Getenv("TITLE"),
		Contents:   os.Getenv("CONTENTS"),
		Period:     isShowPeriod,
		StartTime:  os.Getenv("START_TIME"),
		EndTime:    os.Getenv("END_TIME"),
	}

	tmpl.Execute(w, data)
}

func main() {
	http.HandleFunc("/", Index)
	fmt.Print("Start Server port on 8888")
	log.Fatal(http.ListenAndServe(":8888", nil))
}
