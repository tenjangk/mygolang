package main

import (
	"fmt"
	"sort"
)

type user struct {
	Name string
	No   int
}

type byNo []user

func (a byNo) Len() int           { return len(a) }
func (a byNo) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byNo) Less(i, j int) bool { return a[i].No < a[j].No }

type byName []user

func (b byName) Len() int           { return len(b) }
func (b byName) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b byName) Less(i, j int) bool { return b[i].Name < b[j].Name }

func main() {

	u1 := user{
		Name: "tenz",
		No:   7834562321,
	}
	u2 := user{
		Name: "zen",
		No:   6834562321,
	}
	users := []user{u1, u2}

	fmt.Println(users)
	fmt.Println("--------------------------")
	sort.Sort(byNo(users))
	fmt.Println("sort by int")
	fmt.Println(users)
	fmt.Println("--------------------------")
	fmt.Println("sort by string")
	sort.Sort(byName(users))
	fmt.Println(users)

	s := []int{5, 45, 67, 23, 24, 78, 17}
	l := []string{"djfg", "jdkfh", "jan", "tjhew"}

	fmt.Println("--------------------------")

	sort.Ints(s)
	fmt.Println(s)
	sort.Strings(l)
	fmt.Println(l)

}
