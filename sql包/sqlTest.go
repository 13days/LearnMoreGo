package sql包

import (
	"database/sql"
	"fmt"
	"strings"
)

var (
	dbhostsip  = "127.0.0.1:3306"
	dbusername = "root"
	dbpassword = "123456"
	dbname     = "gdufo2o"
)

// 初始化数据库
func InitDB() (*sql.DB, error) {
	//构建连接信息
	dbinfo := strings.Join([]string{dbusername, ":", dbpassword, "@tcp(", dbhostsip, ")/", dbname, "?charset=utf8"}, "")
	fmt.Println(dbinfo)
	//打开数据库，前面是驱动名称，所以要导入:mysql驱动github.com/go-sql-driver/mysql
	dbins, err := sql.Open("mysql", dbinfo)
	if nil != err {
		fmt.Println("Open Database Error:", err)
		return nil, err
	}
	// 设置数据库的最大连接数
	dbins.SetConnMaxLifetime(100)
	// 设置数据库最大的闲置连接数
	dbins.SetMaxIdleConns(10)
	// 验证连接
	if err = dbins.Ping(); nil != err {
		fmt.Println("Open Database Fail,Error:", err)
		return nil, err
	}
	fmt.Println("Connect Success!!!")
	return dbins, nil
}
