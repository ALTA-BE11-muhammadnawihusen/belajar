package main

import (
	"fmt"
	"reflect"
	"time"
	// "time"
)

func main() {
	now := time.Now()
	fmt.Println(now)

	date := time.Date(2022, 8, 14, 10, 10, 10, 0, time.Local)
	fmt.Println(date)

	// test := now.Weekday().String() -> menghasilkan time.weekday
	// test := now.Local() -> menghasilkan time.Time
	// test := now.Location() -> menghasilkan *time.Location
	// test := now.IsZero()
	// test := now.Unix()
	test := now.String() // sama seperti time now tapi dalam bentuk string

	refv := reflect.ValueOf(test)
	fmt.Println("type", refv.Type())
	// refv2 := reflect.ValueOf(test2)
	// fmt.Println("type", refv2.Type())

	fmt.Println(test)
	// fmt.Println(test2)

}

// gunakan .String untuk type data yang menghasilkan type tidak jelas seperti now.weekday, now.location
