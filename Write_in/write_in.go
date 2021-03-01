package Write_in

import (
	"fmt"
	"os"
	"strconv"
)

func CreateFile(num int,sliceA,sliceB []int, sliceC []string) {
	//创建文件
	f, err := os.Create("a.txt")
	//判断是否出错
	if err != nil {
		fmt.Println(err)
	}

	var str string
	for i:=0;i<num;i++ {
		// WriteString:写入一个字符串
		str = strconv.Itoa(sliceA[i])+sliceC[i]+strconv.Itoa(sliceB[i])+"="
		_ , errW := f.WriteString(str)
		//判断是否出错
		if errW != nil {
			fmt.Println(errW)
		}
		f.WriteString("\n")
	}
	f.Close()
}

func CreateFileWithResult(num int,sliceA,sliceB []int, sliceC []string) {
	//创建文件
	f, err := os.Create("b.txt")
	//判断是否出错
	if err != nil {
		fmt.Println(err)
	}

	var str string
	for i:=0;i<num;i++ {
		// WriteString:写入一个字符串
		if sliceC[i] == "+" {
			str = strconv.Itoa(sliceA[i])+sliceC[i]+strconv.Itoa(sliceB[i])+"="+strconv.Itoa(sliceA[i]+sliceB[i])
		} else {
			str = strconv.Itoa(sliceA[i])+sliceC[i]+strconv.Itoa(sliceB[i])+"="+strconv.Itoa(sliceA[i]-sliceB[i])
		}
		_ , errW := f.WriteString(str)
		//判断是否出错
		if errW != nil {
			fmt.Println(errW)
		}
		f.WriteString("\n")
	}
	f.Close()
}

