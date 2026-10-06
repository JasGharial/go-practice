package basics

import (
	"errors"
	"fmt"
	"log"
	"math"
	"runtime"
	"golang.org/x/tour/pic"
)

func Basic() {
	log.SetPrefix("basics.go: ")
	log.Println("Basics Package Invoked")

	// Declarations
	forLoopExample := []int{1, 2, 6, 7, 8}

	// === Function Calls === //
	LoopSum, err := ForLoop(forLoopExample)
	if err != nil {
		log.Fatal("Error:", err)
	}
	log.Println("Sum:", LoopSum)

	IfElsePowResult := IfElsePow(2, 3, 8)
	log.Println("IfElsePow Result:", IfElsePowResult)

	SwitchExample()

	ArrayExample()

	SlicesLiteralExample()

	pic.Show(Pic)
}

// ForLoop demonstrates a basic for loop in Go and returns the sum of the provided numbers.
func ForLoop(numbers []int) (int, error) {
	log.Println("ForLoop Example")
	if len(numbers) == 0 {
		return 0, errors.New("Please provide a valid list of numbers")
	}

	sum := 0
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}
	return sum, nil
}

// IfElsePow demonstrates a basic if-else statement in Go and returns the result of raising x to
// the power of n if it's less than lim, otherwise it returns lim.
func IfElsePow(x, n, lim float64) float64 {
	log.Println("IfElsePow Example")
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}

	return lim
}

func SwitchExample() {
	os := runtime.GOOS
	log.Println("SwitchExample OS:", os)

	switch os {
		case "darwin":
			log.Println("OS X.")
		case "linux":
			log.Println("Linux.")
		default:
			log.Printf("%s.", os)
	}
}

func ArrayExample() {
	log.Println("Array Examples")
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	log.Println(names)

	a := names[0:2]
	b := names[1:3]
	log.Println(a, b)

	b[0] = "XXX"
	log.Println(a, b)
	log.Println(names)
}

func SlicesLiteralExample() {
	log.Println("Slice Literal Example")
	q := []int{2, 3, 5, 7, 11, 13}
	log.Println("q:", q)

	r := []bool{true, false, true, true, false, true}
	log.Println("r:", r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}

	log.Println("s:", s)
}

func Pic(dx, dy int) [][]uint8 {
	log.Println("Pic Example")
	pic := make([][]uint8, dy)
	for i := range pic {
		pic[i] = make([]uint8, dx)
		for j := range pic[i] {
			pic[i][j] = uint8(i ^ j)
		}
	}
	return pic
}