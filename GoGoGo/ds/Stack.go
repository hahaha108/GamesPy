// 栈
package ds

import (
	"errors"
)

type Stack struct {
	items  []string
	len    int
	maxlen int
}

// 初始化栈
func InitStack(size int) *Stack {
	return &Stack{make([]string, size), 0, size}
}

//获取当前长度
func (self *Stack) Length() int {
	return self.len
}

//获取最大长度
func (self *Stack) Size() int {
	return self.maxlen
}

// 入栈
func (self *Stack) Append(item string) error {
	err := self.verify("append")
	if err != nil {
		return err
	}
	self.items[self.len] = item
	self.len++
	return nil
}

// 出栈
func (self *Stack) Pop() (string, error) {
	err := self.verify("pop")
	if err != nil {
		return "", err
	}
	self.len--
	item := self.items[self.len]
	return item, nil
}



// 插入删除检测
func (self *Stack) verify(action string) error {
	switch action {
	case "append":
		{
			if self.len >= self.maxlen {
				return errors.New("append error")
			}
		}
	case "pop":
		{
			if self.len == 0 {
				return errors.New("pop error")
			}
		}
	}
	return nil
}
