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
age := 24 // this will be an int but we don't know if 32, 64 ecc, it will depend on your machine
height := 1.85

// both ways assign a *STATIC* type

// --------------------------------------

// we can declare multiple variables on the same line

is_human, name, age, height := true, "Spiderman", 24, 1.85

// --------------------------------------

// in go we have also costants
// constants can not use the `:=` operator
// constants need to be known at compile time

const color = "violet"
