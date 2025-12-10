// how to declare variables:

// hard way:

var is_human bool // we do not assign a value here, just declaring the type
var name string
var age int
var height float32

// and then we can use them ...

// --------------------------------------

// easy way:
// using short assignment := this operator infers the type of
// the new variable based on the value.

is_human = true
name := "Spider Man"
age := 24
height := 1.85

// both of the way assign a *STATIC* type

// --------------------------------------

// we can declare multiple variables on the same line

is_human, name, age, height := true, "Spiderman", 24, 1.85
