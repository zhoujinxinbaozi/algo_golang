package _interface

import (
	"fmt"
	"testing"
)
/*
	定义接口，实现了该接口的所有方法的struct都是该接口的实例
 */

// MsgFilter 接口类型 过滤器
type MsgFilter interface {
	Filter(targets []string, target string) []string
	Name() string
}

type MsgFilter1 struct{}
type MsgFilter2 struct{}

// Filter 过滤
func (m1 *MsgFilter1) Filter(targets []string, target string) []string {
	var result []string
	for _, v := range targets {
		if v == target {
			continue
		}
		result = append(result, v)
	}
	return result
}

func (m1 *MsgFilter1) Name() string {
	return "MsgFilter1"
}

func (m2 *MsgFilter2) Filter(targets []string, target string) []string {
	var result []string
	for _, v := range targets {
		if v == target {
			continue
		}
		result = append(result, v)
	}
	return result
}

func (m2 *MsgFilter2) Name() string {
	return "MsgFilter2"
}

func TestInterface(t *testing.T) {
	var filters []MsgFilter
	filters = append(filters, &MsgFilter1{})
	filters = append(filters, &MsgFilter2{})
	for _, filter := range filters {
		fmt.Println(filter.Name())
	}
}
