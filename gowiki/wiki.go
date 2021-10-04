package main

import (
	"log"
	"net/http"
)

func main() {
	// p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	// p1.Save()

	// p2, _ := LoadPage("TestPage")
	// fmt.Println(string(p2.Body))

	http.HandleFunc("/view/", MakeHandler(ViewHandler))
	http.HandleFunc("/edit/", MakeHandler(EditHandler))
	http.HandleFunc("/save/", MakeHandler(SaveHandler))
	http.HandleFunc("/", FrontPageHandler)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
