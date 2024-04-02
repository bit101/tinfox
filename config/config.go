// Package config has config related functions.
package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/bit101/go-ansi"
)

// Config holds the configuration values.
type Config struct {
	TemplatesDir     string `json:"templatesDir"`
	InvalidPathChars string `json:"invalidPathChars"`
}

// NewConfig creates a new, mostly empty config.
func NewConfig() Config {
	return Config{
		InvalidPathChars: "‘“!#$%&+^<=>` ",
	}
}

// LoadConfig loads, parses and returns the app configuration.
func LoadConfig() Config {
	configDir, err := os.UserConfigDir()
	if err != nil {
		log.Fatal("could not load config")
	}
	configPath := filepath.Join(configDir, "tinpig/config")

	_, err = os.Stat(configPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			makeConfig(configDir)
		}
	}

	configStr, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatal("could not load config")
	}
	var configuration Config
	err = json.Unmarshal(configStr, &configuration)
	if err != nil {
		log.Fatal("could not load config")
	}
	return configuration
}

func makeConfig(configDir string) {
	os.Mkdir(filepath.Join(configDir, "tinpig"), 0775)
	cfg := NewConfig()
	cfg.TemplatesDir = filepath.Join(configDir, "tinpig", "templates")
	str, err := json.Marshal(cfg)
	if err != nil {
		log.Fatal(err)
	}
	os.WriteFile(filepath.Join(configDir, "tinpig", "config"), str, 0755)
	os.Mkdir(cfg.TemplatesDir, 0755)
	makeSampleTemplate(cfg)
	ansi.Println(ansi.BoldGreen, "Config Setup")
	ansi.Println(ansi.Yellow, "It looks like this is the first time you're using the app.")
	ansi.Printf(ansi.Yellow, "We set up a configuration dir at %q.\n", filepath.Join(configDir, "tinpig"))
	ansi.Println(ansi.Yellow, "This is also where you can store your templates. We created a sample template to get you started.")
	ansi.Println(ansi.Yellow, "You can change this location if you'd like by editing the config file.")
	fmt.Println()
}

func makeSampleTemplate(cfg Config) {
	err := os.Mkdir(filepath.Join(cfg.TemplatesDir, "html"), 0755)
	if err != nil {
		log.Fatal(err)
	}
	err = os.Mkdir(filepath.Join(cfg.TemplatesDir, "html", "src"), 0755)
	if err != nil {
		log.Fatal(err)
	}
	err = os.Mkdir(filepath.Join(cfg.TemplatesDir, "html", "styles"), 0755)
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile(filepath.Join(cfg.TemplatesDir, "html", "index.html"), []byte(htmlTemplate), 0755)
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile(filepath.Join(cfg.TemplatesDir, "html", "tinpig.json"), []byte(jsonTemplate), 0755)
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile(filepath.Join(cfg.TemplatesDir, "html", "src", "main.js"), []byte(jsTemplate), 0755)
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile(filepath.Join(cfg.TemplatesDir, "html", "styles", "main.css"), []byte(cssTemplate), 0755)
	if err != nil {
		log.Fatal(err)
	}
}
