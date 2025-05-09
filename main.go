package main
import (
	"fmt"
	// "strings"
	"io/ioutil"
	"os"
)
func main(){
	conf_content,err := ioutil.ReadFile("example.conf") 
	if err != nil {
		fmt.Println("Unable to open and read configuration file")
		os.Exit(1)
	}
}
