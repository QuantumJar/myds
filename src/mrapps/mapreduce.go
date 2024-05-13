package mrapps

import (
	"strconv"
	"strings"
	"unicode"
)

type KeyValue struct {
	Key   string
	Value string
}

type ByKey []KeyValue

// Len implements sort.Interface.
func (a ByKey) Len() int {
	return len(a)
}

// ByKey 实现了这个len()方法，他就是一个interface的接口 	Len() int

// Swap swaps the elements with indexes i and j.

func (a ByKey) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a ByKey) Less(i, j int) bool {
	if a[i].Key < a[j].Key {
		return true
	} else {
		return false
	}
}

/*
	Map函数：takes an input pair and produces a set of intermediate key/value pairs.

这里的参数一个是文件名，这个没有特别的作用，一个是文件的内容,返回值就是需要生成的中间键值对
*/
func Map(filename, content string) []KeyValue {

	ff := func(r rune) bool {
		return !unicode.IsLetter(r)
	}
	//将内容分割为单词的切片
	words := strings.FieldsFunc(content, ff)

	//声明中间键值对的容器, 一个[]KeyValue
	intermediatekeyList := []KeyValue{}
	//map函数的关键一步，对这个单词切片的每一个单词map操作，生成一个中间键值对k/v list
	for _, v := range words {
		item := KeyValue{Key: v, Value: "1"}
		intermediatekeyList = append(intermediatekeyList, item)
	}
	return intermediatekeyList
}

// reduce 用于把和这个key相关的值reduce为一个更小的集合。这里的统计单词出现次数是把每一个小key的{"w","1"}
func Reduce(key string, values []string) string {

	return strconv.Itoa(len(values))

}
