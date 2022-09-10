package main

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

type tt struct {
	t time.Time
	s string
}

func main() {
	n := 0
	fmt.Scanf("%d\n", &n)
	ss := make([]string, 0, n)
	m := make(map[string]bool)
	for i := 0; i < n; i++ {
		temp := ""
		fmt.Scanf("%s\n", &temp)
		if !m[temp] {
			ss = append(ss, temp)
			m[temp] = true
		}
	}

	ans := make([]tt, 0, 0)
	for i := 0; i < len(ss); i++ {
		split := strings.Split(ss[i], "/")
		for _, value := range split {
			if len(value) == 19 && value[4] == '-' && value[7] == '-' && value[10] == 'T' && value[13] == ':' && value[16] == ':' {
				parse, err := time.Parse("2006-01-02T15:04:05", value)
				if err != nil {

				} else {
					ans = append(ans, tt{
						t: parse,
						s: ss[i],
					})
					break
				}
			} else {

			}
		}
	}

	sort.Slice(ans, func(i, j int) bool {
		if ans[i].t != ans[j].t {
			return ans[i].t.Before(ans[j].t)
		} else if len(ans[i].s) != len(ans[j].s) {
			return len(ans[i].s) < len(ans[j].s)
		} else if ans[i].s != ans[j].s {
			return ans[i].s < ans[j].s
		}
		return false
	})
	for i := 0; i < len(ans); i++ {
		fmt.Println(ans[i].s)
	}

}
