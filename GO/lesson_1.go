
// package main
// import "fmt"
// func main()  {

// 	var name string="hello"
// 	fmt.Println(name)
// }



// package main
// import "fmt"
// func main()  {

// 	var name string="Bexruz"
// 	var number int= 11
// 	fmt.Printf("hello my name is %v, and Im %d grades student", name, number)
// }


// package main

// import "fmt"

// const PI float64 = 3.14 // global constant

// func main() {
//     var radius float64 = 5.154
//     var area float64

//     area = PI * radius * radius
//     fmt.Printf("Radius: %.2f \nPI:%.1f \n", radius, PI)
//     fmt.Printf("Area of Circle is : %f", area)
//     fmt.Println("Thank You")
// }





// package main
// import "fmt"
// func main()  {
// 	var name string
//     var age int
// 	fmt.Print("Enter your name: ")
// 	fmt.Scanf("%s", &name)

// 	fmt.Print("Enter your age: ")
	
// 	if name=="Bexruz" {
// 		fmt.Print("wellcome Boss u r 18")
// 	}else{
// 		fmt.Scanf("%d", &age)
// 		fmt.Print("wellcome ", name, " You are ",age)

		
// 	}

// }



package main

import "fmt"

func main() {
	name := [1]string{"Bexruz"}
	age := [1]int{18}
	fmt.Printf("Hello, my name is %v, and I'm %d years old.", name, age)
}