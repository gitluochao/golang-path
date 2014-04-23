package controller

import (
	"net/http"
	"html/template"
	"fmt"
)
//use reflect / config by file
type LoginController struct{
}
//golang
func (control LoginController) IndexAction(w http.ResponseWriter,r *http.Request){
	//登陆页面渲染过去
	t,err := template.ParseFiles("webapp/html/login/login.html")
	if err!= nil {
		logger.Info("获取登陆页面异常")
		logger.Error(err.Error())
		fmt.Println(err.Error())
	}
	t.Execute(w,nil)
}
