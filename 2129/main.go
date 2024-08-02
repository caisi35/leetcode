package main

import "fmt"

type Math struct {
	x, y int
}

var m = map[string]Math{
	"foo": Math{x:2, y:3},
}

func main() {
	t := m["foo"]
	t.x = 4
	m["foo"] = t
	fmt.Println(m["foo"].x)
	fmt.Println(capitalizeTitle("cAapiTalIze as TitLE"))
	fmt.Println(capitalizeTitle("L hV"))
}

func capitalizeTitle(title string) string {
	n := len(title)
	ans := make([]byte, n)
	word := 0
	for i := n - 1; i >= 0; i-- {
		// fmt.Println(title[i], 'A' + 32)

		if title[i] >= 'A' && title[i] <= 'Z' {
			ans[i] = title[i] + 32 // 全小写
		} else {
			ans[i] = title[i] // 全小写
		}
		if title[i] != ' ' {
			word += 1
		}

		if title[i] == ' ' && word > 2 {
			ans[i+1] = ans[i+1] - 32
			word = 0
		} else if word > 2 && i == 0 {
			ans[i] = ans[i] - 32
		}
		if title[i] == ' ' {
			word = 0
		}

	}
	return string(ans)
}
