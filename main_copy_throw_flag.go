package main
 
import (
	"flag"
	"fmt"
	"io/ioutil"
)

	var file string 
	var file2 string 

func copy() {
	
	a, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
	}

	recdata := []byte(a)

	err = ioutil.WriteFile(file2, recdata, 0777)
	if  err != nil {
		fmt.Println(err)
	}
	fmt.Println("Copied")	
} 

 func main() {	
	 flagCopy := flag.String("copy", "yes","copied")
	 	 
	 flag.Parse()

	 fmt.Println(flagCopy)

	 if *flagCopy == "copied" {
		copy()
	 }

 }