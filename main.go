package main

import (
	"fmt"
	"os"
)

func getFormatter(name string) PasswordFormat {
	var formatter PasswordFormat
	switch name {
	case "bitWarden":
		formatter = &BitWarden{}
	case "onePassword":
		formatter = &OnePassword{}
	}

	return formatter
}

func main() {
	from := os.Args[1]
	to := os.Args[2]
	file := os.Args[3]

	fromFormatter := getFormatter(from)
	toFormatter := getFormatter(to)

	entry := fromFormatter.read(file)
	fmt.Println(toFormatter.toString(entry))
}
