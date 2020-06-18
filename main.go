package main
 
import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
 
	file, err := os.Open("file.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()
		
	reader := csv.NewReader(file)
	// reader.FieldsPerRecord = 5
	reader.Comment = '#'
	fmt.Println(reader.Read())
 
	for {
		record, e := reader.Read()
		if e != nil {
			fmt.Println(e)
			break
		}
		fmt.Println(record)
	} 
	

}
