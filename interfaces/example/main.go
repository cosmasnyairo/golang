package main

import "fmt"

type englishbot struct {
	greeting string
}

type swahilibot struct {
	greeting string
}

type bot interface {
	getGreeting() string
}

func main() {
	ebot := englishbot{
		"Good Morning!",
	}
	sbot := swahilibot{
		"Habari ya Asubuhi!",
	}

	printGreeting(ebot)
	printGreeting(sbot)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// We replace the 2 below functions with the above function
// func (eb englishbot) printGreeting() {
// 	fmt.Println(eb.getGreeting())
// }

// func (sb swahilibot) printGreeting() {
// 	fmt.Println(sb.getGreeting())
// }

func (eb englishbot) getGreeting() string {
	// Custom greeting
	return eb.greeting
}

func (sb swahilibot) getGreeting() string {
	// Custom greeting
	return sb.greeting
}
