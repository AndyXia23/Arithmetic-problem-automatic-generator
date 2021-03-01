package Generate

import (
	"math/rand"
)

func GenerateNumber(num int) (sliceA,sliceB []int, sliceC []string){
	var a int
	var b int
	var c int
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
		if (slice1[i]-slice2[i])<0 {
			i--
		}
	}
	return  slice1,slice2,slice3
}
