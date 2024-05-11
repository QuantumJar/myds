package main

import (
	"fmt"
	"io"
	"log"
	"myds/src/mrapps"
	"os"
	"sort"
)

func main() {

	file, err := os.Open("pg-being_ernest.txt")
	if err != nil {
		log.Fatalf("cannot open %v", "pg-being_ernest.txt")
	}
	content, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("cannot read %v", "pg-being_ernest.txt")
	}
	file.Close()
	//这里的kva无序
	kva := mrapps.Map("pg-being_ernest.txt", string(content))

	// for k, v := range kva {
	// 	fmt.Println(k, v)
	// }
	//对这个kva按照字幕表顺序排序
	sort.Sort(mrapps.ByKey(kva))

	// for k, v := range kva {
	// 	fmt.Println(k, v)
	// }

	oname := "mr-output-0"
	ofile, _ := os.Create(oname)

	//现在要把整个kva拆分为 多个切片，每个切片只包含一个相同key，
	for i, j := 0, 0; i < len(kva); {
		//reduce 需要一个单词和对应的list
		currentkey := kva[i].Key
		relatedList := []string{}
		//
		for kva[i].Key == kva[j].Key && j < len(kva) {
			relatedList = append(relatedList, kva[j].Value)

			if j != len(kva)-1 {
				j++

			} else {
				//若j指向了最后一个元素
				i = j
				break
			}

		}

		output := mrapps.Reduce(currentkey, relatedList)

		// this is the correct format for each line of Reduce output.
		fmt.Fprintf(ofile, "%v %v\n", currentkey, output)
		if i == len(kva)-1 {
			break
		}
		//更新起始位置
		i = j

	}

	ofile.Close()

}
