package handler

import (
	"net/http"
	"strings"
	"controller"
	"reflect"
)
func GetAdminHanlder() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter,r *http.Request){
	    cookie,err := r.Cookie("username")

		if err != nil || cookie == nil  {
			http.Redirect(w,r,"/login/index",http.StatusFound)
		}
		pathInfo := strings.Trim(r.URL.Path,"/")
		paths := strings.Split(pathInfo,"/")
		var actionName = ""
		if len(paths) > 1 {
			actionName = strings.Title(paths[len(paths) - 1]) + "Action"
		}
		username := cookie.Value
		controller := &controller.AdminController{}
		responseVal := reflect.ValueOf(w)
		requestVal := reflect.ValueOf(r)
		usernameVal := reflect.ValueOf(username)
		reflect.ValueOf(controller).MethodByName(actionName).Call([]reflect.Value{responseVal,requestVal,usernameVal})
	}
}

