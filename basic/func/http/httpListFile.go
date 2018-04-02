package main

import (
	"net/http"
	"os"
	"io/ioutil"
)

func errWrapper(writer http.ResponseWriter, request *http.Request) {
	path := request.URL.Path[len("/list/"):]

	file,err := os.Open(path)

	if err != nil {
		http.Error(writer , http.StatusText(http.StatusNotFound) , http.StatusNotFound)
		return
	}

	defer file.Close()

	all , err := ioutil.ReadAll(file)

	if err != nil {
		http.Error(writer , http.StatusText(http.StatusNotFound) , http.StatusNotFound)
		return
	}

	writer.Write(all)
}

func main() {
	http.HandleFunc("/list/" , )

	http.ListenAndServe(":8888" , nil)
}
