package GoPattern

import (
	"fmt"
	"strconv"
	"sync"
)

type KillOrder struct {
	FromID string
}

func DaisyChain() {
	KillChan := make(chan KillOrder, 1)
	var wg sync.WaitGroup
	do := func(id int) {
		defer wg.Done()
		for {
			select {
			case x := <-KillChan:
				fmt.Printf("[%s->%d]", x.FromID, id)
				x.FromID = strconv.Itoa(id)
				KillChan <- x
				return
			}
		}
	}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go do(i)
	}
	KillChan <- KillOrder{"Main"}

	wg.Wait()
}
