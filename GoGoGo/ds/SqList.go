// 顺序表
package ds

import (
	"errors"
	"fmt"
)

type SqList struct {
	items  []string //元素数组
	len    int    //顺序表当前长度
	maxlen int    //顺序表最大长度
}

// 初始化顺序表
func InitSqList(size int) *SqList {
	return &SqList{make([]string, size), 0, size}
}

//获取当前长度
func (self *SqList) Length() int {
	return self.len
}

//获取最大长度
func (self *SqList) Size() int {
	return self.maxlen
}

//追加一个元素
func (self *SqList) Add(item string) error {
	if self.len == self.maxlen {
		return errors.New("SqList is full")
	}
	self.len++
	self.items[self.len-1] = item
	return nil
}

//删除一个元素
func (self *SqList) Delete(index int) error {
	err := self.verify(index)
	if err != nil {
		return err
	}
	if self.len-1 == index {
		self.len--
	} else {
		self.len--
		for i := index; i < self.len; i++ {
			self.items[i] = self.items[i+1]
		}
	}
	return nil
}

//修改一个元素
func (self *SqList) Update(index int, item string) error {
	err := self.verify(index)
	if err != nil {
		return err
	}
	self.items[index] = item
	return nil
}

//根据索引查询元素
func (self *SqList) Select(index int) (string, error) {
	err := self.verify(index)
	if err != nil {
		return "", err
	}
	return self.items[index], nil
}

//打印顺序表
func (self *SqList) Show() {
	fmt.Print("[")
	for i := 0; i < self.len; i++ {
		fmt.Print(self.items[i])
		if i < self.len-1 {
			fmt.Print(",")
		}
	}
	fmt.Print("]")
}

func (self *SqList) verify(index int) error {
	if self.len-1 < index || index < 0 {
		return errors.New("Index Out of Bounds")
	}
	return nil
}
