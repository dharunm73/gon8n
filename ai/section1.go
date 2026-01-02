package ai

import (
	"fmt"

	"gon8n/jd"

	"context"
	"log"
	"os"

	"google.golang.org/genai"
)

var job string = `Responsibilities:

Design, develop, and maintain backend services using Python and FastAPI
Build scalable, secure, and high-performance REST APIs
Implement complex business logic with clean, testable code
Optimize backend systems using strong DSA and algorithmic thinking
Design and manage database schemas and queries (SQL / NoSQL)
Handle authentication, authorization, and role-based access control
Write unit tests and ensure high code quality
Collaborate with frontend, product, and data teams
Required Skills:

Strong proficiency in Python
Hands-on experience with FastAPI
Excellent Data Structures & Algorithms fundamentals
Strong problem-solving and logical thinking skills
Experience building RESTful APIs
Knowledge of databases (PostgreSQL / MySQL / MongoDB)
Understanding of async programming and API performance optimization
Familiarity with Git and version control workflows`

var Response string

func Gemini() {

	context := context.Background()
	client, err := genai.NewClient(context, nil)
	if err != nil {
		log.Fatal(err)
	}

	config := &genai.GenerateContentConfig{
		SystemInstruction: genai.NewContentFromText("You are a recruiter with years of experience , change the content of the json text i have provided below in each field such as summary content , project content and others according to the jd so that my resume gets the highest ats score possible and Return the response as raw JSON only. Do not wrap it in markdown code blocks. Do not use backticks. and do not say anything except giving the raw json text"+jd.Master, genai.RoleUser),
	}

	result, err := client.Models.GenerateContent(
		context,
		"gemini-2.5-flash",
		genai.Text(job),
		config,
	)
	if err != nil {
		log.Fatal(err)
	}

	Response = result.Text()
	fmt.Println(Response)

}

func TempResume() {
	//Tempdata := []byte(Master)
	err := os.WriteFile("C:/Users/HP/Desktop/Dharun/gon8n/temp_resume.json", []byte(Response), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
