package main

//package a importer
import(
	"fmt"
	"errors"
	"math"
	//"time"
)

func main(){
	
	// display fonction
	fmt.Println("hello world")
	
	// variables
	var x int = 5
	var y int = 7
	var res int = x + y
	fmt.Println(res)

	x2 := 5
	y2 := 7
	res2 := x2 + y2
	fmt.Println(res2)

	// if condition
	test := 7
	if test > 6 {
		fmt.Println("More than 6")
	} else if x < 2 {
		fmt.Println("tata")
	} else{
		fmt.Println("toto")
	}

	//for loop
	for i := 0; i<5; i++ {
		fmt.Println(i)
	}
	//ou bien
	i:=0
	for i < 5 {
		fmt.Println(i)
		i++
	}
	
	//arrays
	var a [5]int
	a[2] = 6
	fmt.Println(a)

	b := [5]int {5,4,3,2,1}
	fmt.Println(b)

	c := []int {5,4,3,2,1}
	c = append(c, 13)
	fmt.Println(c)


	//boucle for pour un array
	arr := []string{"a", "b", "c"}

	for i, v := range arr {
		fmt.Println("index", i, "value", v)
	}

	//maps
	vertices := make(map[string]int)

	vertices["triangle"] = 2
	vertices["square"] = 4
	vertices["dodecagon"] = 12 
	fmt.Println(vertices)
	fmt.Println(vertices["triangle"])

	fmt.Println("square worth ", vertices["square"], "\nDeleting square of map...")
	delete(vertices, "square")
	fmt.Println("Square deleted\nSquare Now worth ", vertices["square"])

	//boucle for pour un map
	grec_alpha := make(map[string]string) 
	grec_alpha["a"] = "alpha"
	grec_alpha["b"] = "beta"
	grec_alpha["c"] = "gamma"

	for i, v := range grec_alpha {
		fmt.Println("index", i, "value", v)
	}

	fmt.Println("Test de la fonction sum", sum(2 , 3))

	fmt.Println("Test de la fonction SQRT")
	testSqrt, err := sqrt(16)
	if err != nil {
		fmt.Println(err)
	}else {
		fmt.Println(testSqrt)
	}


	// utilisation des struct
	p := person{name: "Jake", age: 22}
	fmt.Println(p)
	fmt.Println(p.age)

	//pointeurs
	point := 7
	fmt.Println("adresse de la variable point est", &point)
	inc(&point)
	fmt.Println("adresse de la variable point est", &point)

}

////////////// FONCTIONS //////////////////////

func sum (x int, y int) int {
	return x + y
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Undefinded for negative numbers")	
	}

	return math.Sqrt(x), nil
}

//fonction pour pointeur
func inc(x *int){
	//fmt.Println("adresse de la variable point est", x, *x)
	*x++
}

////////////// Struct //////////////////////
type person struct {
	name string
	age int
}
