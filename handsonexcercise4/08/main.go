package main

import "fmt"

func main() {
	m := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	m[`jangk_ten`] = []string{`kindness`, `painting`, `hiking`}

	for k, v := range m {
		fmt.Println("record for ", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}

	}

	delete(m, `no_dr`)
	fmt.Println(m[`no_dr`])

	if v, ok := m[`no_dr`]; ok { //use comma ok idiom to verify if deleted one still exist or not
		fmt.Println("value of no_dr:", v) //since it doesnt exist, bool for ok will be false so wont show anything
	}
	fmt.Println("")
	fmt.Println(m)
}
