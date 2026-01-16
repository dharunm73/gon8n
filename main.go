package main

import (
	"fmt"
	"gon8n/ai"
	"gon8n/jd"
)

func main() {
	jd.GetJD()
	jd.MasterResume()
	ai.Gemini(jd.JobDesc, ai.Roleplay(), jd.MasterResume())
	ai.UpdateResume()
	ai.TypstCompilation()
	fmt.Println("Everything done")
}
