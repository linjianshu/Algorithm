package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Text()
	s := ""
	fmt.Scanf("%s\n", &s)
	b := false
	if s[0] == '-' {
		b = true
		s = s[1:]
	} else if s[0] == '+' {
		b = false
		s = s[1:]
	}

	ans := ""
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= 9 {
			ans += string(s[i])
		} else {
			break
		}
	}

	res := 0
	atoi, err := strconv.Atoi(ans)
	if err != nil {
		fmt.Println(err)
	}
	res = atoi
	if b {
		res = -res
	}
	fmt.Println(res)
}
