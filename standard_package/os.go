package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// os.Exit(1)
	// fmt.Println("start")

	// defer func() { fmt.Println("defer") }()
	// os.Exit(1)

	// f, err := os.Open("os.txt")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// defer f.Close()

	// go build -o getargs os.go
	// ./getargs 123 456 789

	// fmt.Println(os.Args[0])
	// fmt.Println(os.Args[1])
	// fmt.Println(os.Args[2])
	// fmt.Println(os.Args[3])
	// fmt.Printf("length=%d\n", len(os.Args))
	// for i, v := range os.Args {
	// 	fmt.Println(i, v)
	// }

	// f, _ := os.Create("foo.txt")
	// f.Write([]byte("Hello\n"))
	// f.WriteAt([]byte("Golang"), 6)
	// f.Seek(0, os.SEEK_END)
	// f.WriteString("yeah")

	// f, err := os.Open("foo.txt")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// defer f.Close()
	// bs := make([]byte, 128)
	// n, err := f.Read(bs)
	// fmt.Println(n)
	// fmt.Println(string(bs))
	// bs2 := make([]byte, 128)
	// nn, err := f.ReadAt(bs2, 10)
	// fmt.Println(nn)
	// fmt.Println(string(bs2))

	f, err := os.OpenFile("foo.txt", os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	bs := make([]byte, 128)
	n, err := f.Read(bs)
	fmt.Println(n)
	fmt.Println(string(bs))
}
