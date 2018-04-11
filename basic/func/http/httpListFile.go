package main

import (
	"net/http"
	"os"
	"github.com/zhongliangshan/test/basic/func/http/handlelist"
	"log"
)

type appHandle func(writer http.ResponseWriter, request *http.Request) error

func errWrapper(handle appHandle) func(http.ResponseWriter,*http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {

		err := handle(writer , request)

		if err != nil {
			log.Printf("Error occurred "+
				"handling request: %s",
				err.Error())
			if userErr , ok  := err.(userError);ok{
				http.Error(writer,
					userErr.Message(),
					http.StatusBadRequest)
				return
			}
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer , http.StatusText(code) , code)
		}
	}
}

type userError interface {
	error
	Message() string
}

func main() {
	http.HandleFunc("/" , errWrapper(handlelist.Handlefunc))

	err := http.ListenAndServe(":8888" , nil)
	if err != nil {
		panic(err)
	}
}
