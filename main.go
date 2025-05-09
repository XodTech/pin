package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"os"
)
func main(){
	conf_content,err := ioutil.ReadFile("example.conf") 
	if err != nil {
		fmt.Println("Unable to open and read configuration file")
		os.Exit(1)
	}
	config := make(map[string]string)
	for _,pair := range strings.Split(string(conf_content),`;`){
		kv := strings.Split(pair,`=`)
		if len(kv) == 2 {
			config[kv[0]] = kv[1]
		}else if len(kv) == 1{
			break
		}else {
			fmt.Printf("Invalid pair %s\n",pair)
		}
	}
}
