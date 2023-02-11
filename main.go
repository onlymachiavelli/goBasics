package main 
import (
	"fmt"
	"math/rand"
) 



func fill(n int, arr[5] int) {
	for i:=0;i<n;i++ {
		arr[i] = rand.Intn(100) // min  max 100 or 99, idk 
	}
}

func addShit(n int, arr[5] int) {
	fmt.Println("Enter the number you want to add ! ")
	var num int
	fmt.Scan(&num) 
	n++

	//push to the array 
	//arr[n] = num
	arr = append(array, num) 


}

func main () {
	var n int=5 ; 
	//fmt.Println("Enter the length of the array ! ") 
	//fmt.Scan(&n) 

	arr := make(int, 5) 

	fill(n, &arr)

	fmt.Println(arr)
	addShit(n, &arr)
	fmt.Println(arr)

	
	

}