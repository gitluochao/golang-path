package controller

import (
	"net/http"
	"encoding/json"
	"dao"
)
type Result struct {
	Ret int
	RetMessage string
	Data interface {}
}
var userDao dao.UserDao

type AjaxController struct {

}
func init(){
//	userDao = &UserDaoImpl{}
}
func (constroller *AjaxController) LoginAction(w http.ResponseWriter,r *http.Request){
	w.Header().Set("content-type", "application/json") //response json format
	err := r.ParseForm()
	if err != nil {
		 logger.Error(err.Error())
		 toJson(w,101,"参数错误",nil)
	}
	username := r.FormValue("username")
	password := r.FormValue("password")
	student,err := userDao.QueryUserByUserName(username)
	if err!= nil {
	    logger.Error("查询异常")
		toJson(w,102,"系统错误",nil)
	}
	if student.Username== ""  {
	    toJson(w,103,"用户不存在",nil)
	}
	if student.Password != password {
	    toJson(w,104,"密码错误",nil)
	}
	//登录成功写cookie

	cookie := http.Cookie{Name:"username",Value:username,Path:"/"}
	http.SetCookie(w,&cookie)
	toJson(w,100,"登录成功",nil)
	return
}
func toJson(w http.ResponseWriter,ret int,retMessage string,data interface {}){
	out := &Result{ret,retMessage,data}
	jsonStr,err := json.Marshal(out)
	if err != nil {
	   return
	}
	w.Write(jsonStr)
}
