package main

import (
	"container/list"
	"fmt"
)


type Stack struct {
	list *list.List
}

func NewStack() *Stack {
	list := list.New()
	return &Stack{list}
}

func (stack *Stack) Push(value interface{}) {
	stack.list.PushBack(value)
}

func (stack *Stack) Pop() interface{} {
	e := stack.list.Back()
	if e != nil {
		stack.list.Remove(e)
		return e.Value
	}
	return nil
}

func (stack *Stack) Peak() interface{} {
	e := stack.list.Back()
	if e != nil {
		return e.Value
	}

	return nil
}

func (stack *Stack) Len() int {
	return stack.list.Len()
}

func (stack *Stack) Empty() bool {
	return stack.list.Len() == 0
}

func isValid(s string) bool {
	result := true
	stack := NewStack()
	for _, value := range s {
		switch value {
		case '(', '{', '[':
			stack.Push(value)
		case ')':
			if !stack.Empty() && stack.Pop() == '(' {

			} else {
				result = false
				break
			}
		case ']':
			if !stack.Empty() && stack.Pop() == '[' {

			} else {
				result = false
				break
			}
		case '}':
			if !stack.Empty() && stack.Pop() == '{' {

			} else {
				result = false
				break
			}
		}
	}
	return result && stack.Empty()
}

/**
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。

示例 1:

输入: "()"
输出: true
示例 2:

输入: "()[]{}"
输出: true
示例 3:

输入: "(]"
输出: false
示例 4:

输入: "([)]"
输出: false
示例 5:

输入: "{[]}"
输出: true

solution:
	难点在于自己实现stack。 左括号全都入栈，遇到右括号就弹出一个看是否匹配，不匹配直接false，所有的都遍历完，全都匹配并且stack为空，返回true
 */
func main() {
	fmt.Println(isValid("()[]{}"))
}