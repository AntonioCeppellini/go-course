// how to declare a function

func sub(x int, y int) int {  // this is the function signature: 
  return x-y                  // it tell us the name of it, what it takes
}                             // and what it returns

// if the arguments are of the same type i can do:

func sub(x, y int) int {
  return x-y
}

// if a function has multiple return values
// we need to wrap them like that:

func fullName() (string, string){
  return "Spider", "Man"
}

// variables aren't changed in a function
// a function works with a copy of a variable:

x := 15

func increment(x int) int {
  return x+1
}

increment(x)

fmt.Println(x) // -> output will be 15
// but if we do:
x = increment(x)

fmt.Println(x) // -> output will be 1

// IMPORTANT:
// go doesn't allow us to have unused variables
// we can ignore a variable doing like that:

func getPoint() (int, int){
  return 2, 3
}

x, _ = getPoint()
