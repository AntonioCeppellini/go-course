// Some of the comparison operators in go:
// == equal to
// != not equal to
// < less than
// > greater than
// <= less than or equal to
// >= greater than or equal to
import fmt

if mask_is_on == true {
  fmt.Println("I'm SpiderMan")
} else {
  fmt.Println("I'm Peter")
}

// i could do this
// it allows us to write a more concise code but
// also to limit the scope that the variable exists within

if length := getLength(name); length > 20 {
  fmt.Println("Your name is quite long")
}
