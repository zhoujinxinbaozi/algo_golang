package main

import (
	"encoding/json"
	"fmt"
)
/*
	结构体与json转换
 */
func main() {
	str := `{"name":"周金鑫", "age":24}`
	var p1 person
	json.Unmarshal([]byte(str), &p1)
	fmt.Println(p1)
	bytes, _ := json.Marshal(p1)
	fmt.Println(string(bytes))

}

type person struct {
	Name string `json:name` // 字段暴露需要大写，在json中的字段为小写name
	Age  int8   `json:age`
}
