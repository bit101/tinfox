// Package templates has file related functions.
package templates

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/bit101/tinfox/clui"
	"github.com/bit101/tinfox/config"
	"github.com/bit101/tinfox/theme"
)

// Token describes a single token.
type Token struct {
	Name       string `json:"name"`
	Default    string `json:"default"`
	IsPath     bool   `json:"isPath"`
	IsRequired bool   `json:"required"`
}

// Template is a struct holding template data.
type Template struct {
	Name              string   `json:"name"`
	Description       string   `json:"description"`
	Tokens            []Token  `json:"tokens"`
	PreMessage        string   `json:"preMessage"`
	PostMessage       string   `json:"postMessage"`
	Ignore            []string `json:"ignore"`
	TemplateSourceDir string
	ProjectDir        string
	TokenValues       map[string]string
}

// TemplateParser reads and parses a template.
type TemplateParser struct {
	template *Template
}

// NewTemplateParser creates a new TemplateParser.
func NewTemplateParser() *TemplateParser {
	config.LoadConfig()
	return &TemplateParser{}
}

// LoadAndParse loads the template list, gets the user's choice, dir, tokens values and creates the project.
func (t *TemplateParser) LoadAndParse() {
	t.GetTemplateChoice()
	t.GetProjectDir()
	t.DefineTokens()
	t.CreateProject()
	t.ShowSuccess()
}

// GetTemplateChoice shows the template ui and returns the choice.
func (t *TemplateParser) GetTemplateChoice() {
	list := t.GetTemplateList()
	if len(list) == 0 {
		theme.PrintErrorln("No templates found.")
		fmt.Printf("  Add some templates in %q.\n", config.ActiveConfig.TemplatesDir)
		fmt.Println("  Or adjust the `templatesDir` location in the config file.")
		os.Exit(1)
	}
	nameList := []string{}
	for _, template := range list {
		nameList = append(nameList, template.Name)
	}
	index, _ := clui.MultiChoice(nameList, "Choose a project type:")
	t.template = list[index]
	t.DisplayChoice()
}

// DisplayList displays the list of available templates.
func (t *TemplateParser) DisplayList() {
	list := t.GetTemplateList()
	for _, item := range list {
		theme.PrintInstructionf("%s\n", item.Name)
		fmt.Printf("  %s\n", item.Description)
	}
}

// DisplayChoice shows info about the template the user has chosen.
func (t *TemplateParser) DisplayChoice() {
	if config.ActiveConfig.Verbose {
		fmt.Println()
		theme.PrintHeaderln("Project Info:")
		theme.PrintInstruction("Project: ")
		fmt.Println(t.template.Name)
		theme.PrintInstruction("Description: ")
		fmt.Println(t.template.Description)
	}
	if t.template.PreMessage != "" {
		theme.PrintInstruction("Message: ")
		fmt.Println(t.template.PreMessage)
	}
	fmt.Println()
}

// GetTemplateList returns the list of available templates
func (t *TemplateParser) GetTemplateList() []*Template {
	dirList, err := os.ReadDir(config.ActiveConfig.TemplatesDir)
	if err != nil {
		log.Fatal(err)
	}
	list := []*Template{}
	for _, d := range dirList {
		template, err := t.LoadTemplate(d.Name())
		if err == nil {
			list = append(list, template)
		}
	}
	return list
}

// LoadTemplate loads, parses and returns the template.
func (t *TemplateParser) LoadTemplate(name string) (*Template, error) {
	templateSourceDir := filepath.Join(config.ActiveConfig.TemplatesDir, name)
	templateStr, err := os.ReadFile(filepath.Join(templateSourceDir, "template.json"))
	if err != nil {
		return nil, err
	}
	var template Template
	json.Unmarshal(templateStr, &template)
	template.TemplateSourceDir = templateSourceDir
	return &template, nil
}

// DefineTokens gets values for all the tokens and stores the values in the template.
func (t *TemplateParser) DefineTokens() {
	if len(t.template.Tokens) == 0 {
		return
	}
	if config.ActiveConfig.Verbose {
		theme.PrintHeaderln("Define values for any tokens:")
	}
	tokenValues := map[string]string{}
	for _, token := range t.template.Tokens {
		value := clui.ReadToken(token.Name, token.Default, token.IsRequired, token.IsPath)
		tokenValues[token.Name] = value
	}
	tokenValues["PROJECT_PATH"] = t.template.ProjectDir
	tokenValues["PROJECT_DIR"] = filepath.Base(t.template.ProjectDir)
	t.template.TokenValues = tokenValues
	fmt.Println()
}

