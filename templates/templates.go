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
	"tinpig/clui"
	"tinpig/config"

	"github.com/bit101/go-ansi"
)

// Token describes a single token.
type Token struct {
	Name    string `json:"name"`
	Default string `json:"default"`
}

// Template is a struct holding template data.
type Template struct {
	Name              string   `json:"name"`
	Description       string   `json:"description"`
	Tokens            []Token  `json:"tokens"`
	PostMessage       string   `json:"postMessage"`
	Ignore            []string `json:"ignore"`
	TemplateSourceDir string
	ProjectDir        string
	TokenValues       map[string]string
}

// GetTemplateChoice shows the template ui and returns the choice.
func GetTemplateChoice(cfg config.Config) *Template {
	list := GetTemplateList(cfg)
	if len(list) == 0 {
		ansi.Println(ansi.BoldRed, "No templates found.")
		fmt.Printf("  Add some templates in %q.\n", cfg.TemplatesDir)
		fmt.Println("  Or adjust the `templatesDir` location in the config file.")
		os.Exit(1)
	}
	nameList := []string{}
	for _, template := range list {
		nameList = append(nameList, template.Name)
	}
	index, _ := clui.MultiChoice(nameList, "Choose a project type:")
	return list[index]
}

// GetTemplateList returns the list of available templates
func GetTemplateList(cfg config.Config) []*Template {
	dirList, err := os.ReadDir(cfg.TemplatesDir)
	if err != nil {
		log.Fatal(err)
	}
	list := []*Template{}
	for _, d := range dirList {
		template := LoadTemplate(d.Name(), cfg)
		list = append(list, template)
	}
	return list
}

// LoadTemplate loads, parses and returns the template.
func LoadTemplate(name string, cfg config.Config) *Template {
	templateSourceDir := filepath.Join(cfg.TemplatesDir, name)
	templateStr, err := os.ReadFile(filepath.Join(templateSourceDir, "tinpig.json"))
	if err != nil {
		log.Fatal(err)
	}
	var template Template
	json.Unmarshal(templateStr, &template)
	template.TemplateSourceDir = templateSourceDir
	return &template
}

// DisplayChoice shows info about the template the user has chosen.
func DisplayChoice(template *Template) {
	fmt.Println()
	ansi.Println(ansi.BoldGreen, "Project Info:")
	ansi.Print(ansi.Yellow, "Project: ")
	fmt.Println(template.Name)
	ansi.Print(ansi.Yellow, "Description: ")
	fmt.Println(template.Description)
	fmt.Println()
}

// DefineTokens gets values for all the tokens and stores the values in the template.
func DefineTokens(template *Template) {
	if len(template.Tokens) == 0 {
		return
	}
	ansi.Println(ansi.BoldGreen, "Define values for any tokens:")
	tokenValues := map[string]string{}
	for _, token := range template.Tokens {
		if token.Default == "" {
			value := clui.ReadString(fmt.Sprintf("Value for %q:", token.Name))
			tokenValues[token.Name] = value
		} else {
			value := clui.ReadStringDefault(fmt.Sprintf("Value for %q:", token.Name), token.Default)
			tokenValues[token.Name] = value
		}
	}
	tokenValues["TINPIG_PROJECT_PATH"] = template.ProjectDir
	tokenValues["TINPIG_PROJECT_DIR"] = filepath.Base(template.ProjectDir)
	template.TokenValues = tokenValues
	fmt.Println()
}

// GetProjectDir requests the project directory from the user and stores it in the template.
func GetProjectDir(template *Template, cfg config.Config) {
	var dir string
	ok := false
	for !ok {
		ok = true
		ansi.Println(ansi.BoldGreen, "Project Location: ")
		dir = clui.ReadString("Directory to create project in:")

		// is it an empty string?
		if dir == "" {
			ok = false
			ansi.Println(ansi.BoldRed, "Directory name cannot be empty.")
			fmt.Println()
			continue
		}

		// bad path chars?
		for _, c := range cfg.InvalidPathChars {
			if strings.Index(dir, string(c)) > -1 {
				ok = false
				ansi.Printf(ansi.BoldRed, "Directory name cannot contain %q. Try again.\n\n", c)
				continue
			}
		}

		// let's make sure that's what you wanted
		absDir, _ := filepath.Abs(dir)
		ansi.Print(ansi.Yellow, "You entered: ")
		fmt.Println(absDir)
		confirm := clui.ReadString("Is that correct? [y/n]")
		if strings.ToLower(confirm) != "y" {
			ok = false
			continue
		}

		// does this path already exist?
		_, err := os.Stat(dir)
		if err == nil {
			ansi.Printf(ansi.BoldRed, "Something already exists at location %q. Try again.\n\n", dir)
			ok = false
			continue
		}
	}

	absDir, _ := filepath.Abs(dir)
	template.ProjectDir = absDir
	fmt.Println()
}

// ShowSuccess shows the success message and any post message.
func ShowSuccess(template *Template) {
	ansi.Printf(ansi.BoldGreen, "Success creating the %q project!\n", template.Name)
	ansi.Print(ansi.Yellow, "Location: ")
	fmt.Println(template.ProjectDir)
	if template.PostMessage != "" {
		ansi.Print(ansi.Yellow, "Instructions: ")
		fmt.Println(template.PostMessage)
	}
	fmt.Println()
}

// CreateProject creates the project dir, copies the files and updates the tokens.
func CreateProject(template *Template) {
	templateFiles, err := os.ReadDir(template.TemplateSourceDir)
	if err != nil {
		log.Fatal(err)
	}
	os.Mkdir(template.ProjectDir, 0755)
	for _, file := range templateFiles {
		if file.Name() != "tinpig.json" {
			if !slices.Contains(template.Ignore, file.Name()) {
				copyFile(file, template.TemplateSourceDir, template.ProjectDir, template)
			}
		}
	}
}

func copyFile(file os.DirEntry, srcDir, dstDir string, template *Template) {
	srcFilePath := filepath.Join(srcDir, file.Name())
	dstFilePath := filepath.Join(dstDir, file.Name())
	dstFilePath = replaceDirTokens(dstFilePath, template.TokenValues)

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
			copyFile(subFile, srcFilePath, dstFilePath, template)
		}
	} else {
		fileData, err := os.ReadFile(srcFilePath)
		if err != nil {
			log.Fatal(err)
		}
		fileData = replaceFileTokens(fileData, template.TokenValues)
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
