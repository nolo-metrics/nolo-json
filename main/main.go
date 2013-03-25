package main

import (
	"nolo"
	"os"
)

func collect(name, input string) (items []item) {
	l := lex(name, input)
	for {
		item := l.nextItem()
		items = append(items, item)
		if item.typ == itemEOF || item.typ == itemError {
			break
		}
	}
	return
}

func main() {
	input, _ := ioutil.ReadAll(os.Stdin)

	items := collect("bob", string(input))
	for i := 0; i < len(items); i++ {
		item := items[i]
		fmt.Println(item.String())
	}
}
