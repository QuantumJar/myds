package mrapps

import "strings"

type KeyValue struct {
	Key   string
	Value string
}

/*
	Map函数：takes an input pair and produces a set of intermediate key/value pairs.

这里的参数一个是文件名，这个没有特别的作用，一个是文件的内容,返回值就是需要生成的中间键值对
*/
func Map(filename, content string) []KeyValue {

	//1.将内容分词
	words := strings.Fields(content)

	//对words切片进行map
	intermediateKey := make([]KeyValue)
	for _, word := range words {

	}

}
