package something

import "fmt"

func Say() {
	fmt.Println("hello")
}

// 1byte = 8bit
// 00000000~11111111 =0~255 128+64+32+16+8+4+2+1 = 255

// 0 0 0 0 0 0 1 1
// 7 6 5 4 3 2 1 0

// 0 0 0 0 0 1 0 1

// 0 0 0 1 0 0 1 1
