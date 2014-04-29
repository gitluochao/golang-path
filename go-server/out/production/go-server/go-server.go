package main

import (
	"fmt"
	"handler"
	"net/http"
	"os/exec"
	"reflect"
	"util"
)

var logger = &util.LoggerHelp{}

func main() {
	cmd := exec.Command("pwd", []string{})
	cmd.Start()
	fmt.Println(http.Dir("webapp"))
	logger.Info("reflect.Value")
	fmt.Println(reflect.ValueOf(logger).MethodByName("Info").Call([]reflect.Value{reflect.ValueOf("message")}))
	http.Handle("/css/", http.FileServer(http.Dir("webapp")))
	http.Handle("/js/", http.FileServer(http.Dir("webapp")))
	http.HandleFunc("/login/", handler.GetLoginHanlder())
	http.HandleFunc("/admin/", handler.GetAdminHanlder())
	http.HandleFunc("/ajax/", handler.GetAjaxHanlder())
	http.HandleFunc("/", handler.GetNotFundHandler())
	http.ListenAndServe(":9090", nil)
}
