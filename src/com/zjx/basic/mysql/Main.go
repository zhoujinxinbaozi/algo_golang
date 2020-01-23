package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func main() {
	// username:password@tcp(ip:port)/database
	dsn := "root:lidaosi.135@tcp(127.0.0.1:3306)/test"
	// check dsn format is valid or not
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("check dsn failed, dsn:%s, err:%v", dsn, err)
		return
	}
	err = db.Ping()
	if err != nil {
		fmt.Printf("connect database failed, dsn:%v, err:%v", dsn, err)
		return
	}
	// 设置连接池中最大的连接数
	db.SetMaxOpenConns(10)
	// 设置空闲的最大连接数
	db.SetMaxIdleConns(5)
	fmt.Printf("connet database success!\n")
	//queryRow()
	//query()
	insert()
}

type student struct {
	name string
	age  string
}

// 获取唯一的一行数据
func queryRow() {
	var s1 student
	query := "select * from student  where name=?;"
	// go 语言中db是一个连接池，在进行查询之后必须调用Scan(),会释放该链接
	err = db.QueryRow(query, "zhoujinxin").Scan(&s1.name, &s1.age)
	if err != nil {
		fmt.Printf("get row failed, query:%v", query)
		return
	}
	fmt.Printf("name:%s, age:%s", s1.name, s1.age)
}

// 查询多行
func query(){
	sql := "select * from student"
	rows, err2 := db.Query(sql)
	if err2 != nil{
		fmt.Printf("execute sql failed, sql:%v, err:%v", sql, err2)
	}
	// 多行需要自己 释放连接
	defer rows.Close()
	for rows.Next(){
		var s1 student
		rows.Scan(&s1.name, &s1.age)
		fmt.Printf("name:%s, age:%s\n", s1.name, s1.age)
	}
}
// 插入，删除，修改
func insert(){
	sql := "insert into student values(\"姚依林3\",25);"
	n, err2 := db.Exec(sql)
	if err2 != nil{
		fmt.Println(err2)
	}
	i, _:= n.RowsAffected()
	id, _ := n.LastInsertId()
	fmt.Println("id:", id)
	fmt.Printf("affect %d 行", i)

}
