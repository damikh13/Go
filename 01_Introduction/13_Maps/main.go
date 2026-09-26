package main

import (
	"fmt"
	"sort"
)

func addNew(mp map[string]int) {
	mp["new"] = 1000
}

type Person struct {
	Name string
	Age  int
}

func main() {
	ages := make(map[string]int)
	ages["John"] = 30
	fmt.Println(ages["John"])

	var mp2 map[string]int            // == nil
	fmt.Println("mp2[x] =", mp2["x"]) // 0
	// mp2["x"] = 10                     // нельзя

	ages["Alice"] = 10
	JohnAge := ages["John"]                  // 30
	missingPersonsAge := ages["missing key"] // 0
	fmt.Println(JohnAge)
	fmt.Println(missingPersonsAge)
	fmt.Println(len(ages))
	delete(ages, "John")
	fmt.Println(ages, len(ages)) // map[Alice:10] 1

	mp5 := map[string]int{"1st": 10, "2nd": 20, "3rd": 30, "4th": 40}
	mp5Keys := make([]string, 0, len(mp5))
	for k := range mp5 {
		mp5Keys = append(mp5Keys, k)
	}
	fmt.Println(mp5Keys)  // 4th 1st 3rd 2nd
	sort.Strings(mp5Keys) // 1st 2nd 3rd 4th
	for _, v := range mp5Keys {
		fmt.Printf("mp5[%q] = %v\n", v, mp5[v])
	}

	fmt.Println("mp5:", mp5) // mp5: map[1st:10 2nd:20 3rd:30 4th:40]
	addNew(mp5)
	fmt.Println("mp5:", mp5) // mp5: map[1st:10 2nd:20 3rd:30 4th:40 new:1000]

	personsInfo := map[string]Person{"1st": {"John", 30}}
	// personsInfo["1st"].Age = 31 // нельзя!
	p1 := personsInfo["1st"]
	p1.Age = 31
	personsInfo["1st"] = p1
	fmt.Println(personsInfo)

	personsInfoBetter := map[string]*Person{"1st": {"John", 30}}
	personsInfoBetter["1st"].Age = 31
	fmt.Println(personsInfoBetter["1st"])

	set := make(map[string]struct{})
	set["first item"] = struct{}{}
	val, ok := set["first item"]
	if !ok {
		fmt.Println("no such item in the set")
	}
	fmt.Println(val)
	fmt.Println(set)
}
