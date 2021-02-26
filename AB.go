// 自动生成10以内加减法
package main

import (
   "fmt"
 "math/rand"
)

func main() {
   var num int // the number of questions
 var a int
 var b int
 var c int
 fmt.Println("请输入题目数量")
   fmt.Scanf("%d",&num)
   var slice1 = make([]int,num)
   var slice2 = make([]int,num)
   var slice3 = make([]string,num)

   for i:=0;i<num;i++ {
      a = rand.Intn(10)
      b = rand.Intn(10)
      c = rand.Intn(2)
      slice1[i] = a
 slice2[i] = b
 if c == 0 {
         slice3[i] = "+"
 } else {
         slice3[i] = "-"
 }
   }
      fmt.Println("题目如下，")
   for i:=0;i<num;i++ {
      fmt.Println(slice1[i],slice3[i],slice2[i],"=")
   }
}
