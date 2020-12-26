package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	Name := "Sunny Bhandare "

	tpl := `
		<!DOCTYPE html>
		<html lamg = "en">
		<head>
			<meta charset = "UTF-8">
			<title>
			Hello World
			</title>
		</head>
		<body>
		
		<h1> ` + Name + `</h1>
		
		</body>
		</html>
	`
	fmt.Println(tpl)

	nf, err := os.Create("hi.html")
	if err != nil {
		fmt.Println(err)
	}

	defer nf.Close()

	io.Copy(nf, strings.NewReader(tpl))
}
