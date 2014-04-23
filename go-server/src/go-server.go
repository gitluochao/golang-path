package main

import (
	"net/http"
	"fmt"
	"handler"
)

func main() {
	fmt.Println(http.Dir("webapp"))
	http.Handle("/css/",http.FileServer(http.Dir("webapp")))
	http.Handle("/js/",http.FileServer(http.Dir("webapp")))
	http.HandleFunc("/login/",handler.GetLoginHanlder())
	http.HandleFunc("/admin/",handler.GetAdminHanlder())
	http.HandleFunc("/ajax/",handler.GetAjaxHanlder())
	http.HandleFunc("/",handler.GetNotFundHandler())
	http.ListenAndServe(":8090",nil)

}
