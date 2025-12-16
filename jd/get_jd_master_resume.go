package jd

import "fmt"

var JD string

func GetJD() {
	fmt.Println("Enter the JD")
	fmt.Scanln(&JD)
}
