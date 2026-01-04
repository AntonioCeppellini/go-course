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
fmt.Println(x) // -> output will be 16

// IMPORTANT:
// go doesn't allow us to have unused variables
// we can ignore a variable doing like that:

func getPoint() (int, int){
  return 2, 3
}

x, _ = getPoint()

// doing this:
func getCoordinates() (x, y int){ // we could see this as a way of documenting the function but the implicit return is a bit difficult to read, better to be always explicit
  // x and y are initialized with zero values
  return // automatically returning x and y
}

// is equal to:
func getCoordinates() (int, int){
  var x int
  var y int
  return x, y
}

// GUARD CLAUSES
// early return or continue through a loop
func divide(dividend, divisor int) (int, error) {
  if divisor == 0 {
    return 0, errors.New("Can not divide by zero")
  }
  return dividend/divisor, nil
}
