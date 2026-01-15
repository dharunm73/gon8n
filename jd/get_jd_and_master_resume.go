package jd

import (
	"log"
	"os"
)

var JobDesc string

var Master string

func GetJD() string {
	data, err := os.ReadFile("jd.txt")
	if err != nil {
		log.Fatal(err)
	}
	JobDesc = string(data)
	return JobDesc
}

func MasterResume() string {
	data, err := os.ReadFile("master_resume.json")
	if err != nil {
		log.Fatal(err)
	}
	Master = string(data)

	return Master
}
