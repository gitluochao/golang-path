package controller

import (
	"net/http"
	"html/template"
	 "util"
)

var logger = &util.LoggerHelp{}
type AdminController struct {

}
type User  struct {
	username string
}
func (controller *AdminController) IndexAction(w http.ResponseWriter,r http.Request,username string){
	t,err := template.ParseFiles("webapp/html/index.html")
	if err != nil {
		logger.Error("解析文件异常" + err.Error())
	}
	t.Execute(w,&User{username})
	return
}
