package dao
import (
	"fmt"
)
var dataSource = &DataSource{}

type UserDao interface {
	QueryUserByUserName(username string) (*Student,error)
}

type UserDaoImpl struct {

}
type Student struct {
	Student_id  int64
	Username string
	Password string
}

func (userDaoImpl *UserDaoImpl) QueryUserByUserName(username string) (student *Student,err error){
	db,err := dataSource.Open();
	defer db.Close()
	if err != nil {
	     return
	}
	rows,err := db.Query("select student_id,username,password from online_student where username = ?",username)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer rows.Close()
    if rows.Next() {
		student = &Student{}
		rerr := rows.Scan(student.Student_id,student.Username,student.Password)
		if rerr!=nil {
		   fmt.Println()
		}
	}
	return
}

