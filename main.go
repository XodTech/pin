package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"os"
)
const conf_path = "example.conf"
const default_filename = ".pin"
func parse_config() map[string]string {
	conf_content,err := ioutil.ReadFile(conf_path) 
	if err != nil {
		fmt.Println("Unable to open and read configuration file")
		os.Exit(1)
	}
	config := make(map[string]string)
	for _,pair := range strings.Split(string(conf_content),`;`){ // TODO: Add function to remove spaces
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
	var filename string
	if len(os.Args) >= 2{
		config := parse_config()
		arg := os.Args[1]
		if config[arg] != "" {
			filename = config[arg]
		}else {
			filename = default_filename
		}

	}else if len(os.Args) == 1 {
		filename = default_filename
	}
}
