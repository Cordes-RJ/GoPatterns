package intro

import (
	"sync"
	"time"
)

// this is a bit of a breakaway from the themes of previous examples, but this
// is still in the spirit of sharing memory and distributed performance. In
// general computer science there are a number of idioms or common use patterns
// to achieve certain goals, one of which is referred to as the object pool
// pattern. Object pools, in short, are used when the instantiation or
// destruction of a class is expensive. A good example of this is a database
// connection. We want to re-use this kind of asset not destroy and remake it.
// Sync offers a concurrent safe implementation through the pool method.

// // NOTE: I'll do a comparison func at a later date, time permitting. [x]
// Intro06ExpensiveObj is a simple object that is produced from an expensive
// maker function.
type Intro06ExpensiveObj struct {
	A int
}

// Intro06MakeExpensiveObj makes an expensive object
func Intro06MakeExpensiveObj(cost int) Intro06ExpensiveObj {
	time.Sleep(time.Nanosecond * time.Duration(cost))
	return Intro06ExpensiveObj{A: cost}
}

// Intro06Pool is an example use-case of pool
func Intro06Pool(poolInitsize, workerCt, loops, cost int) {
	// pool has an exported attribute, "New", which contains a func that returns
	// a blank interface.
	pool := &sync.Pool{
		New: func() interface{} {
			obj := Intro06MakeExpensiveObj(cost)
			return &obj
		},
	}
	// now we seed the pool with a few structures. We do instantiation here
	for i := 0; i < poolInitsize; i++ {
		pool.Put(pool.New())
	}
	var wg sync.WaitGroup
	do := func() {
		defer wg.Done()
		//startTime := time.Now()
		obj := pool.Get().(*Intro06ExpensiveObj)
		defer pool.Put(obj)
		var x int
		for i := 0; i < 1000; i++ {
			x = obj.A
		}
		//fmt.Println(time.Now().Sub(startTime).Nanoseconds())
		_ = x
	}
	for loop := 0; loop < loops; loop++ {
		for i := 0; i < workerCt; i++ {
			wg.Add(1)
			go do()
		}
		wg.Wait()
	}

}

// Intro06PoolComparison is a comparison against the pool use-case above
func Intro06PoolComparison(workerCt, loops, cost int) {
	var wg sync.WaitGroup
	do := func() {
		defer wg.Done()
		//startTime := time.Now()
		obj := Intro06MakeExpensiveObj(cost)
		var x int
		for i := 0; i < 1000; i++ {
			x = obj.A
		}
		//fmt.Println(time.Now().Sub(startTime).Nanoseconds())
		_ = x
	}
	for loop := 0; loop < loops; loop++ {
		for i := 0; i < workerCt; i++ {
			wg.Add(1)
			go do()
		}
		wg.Wait()
	}
}

// this is a very contrived example focusing mostly on runtime. While it does
// decrease runtime by an order of magnitude, it does not display the range of
// its effects in an appropriate manner. Pool not only reduces runtime but also
// reduces memory allocation and cpu burden.
