package main

import (
	"example/hello/module/foo"
	f "fmt"
)

func appName() string {
	const AppName = "GoApp"
	var Version string = "1.0"
	return AppName + " " + Version
}

func Do(s string) (b string) {
	// var b string = s
	b = s
	{
		b := "bbbb"
		f.Println(b)
	}
	f.Println(b)
	return b
}
func main() {
	f.Println(foo.Max)
	f.Println(foo.Min)
	f.Println(foo.ReturnMin())
	f.Println(appName())
	f.Println(Do("aaa"))
}
