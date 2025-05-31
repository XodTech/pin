package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"os"
	"regexp"
)

const config_path = "example.conf"
const default_target_filename = ".pin"

func parse_config() map[string]string {
	config_content_bytes,err := ioutil.ReadFile(config_path) 
	if err != nil {
		fmt.Println("Unable to open and read configuration file")
		os.Exit(1)
	}
	re := regexp.MustCompile(`\s+`)
	config_content := re.ReplaceAllString(string(config_content_bytes),"")
	config := make(map[string]string)
	for _,pair := range strings.Split(string(config_content),`;`){
		kv := strings.Split(pair,`=`)
		if len(kv) == 2 {
			config[kv[0]] = kv[1]
		}else if len(kv) == 1{
			break
		}else {
			fmt.Printf("Invalid pair %s\n",pair)
		}
	}
	return config
}
func main(){	
	var target_filename string
	if len(os.Args) >= 2{
		config := parse_config()
		arg := os.Args[1]
		if config[arg] != "" {
			target_filename = config[arg]
		}else {
			target_filename = default_target_filename
		}

	}else if len(os.Args) == 1 {
		target_filename = default_target_filename
	}
	files, err := ioutil.ReadDir(".")
	if err != nil {
		fmt.Println("Unable to get list of files in the current directory")
		os.Exit(1)
	}
	for _,file := range files {
		if file.Name() == target_filename{
			err := ioutil.WriteFile(target_filename,[]byte(os.Args[2]),0644)
			if err != nil {		
				fmt.Println("Unable to write to the target file")
				os.Exit(1)
			}
		}
	}

}
