package main

import (
	"fmt"
)


//we created interface ans value within this interface
//any type electric device can be connect with this Socket() method on this socket board
type SocketBoard interface {
	Socket()
}
//we created smart tv struct
type SmartTv struct {
	TvName string
	MadeYr int32
}
//create smart tv socket
func (s SmartTv) Socket() {
	fmt.Printf("Tv name is %s and Made date %d", s.TvName, s.MadeYr)
}

// created describe function of that method
func Describe(s SocketBoard) {
	fmt.Printf("Type is %T, value is %v", s, s)
}

func main() {
	// created an instance of smart tv struct.
	samsung := SmartTv{
		TvName: "Samsung",
		MadeYr: 2021,
	}
	var s SocketBoard
	s = samsung
	Describe(s)
	fmt.Println()
	s.Socket()

}
