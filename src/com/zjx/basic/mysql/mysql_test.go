package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"testing"
	"time"
)

var conn *gorm.DB

const (
	tableName = "student"
)

type Student struct {
	ID          int64  `gorm:"column:id"`
	StudentID   int64  `gorm:"column:student_id"`
	StudentName string `gorm:"column:student_name"`
	Timestamp   int64  `gorm:"column:time_stamp"`
}

func TestMysql(t *testing.T) {
	conn, err = gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/sticker")
	if err != nil {
		fmt.Printf("conn failed, err : %v", err)
		panic("open failed")
	}
	fmt.Println(conn)
	conn.LogMode(true)

	student := Student{
		StudentID:   100,
		StudentName: "1",
		Timestamp:   time.Now().Unix(),
	}
	createStudent(student)
	//updateStudent(map[string]interface{}{"student_id": 100})
	//deleteStudent(student)
	//findOne(student.StudentID)
	findList(student.StudentName)
}

func createStudent(student Student) {
	if err := conn.Table(tableName).Create(&student).Error; err != nil {
		fmt.Printf("create failed, err : %v\n", err)
		return
	}
	fmt.Println("create success")
}

func deleteStudent(student Student) {
	if err := conn.Table(tableName).Where("student_name = ? and student_id = ?", student.StudentID, student.StudentName).Delete(&Student{}).Error; err != nil {
		fmt.Printf("delete failed, err : %v", err)
		return
	}
	fmt.Println("delete success")
}

func updateStudent(m map[string]interface{}) {
	if err := conn.Table(tableName).Where("student_id = ?", 1).UpdateColumns(m).Error; err != nil {
		fmt.Printf("updateStudent failed, err : %v", err)
		return
	}
	fmt.Println("updateStudent success")
}

func findOne(studentID int64) {
	student := &Student{}
	if err := conn.Table(tableName).Where("student_id = ?", studentID).First(student).Error; err != nil {
		fmt.Printf("findOne failed, err : %v", err)
		return
	}
	fmt.Printf("findOne success, student : %v\n", student)
}

func findList(studentName string) {
	var studentList []Student
	if err := conn.Table(tableName).Where("student_name = ?", studentName).Find(&studentList).Error; err != nil {
		fmt.Printf("findOne failed, err : %v", err)
		return
	}
	fmt.Printf("findOne success, student : %v\n", studentList)
}
