package pointers

import (
	"log"
)

func PointersInvocation() {
	log.SetPrefix("pointers: ")
	log.SetFlags(0)

	PointerIntroExample()
	StructExample()
}

func PointerIntroExample() {
	log.Println("Pointer Introduction Example")

	i, j := 41, 2701
	p := &i
	log.Println(*p) // read i through the pointer p
	*p = 21
	log.Println(i) // see the new value of i

	p = &j
	log.Println(*p) // read j through the pointer p
	*p = *p / 37
	log.Println(j) // see the new value of j
}

type Vertex struct {
	X int
	Y int
}

func StructExample() {
	log.Println("Struct Example")
	v := Vertex{1, 2}
	log.Println(v)

	log.Println("Struct Accessed Values")
	log.Println("v.X:", v.X)
	log.Println("v.Y:", v.Y)

	log.Println("Struct Accessed Values through Pointer")
	p := &v
	log.Println("p.X:", p.X)
	log.Println("p.Y:", p.Y)

	p.X = 1e9
	log.Println("v:", v)
}
