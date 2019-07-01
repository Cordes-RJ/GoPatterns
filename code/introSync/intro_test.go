package introSync

import (
	"fmt"
	"testing"
	"time"
)

func TestIntro00(t *testing.T) {
	name := "intro_00_Routines"
	fmt.Printf("<%s>\n", name)
	defer fmt.Printf("\n</%s>\n", name)
	intro00GoRoutines()
	x := false
	if x {
		t.Error("error")
	}
}

func TestIntro01A(t *testing.T) {
	name := "intro_01_WaitGroups"
	fmt.Printf("<%s>\n", name)
	defer fmt.Printf("\n</%s>\n", name)
	intro01WaitGroupsA()
	x := false
	if x {
		t.Error("error")
	}
}

func TestIntro01B(t *testing.T) {
	name := "intro_01_WaitGroups"
	fmt.Printf("<%s>\n", name)
	defer fmt.Printf("\n</%s>\n", name)
	intro01WaitGroupsB()
	x := false
	if x {
		t.Error("error")
	}
}

func TestIntro02A(t *testing.T) {
	name := "intro_02_GoRoutineMemoryA"
	fmt.Printf("<%s>\n", name)
	defer fmt.Printf("\n</%s>\n", name)
	intro02memoryA()
	x := false
	if x {
		t.Error("error")
	}
}

func TestIntro02B(t *testing.T) {
	name := "intro_02_GoRoutineMemoryB"
	fmt.Printf("<%s>\n", name)
	defer fmt.Printf("\n</%s>\n", name)
	intro02memoryB()
	x := false
	if x {
		t.Error("error")
	}
}

func TestIntro02C(t *testing.T) {
	name := "intro_02_GoRoutineMemoryC"
	fmt.Printf("<%s>\n", name)
	defer fmt.Printf("\n</%s>\n", name)
	intro02memoryC()
	x := false
	if x {
		t.Error("error")
	}
}

func TestIntro03A(t *testing.T) {
	name := "intro_03_MutexA"
	fmt.Printf("<%s>\n", name)
	defer fmt.Printf("\n</%s>\n", name)
	intro03mutexA()
	x := false
	if x {
		t.Error("error")
	}
}

func TestIntro03B(t *testing.T) {
	name := "intro_03_MutexB"
	fmt.Printf("<%s>\n", name)
	defer fmt.Printf("\n</%s>\n", name)
	intro03mutexB()
	x := false
	if x {
		t.Error("error")
	}
}

func TestIntro04rwMutexA(t *testing.T) {
	name := "intro_04_rwMutexA"
	fmt.Printf("<%s>\n", name)
	defer fmt.Printf("\n</%s>\n", name)
	x := false
	if x {
		t.Error("error")
	}
	startTime := time.Now()
	for i := 0; i < 100; i++ {
		intro04rwMutexA()
	}
	fmt.Print(time.Now().Sub(startTime).Nanoseconds())
}

func TestIntro04rwMutexB(t *testing.T) {
	name := "intro_04_rwMutexB"
	fmt.Printf("<%s>\n", name)
	defer fmt.Printf("\n</%s>\n", name)
	x := false
	if x {
		t.Error("error")
	}
	startTime := time.Now()
	for i := 0; i < 10; i++ {
		intro04rwMutexB()
	}
	fmt.Print(time.Now().Sub(startTime).Nanoseconds())
}

func TestIntro05primitiveA(t *testing.T) {
	name := "intro_05_primitiveA"
	fmt.Printf("<%s>\n", name)
	defer fmt.Printf("\n</%s>\n", name)
	x := false
	if x {
		t.Error("error")
	}
	startTime := time.Now()
	for i := 0; i < 10; i++ {
		Intro05primitiveA(3)
	}
	fmt.Println()
	fmt.Print(time.Now().Sub(startTime).Nanoseconds())
}

func TestIntro05primitiveB(t *testing.T) {
	name := "intro_05_primitiveB"
	fmt.Printf("<%s>\n", name)
	defer fmt.Printf("\n</%s>\n", name)
	x := false
	if x {
		t.Error("error")
	}
	startTime := time.Now()
	for i := 0; i < 10; i++ {
		Intro05primitiveB(3)
	}
	fmt.Println()
	fmt.Print(time.Now().Sub(startTime).Nanoseconds())
}

func TestIntro05Cond(t *testing.T) {
	name := "intro_05_cond"
	fmt.Printf("<%s>\n", name)
	defer fmt.Printf("\n</%s>\n", name)
	x := false
	if x {
		t.Error("error")
	}
	startTime := time.Now()
	for i := 0; i < 5; i++ {
		Intro05Cond(3)
	}
	fmt.Println()
	fmt.Print(time.Now().Sub(startTime).Nanoseconds())
}

func TestIntro06Pool(t *testing.T) {
	name := "intro_06_Pool"
	fmt.Printf("<%s>\n", name)
	defer fmt.Printf("\n</%s>\n", name)
	x := false
	if x {
		t.Error("error")
	}
	wCt := 25
	runs := 5
	cost := 5000
	loops := 500
	startTime := time.Now()
	for i := 0; i < runs; i++ {
		Intro06Pool(25, wCt, loops, cost)
	}
	fmt.Print("\n      pool: ")
	fmt.Print(time.Now().Sub(startTime).Nanoseconds())
	startTime = time.Now()
	for i := 0; i < runs; i++ {
		Intro06PoolComparison(wCt, loops, cost)
	}
	fmt.Print("\nvs no pool: ")
	fmt.Print(time.Now().Sub(startTime).Nanoseconds())
}
