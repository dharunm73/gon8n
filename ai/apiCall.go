package ai

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"

	"google.golang.org/genai"
)

var Response string

func Roleplay() string {

	fmt.Println("Reading Instructions")

	data, err := os.ReadFile("Instruction.txt")
	if err != nil {
		log.Fatal(err)
	}
	Instruction := string(data)

	return Instruction
}

func Gemini(JobDesc string, Instruction string, Master string) {

	fmt.Println("API call is running")

	context := context.Background()
	client, err := genai.NewClient(context, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}

	config := &genai.GenerateContentConfig{
		SystemInstruction: genai.NewContentFromText(Instruction+"\n"+Master+"\nBelow is the Job Description\n", genai.RoleUser),
	}

	result, err := client.Models.GenerateContent(
		context,
		"gemini-2.5-flash",
		genai.Text(JobDesc),
		config,
	)
	if err != nil {
		log.Fatal(err)
	}

	Response = result.Text()

}

func UpdateResume() {

	fmt.Println("Resume being loaded onto updated resume json")
	err := os.WriteFile("updated_resume.json", []byte(Response), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func TypstCompilation() {

	fmt.Println("Resume being compiled")
	cmd := exec.Command("typst", "compile", "template.typ", "--font-path", "./")
	_, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
}
