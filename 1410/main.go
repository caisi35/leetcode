package main

import (
	"fmt"
)

var p *int

func foo() (*int, error) {
	var i int = 5
	return &i, nil
}

func bar() {
	//use p
	fmt.Println(*p)
}

func main() {
	var err error
	p, err = foo()
	if err != nil {
		fmt.Println(err)
		return
	}
	bar()
	fmt.Println(*p)
	fmt.Println(entityParser("&amp; is an HTML entity but &ambassador; is not."))
	fmt.Println(entityParser("and I quote: &quot;...&quot;"))
	fmt.Println(entityParser("x &gt; y &amp;&amp; x &lt; y is always false"))
	fmt.Println(entityParser("leetcode.com&frasl;problemset&frasl;all"))

}

func entityParser(text string) string {
	ans := ""
	n := len(text)
	s := 0

	for i := 0; i < n; i++ {
		v := text[i]
		t := ""
		if s > 0 {
			s -= 1
			continue
		}
		if v == '&' {
			if n-i-1 >= 3 {
				if text[i+1:i+4] == "gt;" {
					t = ">"
					s = 3
				} else if text[i+1:i+4] == "lt;" {
					t = "<"
					s = 3
				}
			}
			if n-i-1 >= 4 {
				if text[i+1:i+5] == "amp;" {
					t = "&"
					s = 4
				}
			}
			if n-i-1 >= 5 {
				if text[i+1:i+6] == "apos;" {
					t = "'"
					s = 5
				} else if text[i+1:i+6] == "quot;" {
					t = `"`
					s = 5
				}
			}
			if n-i-1 >= 6 {
				if text[i+1:i+7] == "frasl;" {
					t = "/"
					s = 6
				}
			}
			if t != "" {
				ans += t
				continue
			}

		}
		ans += string(v)
	}
	return ans
}

// curl 'http://127.0.0.1:5555/api/traffic_mirror/mirror_filter/update/' -H "Content-Type: application/json" -d '{"os_id":"nm8","ct_user_id":"1000000590","mirror_filter_uuid":"mrfilter-82ryrscjxs1","name":"aa"}'

// { 'order_time': None,
// 'user_id': 1000000590,
// 'ct_user_platforms_id': 1216047039,
// 'dc_id': 3,
// 'deleted': False,
// 'uuid': 'mrfilter-82ryrscjxs',
// 'detail': '',
// 'extra_info': '',
// 'removed_time': None,
// 'is_paas': False,
// 'zone_id': '',
// 'jobid': None,
// 'name': 'aa1',
// 'description': ''}

// curl 'http://127.0.0.1:55552/api/peering/query/?os_id=nm8&ct_user_id=1000000662&query_content=a'