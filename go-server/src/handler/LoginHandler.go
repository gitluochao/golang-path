package handler
import (
	"net/http"
	"strings"
	 "controller"
	"reflect"
)
func GetLoginHanlder() func(http.ResponseWriter,*http.Request){
	return func(w http.ResponseWriter,r *http.Request){
		pathInfo := strings.Trim(r.URL.Path,"/")

		paths := strings.Split(pathInfo,"/")
		var actionName = ""
		//取出actionName
		if len(paths) > 1 {
			actionName = strings.Title(paths[1]) + "Action"
		}

		login := &controller.LoginController{}
		responseValue := reflect.ValueOf(w)
		requestValue := reflect.ValueOf(r)
		reflect.ValueOf(login).MethodByName(actionName).Call([]reflect.Value{responseValue,requestValue})
	}
}


