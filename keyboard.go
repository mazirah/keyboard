//Package keyboard reads keyboard input
package keyboard

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

//GetFloat takes a string from Stdin trims it and coverts it float64, then returns it
func GetFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	number, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatal(err)
	}
	return number, nil
}
