package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"./mongoDb"
	"./scan"
)

func getAsync(uri string, ch chan<- string) {

	resp, err := http.Get(uri)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	ch <- string(body)
}

func extractFile(r *http.Request) []byte {
	//1. Parse input
	r.ParseMultipartForm(10 << 20)

	//2. retreive file
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving file from form-data")
		fmt.Println(err)
		return nil
	}
	defer file.Close()

	fmt.Printf("Uploaded file: %v\n", handler.Filename)
	fmt.Printf("File Size: %v\n", handler.Size)
	fmt.Printf("MIME Header: %v\n", handler.Header)

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return fileBytes
}

func unsubscribe(w http.ResponseWriter, r *http.Request) {

	//fmt.Fprintf(w, "Uploading file\n")
	file := extractFile(r)

	links := scan.ExtractAllUnsubLinks(string(file))

	ch := make(chan string)

	for _, url := range links {
		go getAsync(url, ch)
	}

	for range links {
		l := mongoDb.UserLink{UserId: 1, UnsubLink: <-ch}
		mongoDb.UserLinkAdd(l)
	}
}

func serveHtml(w http.ResponseWriter, r *http.Request) {
	html, err := ioutil.ReadFile(".\\html\\index.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprintf(w, string(html))

}

func setupRoutes() {
	http.HandleFunc("/unsubscribe", unsubscribe)
	http.HandleFunc("/", serveHtml)
	http.ListenAndServe(":8080", nil)
}

func main() {
	setupRoutes()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
