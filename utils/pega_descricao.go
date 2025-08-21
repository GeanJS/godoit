// Package utils
package utils

import (
	"bufio"
	"fmt"
	"os"
)


func PegaDescricao() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println(input)
	return input
	
}
