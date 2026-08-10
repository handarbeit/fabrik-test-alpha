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
	fmt.Println(greeting.HelloIssue3737())
	fmt.Println(greeting.HelloIssue3834())
	fmt.Println(greeting.HelloIssue3855())
	fmt.Println(greeting.HelloIssue3922())
	fmt.Println(greeting.HelloE2E202608011257146717())
	fmt.Println(greeting.HelloE2E202608011442176757())
	fmt.Println(greeting.HelloE2E202608030436207283())
	fmt.Println(greeting.HelloE2E202608041142376848())
	fmt.Println(greeting.HelloE2E202608100155218654())
}