// GetProjectDir requests the project directory from the user and stores it in the template.
func (t *TemplateParser) GetProjectDir() {
	var dir string
	ok := false
	for !ok {
		ok = true
		if config.ActiveConfig.Verbose {
			theme.PrintHeaderln("Project Location: ")
		}
		dir = clui.ReadString("Directory to create project in:")

		// is it an empty string?
		if dir == "" {
			ok = false
			theme.PrintErrorln("Directory name cannot be empty.")
			fmt.Println()
			continue
		}

		// bad path chars?
		for _, c := range config.ActiveConfig.InvalidPathChars {
			if strings.Index(dir, string(c)) > -1 {
				ok = false
				theme.PrintErrorf("Directory name cannot contain %q. Try again.\n\n", c)
			}
		}
		if !ok {
			continue
		}

		if config.ActiveConfig.Verbose {
			// let's make sure that's what you wanted
			absDir, _ := filepath.Abs(dir)
			theme.PrintInstruction("You entered: ")
			fmt.Println(absDir)
			confirm := clui.ReadString("Is that correct? [Y/n]")
			confirm = strings.ToLower(confirm)
			if confirm != "" && confirm != "y" {
				ok = false
				theme.PrintErrorln("OK, let's try again.")
				continue
			}
		}

		// does this path already exist?
		_, err := os.Stat(dir)
		if err == nil {
			theme.PrintErrorf("Something already exists at location %q. Try again.\n\n", dir)
			ok = false
			continue
		}
	}

	absDir, _ := filepath.Abs(dir)
	t.template.ProjectDir = absDir
	fmt.Println()
}

// ShowSuccess shows the success message and any post message.
func (t *TemplateParser) ShowSuccess() {
	theme.PrintHeaderf("Success creating the %q project!\n", t.template.Name)
	theme.PrintInstruction("Location: ")
	fmt.Println(t.template.ProjectDir)
	if t.template.PostMessage != "" {
		theme.PrintInstruction("Instructions: ")
		fmt.Println(t.template.PostMessage)
	}
	fmt.Println()
}

// CreateProject creates the project dir, copies the files and updates the tokens.
func (t *TemplateParser) CreateProject() {
	templateFiles, err := os.ReadDir(t.template.TemplateSourceDir)
	if err != nil {
		log.Fatal(err)
	}
	os.Mkdir(t.template.ProjectDir, 0755)
	for _, file := range templateFiles {
		if file.Name() != "template.json" {
			if !slices.Contains(t.template.Ignore, file.Name()) {
				t.copyFile(file, t.template.TemplateSourceDir, t.template.ProjectDir)
			}
		}
	}
}

func (t *TemplateParser) copyFile(file os.DirEntry, srcDir, dstDir string) {
	srcFilePath := filepath.Join(srcDir, file.Name())
	dstFilePath := filepath.Join(dstDir, file.Name())
	dstFilePath = replaceDirTokens(dstFilePath, t.template.TokenValues)

	fileInfo, err := file.Info()
	if err != nil {
		log.Fatal(err)
	}
	mode := fileInfo.Mode()

	if file.IsDir() {
		os.Mkdir(dstFilePath, mode)
		files, err := os.ReadDir(srcFilePath)
		if err != nil {
			log.Fatal(err)
		}
		for _, subFile := range files {
			t.copyFile(subFile, srcFilePath, dstFilePath)
		}
	} else {
		fileData, err := os.ReadFile(srcFilePath)
		if err != nil {
			log.Fatal(err)
		}
		fileData = replaceFileTokens(fileData, t.template.TokenValues)
		os.WriteFile(dstFilePath, fileData, mode)
	}
}

func replaceFileTokens(fileData []byte, tokens map[string]string) []byte {
	text := string(fileData)
	for token, value := range tokens {
		token = "${" + token + "}"
		text = strings.ReplaceAll(text, token, value)
	}
	return []byte(text)
}

func replaceDirTokens(path string, tokens map[string]string) string {
	for token, value := range tokens {
		token = "%" + token + "%"
		path = strings.ReplaceAll(path, token, value)
	}
	return path
}
