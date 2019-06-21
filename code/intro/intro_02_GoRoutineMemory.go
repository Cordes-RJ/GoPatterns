package intro

import "sync"
import util "GoPatterns/code/util"

// If we look at the function below:
func intro02memoryA() {
	var wg sync.WaitGroup // this is a WaitGroup
	for i, item := range []string{"a", "b", "c", "d"} {
		wg.Add(1)
		go util.MakeAThing(item, i).DoWithWG(&wg)
	}
	wg.Wait()
}
