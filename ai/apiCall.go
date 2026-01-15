package ai

import (
	"context"
	"log"
	"os"
	"os/exec"

	"google.golang.org/genai"
)

var Response string

func Roleplay() string {

	data, err := os.ReadFile("Instruction.txt")
	if err != nil {
		log.Fatal(err)
	}
	Instruction := string(data)

	return Instruction
}

func Gemini(JobDesc string, Instruction string, Master string) {

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
	err := os.WriteFile("updated_resume.json", []byte(Response), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func TypstCompilation() {
	cmd := exec.Command("typst", "compile", "template.typ", "--font-path", "./")
	_, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
}
