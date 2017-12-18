package dynaport

import (
	"fmt"
)

func ExampleGet() {
	ports := Get(3)
	fmt.Printf("%#v", ports)
}

func ExampleGetWithErr() {
	ports, err := GetWithErr(3)
	if err != nil {
		panic("should not have an error")
	}
	fmt.Println(ports)
}

func ExampleGetWithErr_err() {
	_, err := GetWithErr(100000)
	if err == nil {
		panic("should have an error")
	}
	fmt.Println(err.Error())
	// Output:
	// dynaport: block size is too small for ports requested
}
