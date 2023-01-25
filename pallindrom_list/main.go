package main

import (
	"fmt"
	"practice_codes/data-structures/pallindrom_list/list"
	"practice_codes/data-structures/pallindrom_list/pallindrom"
)

func main() {
	lt := new(list.LinkedList)
	lt.AppendInList(1)
	lt.AppendInList(2)
	lt.AppendInList(2)
	lt.AppendInList(1)

	isPall := pallindrom.IsListPallindrom(lt)
	fmt.Printf("is list pallindrom :: %v \n", isPall)

}
