package dao

import (
	"database/sql"
    _ "mysql"
)
type DataSource struct {

}
func (dataSource *DataSource) Open() (*sql.DB ,error){
	return sql.Open("mysql","sonar:123456@tcp(192.168.192.135:3358)/sonar?charset=utf8")
}





