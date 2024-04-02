// package main is the shit
package main

import (
	"tinpig/config"
	"tinpig/templates"
)

/*
- make modules into structs with methods
- cli standard stuff
- colors in config
*/

func main() {
	cfg := config.LoadConfig()
	template := templates.GetTemplateChoice(cfg)
	templates.DisplayChoice(template)
	templates.GetProjectDir(template, cfg)
	templates.DefineTokens(template)
	templates.CreateProject(template)
	templates.ShowSuccess(template)
}
