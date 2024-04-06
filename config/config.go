// Package config has config related functions.
package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/bit101/tinfox/theme"
)

// Config holds the configuration values.
type Config struct {
	TemplatesDir      string `json:"templatesDir"`
	InvalidPathChars  string `json:"invalidPathChars"`
	HeaderColor       string `json:"headerColor"`
	InstructionColor  string `json:"instructionColor"`
	ErrorColor        string `json:"errorColor"`
	DefaultValueColor string `json:"defaultValueColor"`
	Verbose           bool   `json:"verbose"`
	ConfigDir         string `json:"-"`
}

// LoadConfig loads, parses and returns the app configuration.
func LoadConfig() Config {
	configDir, err := os.UserConfigDir()
	checkError(err, "could not find config dir.")
	configPath := filepath.Join(configDir, "tinfox/config")

	initializedConfig := false
	_, err = os.Stat(configPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			initConfig(configDir)
			initializedConfig = true
		}
	}

	configStr, err := os.ReadFile(configPath)
	checkError(err, "could not read config.")

	var configuration Config
	err = json.Unmarshal(configStr, &configuration)
	checkError(err, "could not parse config.")
	if initializedConfig {
		displayConfigSetupMessage(configDir)
	}

	theme.SetTheme(
		configuration.HeaderColor,
		configuration.InstructionColor,
		configuration.ErrorColor,
		configuration.DefaultValueColor,
	)
	return configuration
}

func initConfig(configDir string) {
	err := os.Mkdir(filepath.Join(configDir, "tinfox"), 0775)
	checkError(err, "could not create new config.")

	var cfg Config
	cfg.InvalidPathChars = "‘“!#$%&+^<=>` "
	cfg.TemplatesDir = filepath.Join(configDir, "tinfox", "templates")
	cfg.HeaderColor = "boldgreen"
	cfg.InstructionColor = "yellow"
	cfg.ErrorColor = "boldred"
	cfg.DefaultValueColor = "blue"
	cfg.Verbose = true

	str, err := json.MarshalIndent(cfg, "", "  ")
	checkError(err, "could not create new config.")

	os.WriteFile(filepath.Join(configDir, "tinfox", "config"), str, 0755)
	os.Mkdir(cfg.TemplatesDir, 0755)

	makeSampleTemplate(cfg)
}

func makeSampleTemplate(cfg Config) {
	err := os.Mkdir(filepath.Join(cfg.TemplatesDir, "html"), 0755)
	checkError(err, "could not create sample template.")
	err = os.Mkdir(filepath.Join(cfg.TemplatesDir, "html", "src"), 0755)
	checkError(err, "could not create sample template.")
	err = os.Mkdir(filepath.Join(cfg.TemplatesDir, "html", "styles"), 0755)
	checkError(err, "could not create sample template.")
	err = os.WriteFile(filepath.Join(cfg.TemplatesDir, "html", "index.html"), []byte(htmlTemplate), 0755)
	checkError(err, "could not create sample template.")
	err = os.WriteFile(filepath.Join(cfg.TemplatesDir, "html", "template.json"), []byte(jsonTemplate), 0755)
	checkError(err, "could not create sample template.")
	err = os.WriteFile(filepath.Join(cfg.TemplatesDir, "html", "src", "main.js"), []byte(jsTemplate), 0755)
	checkError(err, "could not create sample template.")
	err = os.WriteFile(filepath.Join(cfg.TemplatesDir, "html", "styles", "main.css"), []byte(cssTemplate), 0755)
	checkError(err, "could not create sample template.")
}

func displayConfigSetupMessage(configDir string) {
	fmt.Println("It looks like this is the first time you're using tinfox.")
	fmt.Printf("We set up a configuration dir at %q.\n", filepath.Join(configDir, "tinfox"))
	fmt.Printf("Add your templates to %q.\n", filepath.Join(configDir, "tinfox", "templates"))
	fmt.Println("You can change this location if you'd like by editing the config file.")
	fmt.Println("We threw in a sample HTML template there to get you started.")
	fmt.Println()
}

func checkError(err error, msg string) {
	if err != nil {
		log.Fatal(msg, err)
	}
}
