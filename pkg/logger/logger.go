package logger

import "fmt"

var prefix string = "Qi"

func Log(content string) {
	fmt.Println(prefix, "-", content)
}

func init() {
	prefix = "Ellen"
}
