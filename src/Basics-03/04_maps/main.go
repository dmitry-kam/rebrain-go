package main

import (
	"fmt"
	"sort"
)

func main() {
	var book map[string]int

	// map[KeyType]ValueType
	// 	KeyType may be any type that is comparable (more on this later)
	// 	ValueType may be any type at all, including another map!
	var author map[string]string = map[string]string{
		"name":     "Stephen",
		"lastName": "King",
	}

	reader := map[string]string{
		"name":     "John",
		"lastName": "Smith",
		"city":     "NY",
	}

	fmt.Println(book)
	fmt.Println(author)
	fmt.Println(reader)

	months := map[string]int{
		"Jan": 31,
		"Feb": 28,
		"Mar": 31,
	}
	fmt.Println(len(months)) // 3

	items := map[int]map[string]int{
		2020: {
			"books":      10,
			"periodical": 8,
		},
		2019: {
			"books":      12,
			"periodical": 10,
		},
	}

	fmt.Println(items[2020]) // map[books:10 periodical:8]

	items[2018] = map[string]int{
		"books": 5,
	}
	fmt.Println(items) // map[2018:map[books:5] 2019:map[books:12 periodical:10] 2020:map[books:10 periodical:8]]

	book = map[string]int{}
	fmt.Println(book["Book1"]) // 0, not error
	a, ok := book["Book1"]
	fmt.Println(a, ok) // 0 false

	delete(items, 2020)
	fmt.Println(items) // map[2018:map[books:5] 2019:map[books:12 periodical:10]]

	var book1 map[string]int
	fmt.Println(book1) // map[]
	//book1["Book1"] = 2        // panic: assignment to entry in nil map
	fmt.Println(book1 == nil) // true

	//////////////////////////////
	var book2 map[string]int
	reader2 := map[string]string{}
	fmt.Println(book2 == nil, len(book2))     // true 0
	fmt.Println(reader2 == nil, len(reader2)) // false 0

	///////////////
	book3 := make(map[string]int, 100)
	book3["Book1"] = 10
	book3["Book2"] = 6
	fmt.Println(book3, len(book3))

	///////////////////////////////////
	book4 := map[string]int{
		"Book2": 2,
		"Book4": 4,
		"Book1": 1,
		"Book3": 3}

	keys := make([]string, 0, len(book4))
	// always different order
	for k, v := range book4 {
		fmt.Println(k, v)
	}

	// ordered keys
	for k := range book4 {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, v := range keys {
		fmt.Println(book4[v]) // 1 2 3 4
	}

	fmt.Println(len(keys), cap(keys)) // 4 4

	//////////////////////////////
	type Node struct {
		Next  *Node
		Value interface{}
	}
	//var first *Node
	third := &Node{Value: "C", Next: nil}
	second := &Node{Value: "B", Next: third}
	first := &Node{Value: "A", Next: second}

	visited := make(map[*Node]bool)
	for n := first; n != nil; n = n.Next {
		if visited[n] {
			fmt.Println("cycle detected")
			break
		}
		visited[n] = true
		fmt.Println(n.Value)
	}

	////////////////////////////////

	m := map[string]int{}
	for _, word := range []string{"hello", "world", "from", "the",
		"best", "language", "in", "the", "world"} {
		m[word]++
	}
	for k, v := range m {
		println(k, v)
	}

	////////////////////////////////
	m1 := make(map[int]int)
	m1[10] = 15
	println("m1[10] before foo =", m1[10]) // 15
	foo(m1)                                // &m1
	println("m1[10] after foo =", m1[10])  // 10

	var m2 map[int]int
	fn(m2)
	fmt.Println("m2 == nil in main?:", m2 == nil)
	fn(m1)
	fmt.Println("m1 == nil in main?:", m1 == nil)

	///////////////////////////////
	//a := &m[1]
	//fmt.Println(m[1], *a)

}

func foo(m map[int]int) {
	m[10] = 10
}

func fn(m2 map[int]int) {
	//m2 = make(map[int]int)
	fmt.Println("m2 == nil in fn?:", m2 == nil, m2)
}
