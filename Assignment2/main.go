package main

import (
	"fmt"
	"io"
	"os"
)

func main(){
fi, error := os.Open(os.Args[1])

if error != nil {
	fmt.Println("Error:", error)
	os.Exit(1)
}

io.Copy(os.Stdout, fi)


}