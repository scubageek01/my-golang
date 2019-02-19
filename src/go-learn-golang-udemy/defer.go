package main

import "fmt"

func main() {
	defer deferClose()  // Close file after opening file to do something.  Free up resources
	openFile()
}

func deferClose() {
	fmt.Println("closed")
}

func openFile() {
	fmt.Println("open")
}
