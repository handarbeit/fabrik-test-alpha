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
	fmt.Println(greeting.HelloE2E202608100429523676())
	fmt.Println(greeting.HelloE2E202608130450548185())
	fmt.Println(greeting.HelloE2E202608131322370922())
	fmt.Println(greeting.HelloE2E202608132002503163())
	fmt.Println(greeting.HelloE2E202608140336482864())
	fmt.Println(greeting.HelloE2E202608140529281443())
	fmt.Println(greeting.HelloE2E202608161836056156())
	fmt.Println(greeting.HelloE2E202608162025514426())
	fmt.Println(greeting.HelloE2E202608281527563489())
	fmt.Println(greeting.HelloE2E202608282032113677())
	fmt.Println(greeting.HelloE2E202608290258006960())
	fmt.Println(greeting.HelloE2E202609042333171389())
}
