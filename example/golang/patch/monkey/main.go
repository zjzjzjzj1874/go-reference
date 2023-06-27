package main

import (
	"fmt"
	"os"
	"strings"

	"bou.ke/monkey"
)

// Attention:这个在M1上没法执行,因为架构不支持
func main() {
	fmt.Println("what the hell?")
	monkey.Patch(fmt.Println, func(a ...interface{}) (n int, err error) {
		s := make([]interface{}, len(a))
		for i, v := range a {
			s[i] = strings.Replace(fmt.Sprint(v), "hell", "*bleep*", -1)
		}
		return fmt.Fprintln(os.Stdout, s...)
	})
	fmt.Println("what the hell?") // what the *bleep*?
}
