package main
 
import (
	"log"
	"os"
)
 
func main() {
	_, err := os.Stat("test")
 
	if os.IsNotExist(err) { // checks to see whether folder exists
		errDir := os.MkdirAll("test", 0755) // makes folder if none exists
		if errDir != nil {
			log.Fatal(err)
		}
 
	}
}