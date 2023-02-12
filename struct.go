package main
import ("fmt")

type Person struct  {
	fullname string 
	age int 
	height float32 
	weight int  
}
func main() {
	var  Alaa Person ; 
	Alaa.fullname = "Alaa Barka"
	Alaa.age = 20
	Alaa.height = 1.76
	Alaa.weight = 0 

	if Alaa.weight == 0 {
		fmt.Println("Alaa didn't save his weight ! ") 
	} 
	fmt.Println(Alaa)
}
