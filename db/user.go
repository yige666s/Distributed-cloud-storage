package db

import (
	"Distributed-cloud-storage/db/mysql"
	"fmt"
)

// 用户注册
func UserSignup(username string, passwd string) bool {
	stmt, err := mysql.BDConn().Prepare("insert ignore into tbl_user(`user_name`,`user_pwd`) value (?,?)")
	if err != nil {
		fmt.Println("Failed to insert,err:" + err.Error())
		return false
	}
	defer stmt.Close()

	ret, err := stmt.Exec(username, passwd)
	if err != nil {
		fmt.Println("filed to insert,err " + err.Error())
		return false
	}

	if rowsAffected, err := ret.RowsAffected(); nil == err && rowsAffected > 0 {
		return true
	}
	return false
}

func UserSignIn(username string, encpwd string) bool {
	stmt, err := mysql.BDConn().Prepare("select * from tbl_user where user_name = ? limit 1")
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	rows, err := stmt.Query(username)
	if err != nil {
		fmt.Println(err.Error())
		return false
	} else if rows == nil {
		fmt.Println("username not found : " + username)
		return false
	}
	pRows := mysql.ParseRows(rows)
	if len(pRows) > 0 && string(pRows[0]["user_pwd"].([]byte)) == encpwd { // 比较密码是否一致
		return true
	}
	return false
}

// 刷新用户登录的token
func UpdateToken(username string, token string) bool {
	stmt, err := mysql.BDConn().Prepare("replace into tbl_user_token(`user_name`,`user_token`) values (?,?)")
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	defer stmt.Close()

	_, err = stmt.Exec(username, token)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}

type User struct { //返回的用户信息，可以与tbl_user中的字段一致
	Username     string
	Email        string
	Phone        string
	SignupAt     string
	LastActiveAt string
	Status       int
}

// 查询用户信息
func GetUserInfo(username string) (User, error) {
	user := User{}
	stmt, err := mysql.BDConn().Prepare("select usre_name,signup_at fomr tbl_user where user_name = ? limit 1")
	if err != nil {
		fmt.Println(err.Error())
		return user, err
	}
	defer stmt.Close()

	//执行查询操作
	err = stmt.QueryRow(username).Scan(&user.Username, &user.SignupAt)
	if err != nil {
		return user, err
	}
	return user, nil
}
