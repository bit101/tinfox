// package main is the shit
package main

import (
	"tinpig/config"
	"tinpig/templates"
)

/*
- display project name instead of folder name
- ignore
- cli standard stuff
- colors in config
*/

func main() {
	cfg := config.LoadConfig()
	choice := templates.GetTemplateChoice(cfg)
	template := templates.LoadTemplate(choice, cfg)
	templates.DisplayChoice(template)
	templates.GetProjectDir(template, cfg)
	templates.DefineTokens(template)
	templates.CreateProject(template)
	templates.ShowSuccess(template)
}
