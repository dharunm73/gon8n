package jd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

var JobDesc string

var Master string

func GetJD() string {

	fmt.Println("Enter the Job Descritpion in the notepad, if understood, press Enter")
	fmt.Scanln()

	file, _ := os.OpenFile("jd.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	file.Close()

	opencmd := exec.Command("notepad.exe", "jd.txt")
	if err := opencmd.Start(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("After saving the file, press Enter here")
	fmt.Scanln()

	data, _ := os.ReadFile("jd.txt")

	JobDesc = string(data)

	return JobDesc
}

func MasterResume() string {

	fmt.Println("Reading master resume json")

	data, err := os.ReadFile("master_resume.json")
	if err != nil {
		log.Fatal(err)
	}
	Master = string(data)

	return Master
}
