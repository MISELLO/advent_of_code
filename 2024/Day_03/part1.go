package main

import (
	"bufio"
	"fmt"
	"os"
	//"strconv"
)

type tStatus int

const (
	Letter_m tStatus = iota
	Letter_u
	Letter_l
	OpenParenthesis
	Number_1
	Comma
	Number_2
	CloseParenthesis
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please, provide just one file to analize.")
		os.Exit(0)
	}
	fmt.Println("Opening file", os.Args[1])

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer f.Close()

	fmt.Println("File", os.Args[1], "opened")

	reader := bufio.NewReader(f)
	var status tStatus
	var num1, num2, sum int
	var found bool

	// Load input
	fmt.Println("Reading character by character ...")
	for {
		char, err := reader.ReadByte()
		if err != nil {
			if err.Error() == "EOF" {
				fmt.Println("End Of File reached!")
				break
			}
			fmt.Println(err.Error())
			os.Exit((1))
		}
		found = process(&status, &num1, &num2, char)
		if found {
			sum += num1 * num2
		}
	}

	// Result
	fmt.Printf("The sum of all the results of the valid multiplications is \033[1m%d\033[0m.\n", sum)
}

func process(s *tStatus, n1, n2 *int, c byte) bool {

	fmt.Printf("%s", string(c))

	// We simplify this switch to avoid gocyclo over 15

	switch *s {
	case Letter_m, Letter_u, Letter_l:
		return processLetters(s, n1, n2, c)
	case OpenParenthesis:
		return processOpenParenthesis(s, n1, n2, c)
	case Number_1, Number_2:
		return processNumbers(s, n1, n2, c)
	}
	return false
}

func processLetters(s *tStatus, n1, n2 *int, c byte) bool {
	switch *s {
	case Letter_m:
		if c == 'm' {
			*s = Letter_u
		}
		return false
	case Letter_u:
		if c == 'u' {
			*s = Letter_l
			return false
		} else if c == 'm' {
			*s = Letter_u // keep the state
			return false
		}
		*s = Letter_m
		return false
	case Letter_l:
		if c == 'l' {
			*s = OpenParenthesis
			return false
		} else if c == 'm' {
			*s = Letter_u
			return false
		}
		*s = Letter_m
		return false
	}
	return false
}

func processOpenParenthesis(s *tStatus, n1, n2 *int, c byte) bool {
	switch *s {
	case OpenParenthesis:
		if c == '(' {
			*n1 = 0
			*n2 = 0
			*s = Number_1
			return false
		} else if c == 'm' {
			*s = Letter_u
			return false
		}
		*s = Letter_m
		return false
	}
	return false
}

func processNumbers(s *tStatus, n1, n2 *int, c byte) bool {
	switch *s {
	case Number_1:
		if c >= '0' && c <= '9' {
			*n1 = *n1*10 + int(c-'0')
			return false
		} else if c == ',' && *n1 > 0 && *n1 < 1000 {
			*s = Number_2
			return false
		} else if c == 'm' {
			*s = Letter_u
			return false
		}
		*s = Letter_m
		return false
	case Number_2:
		if c >= '0' && c <= '9' {
			*n2 = *n2*10 + int(c-'0')
			return false
		} else if c == ')' && *n2 > 0 && *n2 < 1000 {
			*s = Letter_m
			fmt.Printf(" --> %d x %d = %d\n", *n1, *n2, *n1**n2)
			return true
		} else if c == 'm' {
			*s = Letter_u
			return false
		}
		*s = Letter_m
		return false
	}
	return false
}
