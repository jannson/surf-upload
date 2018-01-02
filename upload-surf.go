package main

import (
	"fmt"
	"log"
	"os"

	"github.com/headzoo/surf"
)

func postFile(filename string, targetUrl string) error {
	bow := surf.NewBrowser()
	err := bow.Open(targetUrl)
	if err != nil {
		log.Fatal(err)
	}

	//log.Println("bow.Dom()", bow.Dom())

	fh, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file")
		return err
	}
	defer fh.Close()

	forms := bow.Forms()
	log.Println("forms len=", len(forms))
	form := forms[1]
	form.SetFile("uploadfile", filename, fh)

	err = form.Submit()
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func main() {
	target_url := "http://localhost:9000/upload"
	filename := "LICENSE"
	postFile(filename, target_url)
}
