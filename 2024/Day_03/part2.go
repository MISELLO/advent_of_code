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
	Letter_d
	Letter_o
	Letter_n
	Apostrophe
	Letter_t
	OpenParMul
	OpenParDo
	OpenParDont
	Number_1
	Comma
	Number_2
	CloseParMul
	CloseParDo
	CloseParDont
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

	var do bool = true

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
		found = process(&status, &num1, &num2, char, &do)
		if found && do {
			sum += num1 * num2
		}
	}

	// Result
	fmt.Printf("The sum of all the results of the valid multiplications is \033[1m%d\033[0m.\n", sum)
}

func process(s *tStatus, n1, n2 *int, c byte, do *bool) bool {

	fmt.Printf("%s", string(c))

	// We simplify this switch to avoid gocyclo over 15

	switch *s {
	case Letter_m, Letter_u, Letter_l:
		return processMul(s, n1, n2, c, do)
	case Letter_o, Letter_n, Letter_t:
		return processDont(s, n1, n2, c, do)
	case Apostrophe, OpenParMul, OpenParDont:
		return processApostropheAndOpenParenthesis(s, n1, n2, c, do)
	case CloseParDo, CloseParDont:
		return processCloseParenthesis(s, n1, n2, c, do)
	case Number_1, Number_2:
		return processNumbers(s, n1, n2, c, do)
	}
	return false
}

func processMul(s *tStatus, n1, n2 *int, c byte, do *bool) bool {
	switch *s {
	case Letter_m:
		if c == 'm' {
			*s = Letter_u
			return false
		} else if c == 'd' {
			*s = Letter_o
			return false
		}
		*s = Letter_m
		return false
	case Letter_u:
		if c == 'u' {
			*s = Letter_l
			return false
		} else if c == 'm' {
			*s = Letter_u // keep the state
			return false
		} else if c == 'd' {
			*s = Letter_o
			return false
		}
		*s = Letter_m
		return false
	case Letter_l:
		if c == 'l' {
			*s = OpenParMul
			return false
		} else if c == 'm' {
			*s = Letter_u
			return false
		} else if c == 'd' {
			*s = Letter_o
			return false
		}
		*s = Letter_m
		return false
	}
	return false
}

func processDont(s *tStatus, n1, n2 *int, c byte, do *bool) bool {
	switch *s {
	case Letter_o:
		if c == 'o' {
			*s = Letter_n
			return false
		} else if c == 'm' {
			*s = Letter_u
			return false
		} else if c == 'd' {
			*s = Letter_o
			return false
		}
		*s = Letter_m
		return false
	case Letter_n:
		if c == 'n' {
			*s = Apostrophe
			return false
		} else if c == '(' {
			*s = CloseParDo
			return false
		} else if c == 'm' {
			*s = Letter_u
			return false
		} else if c == 'd' {
			*s = Letter_o
			return false
		}
		*s = Letter_m
		return false
	case Letter_t:
		if c == 't' {
			*s = OpenParDont
			return false
		} else if c == 'm' {
			*s = Letter_u
			return false
		} else if c == 'd' {
			*s = Letter_o
			return false
		}
		*s = Letter_m
		return false
	}
	return false
}

func processApostropheAndOpenParenthesis(s *tStatus, n1, n2 *int, c byte, do *bool) bool {
	switch *s {
	case Apostrophe:
		if c == '\'' {
			*s = Letter_t
			return false
		} else if c == 'm' {
			*s = Letter_u
			return false
		} else if c == 'd' {
			*s = Letter_o
			return false
		}
		*s = Letter_m
		return false
	case OpenParMul:
		if c == '(' {
			*n1 = 0
			*n2 = 0
			*s = Number_1
			return false
		} else if c == 'm' {
			*s = Letter_u
			return false
		} else if c == 'd' {
			*s = Letter_o
			return false
		}
		*s = Letter_m
		return false
	case OpenParDont:
		if c == '(' {
			*s = CloseParDont
			return false
		} else if c == 'm' {
			*s = Letter_u
			return false
		} else if c == 'd' {
			*s = Letter_o
			return false
		}
		*s = Letter_m
		return false
	}
	return false
}

func processCloseParenthesis(s *tStatus, n1, n2 *int, c byte, do *bool) bool {
	switch *s {
	case CloseParDo:
		if c == ')' {
			*s = Letter_m
			*do = true
			return false
		} else if c == 'm' {
			*s = Letter_u
			return false
		} else if c == 'd' {
			*s = Letter_o
			return false
		}
		*s = Letter_m
		return false
	case CloseParDont:
		if c == ')' {
			*s = Letter_m
			*do = false
			return false
		} else if c == 'm' {
			*s = Letter_u
			return false
		} else if c == 'd' {
			*s = Letter_o
			return false
		}
		*s = Letter_m
		return false
	}
	return false
}

func processNumbers(s *tStatus, n1, n2 *int, c byte, do *bool) bool {
	switch *s {
	case Number_1:
		if digit(c) {
			*n1 = *n1*10 + int(c-'0')
			return false
		} else if c == ',' && *n1 > 0 && *n1 < 1000 {
			*s = Number_2
			return false
		} else if c == 'm' {
			*s = Letter_u
			return false
		} else if c == 'd' {
			*s = Letter_o
			return false
		}
		*s = Letter_m
		return false
	case Number_2:
		if digit(c) {
			*n2 = *n2*10 + int(c-'0')
			return false
		} else if c == ')' && *n2 > 0 && *n2 < 1000 {
			*s = Letter_m
			fmt.Printf(" --> %d x %d = %d (do=%v)\n", *n1, *n2, *n1**n2, *do)
			return true
		} else if c == 'm' {
			*s = Letter_u
			return false
		} else if c == 'd' {
			*s = Letter_o
			return false
		}
		*s = Letter_m
		return false
	}
	return false
}

func digit(c byte) bool {
	return c >= '0' && c <= '9'
}
