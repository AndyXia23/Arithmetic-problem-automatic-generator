package main

import (
	"awesomeProject/Generate"
	"awesomeProject/Write_in"
	"fmt"
)

func main() {
	var num int // the number of questions
	var slice1 = make([]int,num)
	var slice2 = make([]int,num)
	var slice3 = make([]string,num)

	fmt.Println("请输入题目数量")
	fmt.Scanf("%d",&num)
	slice1,slice2,slice3 = Generate.GenerateNumber(num)
	Write_in.CreateFile(num,slice1,slice2,slice3)
	Write_in.CreateFileWithResult(num,slice1,slice2,slice3)

}
