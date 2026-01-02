package jd

import (
	"fmt"
	"log"
	"os"
)

var JD string

var data byte

var Master string

func GetJD() {
	fmt.Println("Enter the JD")
	fmt.Scanln(&JD)
}

func MasterResume() {
	data, err := os.ReadFile("C:/Users/HP/Desktop/Dharun/gon8n/master_resume.json")
	if err != nil {
		log.Fatal(err)
	}
	Master = string(data)
}

/*func TempResume() {
	//Tempdata := []byte(Master)
	err := os.WriteFile("C:/Users/HP/Desktop/Dharun/gon8n/temp_resume.json", []byte(ai.Response) , 0644)
	if err != nil {
		log.Fatal(err)
	}
}*/
