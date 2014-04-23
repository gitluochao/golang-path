package handler
import (
	"net/http"
	"strings"
	"controller"
	"reflect"
)
func GetAjaxHanlder() func(http.ResponseWriter,*http.Request){
	return func(w http.ResponseWriter,r *http.Request){
		pathInfo := strings.Trim(r.URL.Path,"/")
		paths := strings.Split(pathInfo,"/")
		var actionName = ""

		if len(paths) > 1 {
			actionName = strings.Title(paths[len(paths) - 1]) + "Action"
		}
		controller := &controller.AjaxController{}
		responseVal := reflect.ValueOf(w)
		requestVal := reflect.ValueOf(r)
		reflect.ValueOf(controller).MethodByName(actionName).Call([]reflect.Value{responseVal,requestVal})
	}
}


