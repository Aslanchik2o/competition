package main

import "fmt"


type Number interface{
	int64 | float64
}
type User struct{
	email string
	name string
}

func main(){
	a := []int64{1,313,4}
	b := []float64{2.313,1424.235,325}
	c := [] string{"1","2","32"}
	d := []User{
		{
			email: "steal@gmail.com",
			name: "Durka",
		},
		{
			email: "thin@gmail.com",
			name: "Cat",
		},
		{	email: "thin1@gmail.com",
		name: "Cat2",
	},
	}
	fmt.Println(sum(a))
	fmt.Println(sum(b))
	fmt.Println(c)
	fmt.Println(searchElement(d, User{
		email: "thin1@gmail.com",
		name: "Cat2",
	}))
	printAny(d)
}

func sum[V int64 | float64] (input []V)V {
	var result V
	for _, number := range input {
		result += number
	}
	return result
}
func searchElement[C comparable](elements []C, SearchEl C) bool {
	for _, el := range elements{
		if el == SearchEl {
return true
		}
	}
	return false
}
func printAny[A any] (input A){
	fmt.Println(input)

}