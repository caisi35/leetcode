package main

import (
	"fmt"
	"sort"
)

type student struct {
	Name string
	Age  int
}

func main() {
	//定义map
	m := make(map[string]*student)

	//定义student数组
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}

	//将数组依次添加到map中
	for i, stu := range stus {
		m[stu.Name] = &stus[i]
	}

	//打印map
	for k, v := range m {
		fmt.Println(k, "=>", v.Name)
	}
	allIp := []string{"192.1.1.1", "192.1.1.2", "192.1.1.3"}
	subnetGateway := "192.1.1.2"
	firstIp := "192.1.1.1"
	lastIp := "192.1.1.255"
	fmt.Println(sort.SearchStrings(allIp, subnetGateway))
	if subnetGateway != firstIp && subnetGateway != lastIp {
		fmt.Println("error")
	}
	resp := make([]string, 0, len(allIp))
	resp = append(resp, allIp[1:2]...)
	fmt.Println(resp)
}
