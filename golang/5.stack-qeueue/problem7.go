package main

import "fmt"

func main() {
	var n int

	fmt.Scanf("%d", &n)

	q := Queue{}

	for i := 0; i < n; i++ {
		var v int
		var u int

		fmt.Scanf("%d %d", &v, &u)

		if v == 1 {
			// push the value (u) to the queue ( enqueue )
			q.Enqueue(u)
		} else if v == 2 {
			// pop the value from the queue ( dequeue )
			q.Dequeue()
		}

	}

	fmt.Printf("\n %+v", q.in)
	fmt.Printf("\n %+v", q.out)
}
