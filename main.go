package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
)

var config_path string

const default_target_filename = ".pin"

func parse_config() map[string]string {
	os_name := runtime.GOOS
	if os_name == "windows" {
		config_path = os.Getenv("USERPROFILE") + "\\pin\\aliases.conf"
	} else if os_name == "linux" || os_name == "darwin" || strings.Contains(os_name, "bsd") || os_name == "dragonfly" {
		config_path = os.Getenv("HOME") + "/.config/pin/aliases.conf"
	} else {
		fmt.Println("Your OS might be not supported")
		os.Exit(1)
	}

	config_content_bytes, err := ioutil.ReadFile(config_path)
	if err != nil {
		if os.IsNotExist(err) {
			err := os.MkdirAll(filepath.Dir(config_path), 0755)
			if err != nil {
				fmt.Println("Falied to create configuration file in ", config_path)
				fmt.Println("Please make it manually before starting this utility again.")
				os.Exit(1)
			}
			_, err = os.Create(config_path)
			if err != nil {
				fmt.Println("Falied to create configuration file in ", config_path)
				fmt.Println("Please make it manually before starting this utility again.")
				os.Exit(1)
			}

			fmt.Println("A new configuration file has been created at:", config_path)
			fmt.Println("To get started, please add your aliases to this file, following the format described in the GitHub repository:")
			fmt.Println("https://github.com/XodTech/pin")
		} else {
			fmt.Println("Unable to open and read configuration file")
			os.Exit(1)
		}
	}
	re := regexp.MustCompile(`\s+`)
	config_content := re.ReplaceAllString(string(config_content_bytes), "")
	config := make(map[string]string)
	for _, pair := range strings.Split(string(config_content), `;`) {
		kv := strings.Split(pair, `=`)
		if len(kv) == 2 {
			config[kv[0]] = kv[1]
		} else if len(kv) == 1 {
			break
		} else {
			fmt.Printf("Invalid pair %s\n", pair)
		}
	}
	return config
}
func main() {
	var target_filename string
	if len(os.Args) >= 2 {
		config := parse_config()
		option := os.Args[1]
		if config[option] != "" {
			target_filename = config[option]
		} else {
			target_filename = default_target_filename
		}

	} else if len(os.Args) == 1 {
		target_filename = default_target_filename
	}
	files, err := ioutil.ReadDir(".")
	if err != nil {
		fmt.Println("Unable to get list of files in the current directory")
		os.Exit(1)
	}
	for _, file := range files {
		if file.Name() == target_filename {
			content, err := ioutil.ReadFile(target_filename)
			if err != nil {
				fmt.Println("Unable to read the target file")
				os.Exit(1)
			} else if len(content) >= 1 {
				content = append(content, "\n"...)
			}
			content = append(content, []byte(os.Args[2])...)
			err = ioutil.WriteFile(target_filename, content, 0644)
			if err != nil {
				fmt.Println("Unable to write to the target file")
				os.Exit(1)
			}
			return
		}
	}
	err = ioutil.WriteFile(target_filename, []byte(os.Args[2]), 0644)
	if err != nil {
		fmt.Println("Unable to create and write to the target file")
		os.Exit(1)
	}
}
