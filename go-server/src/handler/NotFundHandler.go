package handler
import (
	"net/http"
	"html/template"
	"util"
)
var logger = &util.LoggerHelp{}
func GetNotFundHandler() func(http.ResponseWriter,*http.Request){
	return func(w http.ResponseWriter,r *http.Request){
		if r.URL.Path == "/" {
			http.Redirect(w,r,"/login/index",http.StatusFound)
		}
		t,err := template.ParseFiles("webapp/html/404.html")
		if err != nil {
			logger.Info(err.Error())
		}
		t.Execute(w,nil)
	}
}


