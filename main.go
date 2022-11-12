package main

import (
	"fmt"
)

func main(){
	x := 5
	y := 5
	ans := min(x,y)
	fmt.Println(ans)
	MultiplicationTable()
	result := recursive(5)
	fmt.Println(result)
	array:=[]int{11,12,55,3,1,27}
	fmt.Println(BubbleSort(array))
}

func min(x int,y int) (string) {
	if x > y {
		return "x比較大"
	}else if y > x{
		return "y比較大"
	}else {
		return "一樣大"
	}
}
func MultiplicationTable(){
	for i := 1; i <=9 ; i++ {
		for j := 1; j <= 9; j++ {
			fmt.Println(i , "*" , j , "=" , i*j)
		}
	}
}
func recursive(n int) int {
	if n == 1{
		return 0
	}else{
		return 2 * recursive(n - 1) + 1
	}
}
func BubbleSort(a []int) []int{
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a)-i-1; j++ {
			if(a[j] > a[j+1]) {
				a[j],a[j+1]= a[j+1],a[j]
			}
		}
	}
	return a
}