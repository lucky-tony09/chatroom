package chatroom

import "fmt"

func Start() {
	m := make(chan int, 4)
	for i := range m {
		fmt.Println("0", i)
	}
}
