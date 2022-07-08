package utils

import (
	"bufio"
	"fmt"
	"os"
)

func ReadUserInput() {

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	fmt.Printf("user input is %v ", input)
}
