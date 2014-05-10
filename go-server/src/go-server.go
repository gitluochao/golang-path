package main

import (
	"net/http"
	"fmt"
	"handler"
	"crypto/tls"
)

//golang adapter //interface  func(ResponseWriter, *Request)
func ErrorHanler(f func(http.ResponseWriter,*http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter,r *http.Request){
		err := f(w,r)
		if err!= nil {
		  http.Error(w,err.Error(),http.StatusInternalServerError)
		}
	}
}
func errorHandler(f func(http.ResponseWriter, *http.Request) error) func(http.ResponseWriter,*http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := f(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func betterHandler(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("better Hanlder")
	return nil
}
func main() {
	conn,err := tls.Dial("tcp","")
	fmt.Println(http.Dir("webapp"))
	http.Handle("/css/",http.FileServer(http.Dir("webapp")))
	http.Handle("/js/",http.FileServer(http.Dir("webapp")))
	http.HandleFunc("/login/",handler.GetLoginHanlder())
	http.HandleFunc("/admin/",handler.GetAdminHanlder())
	http.HandleFunc("/ajax/",handler.GetAjaxHanlder())
	http.HandleFunc("/",errorHandler(betterHandler))  // use adapter
	http.ListenAndServe(":8090",nil)

}
