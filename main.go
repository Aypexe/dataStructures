package main

import (
	"datastructures/collection"
	linkedQueue "datastructures/collection/queue/linked"
	linkedStack "datastructures/collection/stack/linked"
	"fmt"
	"time"
)

func testCollection(c collection.Collectioner, name string) {
	fmt.Println("====== " + name + " ======\n" +
		"\n---- (NEWEST)\n" +
		c.String() +
		"---- (OLDEST)\n")
	time.Sleep(time.Second)

	res, err := c.Pop()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("popped:", res)
	}
	time.Sleep(time.Second)

	fmt.Println(
		"\n---- (NEWEST)\n" +
			c.String() +
			"---- (OLDEST)\n")

	fmt.Print("===================")
}

func main() {
	// h := linkedHashTable.New()
	// fmt.Println(h)

	q := linkedQueue.New()
	q.PushAll(1, 2, 3, 4, 5)
	testCollection(q, "Queue")

	fmt.Print("\n\n-- -- -- -- -- -- -- -- --\n\n")
	time.Sleep(4 * time.Second)

	s := linkedStack.New()
	s.PushAll(1, 2, 3, 4, 5)
	testCollection(s, "Stack")
}
