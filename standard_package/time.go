package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t)
	t2 := time.Date(2020, 6, 10, 10, 10, 10, 0, time.Local)
	fmt.Println(t2)
	fmt.Println(t2.Year())
	fmt.Println(t2.YearDay())
	fmt.Println(t2.Month())
	fmt.Println(t2.Weekday())
	fmt.Println(t2.Day())
	fmt.Println(t2.Hour())
	fmt.Println(t2.Minute())
	fmt.Println(t2.Second())
	fmt.Println(t2.Nanosecond())
	fmt.Println(t2.Zone())

	fmt.Println(time.Hour)
	fmt.Printf("%T\n", time.Hour)
	fmt.Println(time.Minute)
	fmt.Println((time.Second))
	fmt.Println(time.Millisecond)
	fmt.Println(time.Microsecond)
	fmt.Println(time.Nanosecond)

	d, _ := time.ParseDuration("2h30m")
	fmt.Println(d)
	t3 := time.Now()
	t3 = t3.Add(2*time.Minute + 15*time.Second)
	fmt.Println(t3)

	t4 := time.Date(2021, 7, 24, 0, 0, 0, 0, time.Local)
	t5 := time.Now()
	fmt.Println(t5)
	d2 := t4.Sub(t5)
	fmt.Println(d2)
	fmt.Println(t5.Before(t4))
	fmt.Println(t5.After(t4))
	fmt.Println(t4.Before(t5))
	fmt.Println(t4.After(t5))

	time.Sleep(10 * time.Second)
	fmt.Println("10秒後")
}
