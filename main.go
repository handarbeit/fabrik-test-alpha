package main

import (
	"fmt"
	"os"

	"github.com/handarbeit/fabrik-test-beta/pkg/greeting"
)

func main() {
	name := "world"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	fmt.Println(greeting.GreetingFor(name))
	fmt.Println(greeting.HelloE2E())
	fmt.Println(greeting.HelloIssue31())
	fmt.Println(greeting.HelloIssue52())
	fmt.Println(greeting.HelloIssue3472())
	fmt.Println(greeting.HelloIssue3496())
	fmt.Println(greeting.HelloIssue3620())
}
