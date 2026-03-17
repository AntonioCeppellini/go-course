// STRUCTS
// Like a python dict or an obj in JS
// Use them to represent structured data.
type car struct {
  Make string
  Model string
  Height int
  Width int
}

// This creates a new struct type called `car`
// all cars have `Make`, `Model`, `Height` and `Width`.

// Structs can be nested to represent more complex entities
type car struct {
  Make string
  Model string
  Height int
  Width int
  FrontWheel Wheel
  BackWheel Wheel
}

type Wheel struct {
  Radius int
  Material string
}

// The field of a struct can be accessed using the dot `.` operator.
myCar := car{}
myCar.FrontWheel.Radius = 5

// ANONYMOUS STRUCTS
// Like a normal strcut but is defined without a
// name and cannot be referenced elsewhere in the code
myCar := struct {
  Make string
  Model string
}

// Can be nested as fields within other structs:
type car struct {
  Make string
  Model string
  Height int
  Width int
  // Wheel is an anonymmous struct
  Wheel struct {
    Radius int
    Material string
  }
}

// Usually is better to use named struct
// BUT
// Anonymous struct can be compiled more quickly

// EMBEDDED STRUCTS

myCar := struct {
  Make string
  Model string
}

type truck struct {
  // "car" is embedded, so the definition of a
  // "truck" now also additionally contains all
  // of the fields of the car struct 
  car
  bedsize int
}

// in an embedded struct we can access to a field at the top-level, unlike the, what we saw before, nested structs
myTruck := truck{}
fmt.Println(myTruck.Make) // instead of `myTruck.car.make`

// GO IS NOT AN OBJECT ORIENTED PROGRAMMING LANGUAGE BUT
// we can call what's happening in the embedded structs "inheritance"

// REPEAT WITH ME, GO IS NOT OBJECT ORIENTED
// BUT
// it supports methods that can be defined on structs.
// methods are functions that have a receiver, a special parameter that syntactically goes before the name of
// the function.

type rectanglem struct {
  width int
  height int
}

// area has a receive of (r rectangle)
func (r rectangle) area () int {
  return r.width * r.height
}

r := rectangle{
  width: 5,
  height: 10,
}

fmt.Println(r.area())
