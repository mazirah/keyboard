//Package keyboard reads keyboard input
package keyboard

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)
//Fast error checking
func check(err error) {
	if err !=nil {
		log.Fatal(err)
	}
}

//GetFloat reads a float number from the keyboard. It returns the number and any error encountered
func GetFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	check(err)
	input = strings.TrimSpace(input)
	number, err := strconv.ParseFloat(input, 64)
	check(err)
	return number, nil
}

//GetInt reads an int number from the keyboard. It returns the number and any error encountered
func GetInt() (int, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	check(err)
	input = strings.TrimSpace(input)
	number, err := strconv.Atoi(input)
	check(err)
	return number, nil
}
