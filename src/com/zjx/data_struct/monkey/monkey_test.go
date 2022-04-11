package monkey

import (
	"fmt"
	"reflect"
	"testing"

	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"
)

// go test com/zjx/algo_golang/src/com/zjx/data_struct/monkey -run=TestMonkeyMethod -v -gcflags=-l
func TestMonkeyMethod(t *testing.T) {
	e := &Entity{
		Age: 18,
	}
	monkey.PatchInstanceMethod(reflect.TypeOf(e), "GetAge", func(entity *Entity) int8 {
		return entity.Age
	})
	assert.EqualValues(t, int8(18), e.GetAge())
}

type Entity struct {
	Age int8
}

//go:noinline  添加上这个注释 就可以不用-gcflag参数
func (e *Entity) GetAge() int8 {
	return 0
}

// go test com/zjx/algo_golang/src/com/zjx/data_struct/monkey -run=TestMonkeyFunc -v -gcflags=-l 禁用编译器优化 才可以使用monkey
func TestMonkeyFunc(t *testing.T) {
	monkey.Patch(GetInfoByUID, func(int64) (*UserInfo, error) {
		return &UserInfo{Name: "zhoujinxin"}, nil
	})

	name := MyFunc(123)
	fmt.Println("name : ", name)
	assert.True(t, name == "hello zhoujinxin")
}

func MyFunc(uid int64) string {
	u, err := GetInfoByUID(uid)
	if err != nil {
		return "welcome"
	}

	// 这里是一些逻辑代码...

	return fmt.Sprintf("hello %s", u.Name)
}

type UserInfo struct {
	Name string
}

func GetInfoByUID(int64) (*UserInfo, error) {
	return &UserInfo{}, nil
}

// go test com/zjx/algo_golang/src/com/zjx/data_struct/monkey -run=TestMonkeyInterface -v -gcflags=-l  必须是导出函数
func TestMonkeyInterface(t *testing.T) {
	var interImpl Inter
	interImpl = &InterImpl{}
	monkey.PatchInstanceMethod(reflect.TypeOf(interImpl), "GetInterType", func(p *InterImpl) int8 {
		return int8(10)
	})
	fmt.Println(interImpl.GetInterType())
}

type Inter interface {
	GetInterType() int8
}

type InterImpl struct {
}

func (p *InterImpl) GetInterType() int8 {
	return int8(0)
}
