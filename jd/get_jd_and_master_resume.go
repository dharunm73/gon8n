package jd

import (
	"log"
	"os"
)

var JobDesc string

var data byte

var Master string

func GetJD() {
	//fmt.Println("Enter the Job Description")
	//fmt.Scanln(&JobDesc)

	JobDesc = `Key job responsibilities
• Collaborate with experienced cross-disciplinary Amazonians to conceive, design, and bring innovative products and services to market.
• Design and build innovative technologies in a large distributed computing environment and help lead fundamental changes in the industry.
• Create solutions to run predictions on distributed systems with exposure to innovative technologies at incredible scale and speed.
• Build distributed storage, index, and query systems that are scalable, fault-tolerant, low cost, and easy to manage/use.
• Design and code the right solutions starting with broadly defined problems.
• Work in an agile environment to deliver high-quality software.

Basic Qualifications
- Bachelor's degree or above in computer science, computer engineering, or related field
- Knowledge of Computer Science fundamentals such as object-oriented design, algorithm design, data structures, problem solving, and complexity analysis.
- Knowledge of programming languages such as C/C++, Python, Java or Perl

Preferred Qualifications
- Previous technical internship(s).
- Experience with distributed, multi-tiered systems, algorithms, and relational databases.
- Experience in optimization mathematics such as linear programming and nonlinear optimization.
- Effectively articulate technical challenges and solutions.
- Adept at handling ambiguous or undefined problems as well as ability to think abstractly.`
}

func MasterResume() string {
	data, err := os.ReadFile("C:/Users/HP/Desktop/Dharun/gon8n/master_resume.json")
	if err != nil {
		log.Fatal(err)
	}
	Master = string(data)

	return Master
}
