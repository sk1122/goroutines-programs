// Map Function - (string) => dict with occurences of each character
// Reduce Function - ([] dict) => number (adding all the occurences)
// This should be done concurrently using goroutine
// 'satyam', 'programmer' => {s:1, a:2, t:1, y:1, m:1}, {p:1, r:3, o:1, g:1, a:1, m:2, e:1}
// => {s:1, a:3, t:1, y:1, m:3, p:1, r:3, o:1, g:1, e:1}

package main

import (
    "fmt"
    "sync"
)

func mapx(s string, x chan map[string]uint) {
    m := make(map[string]uint)
    for i, a := range s {
	fmt.Println(i, string(a))
	m[string(a)] += 1
    }

    x <- m
}

func reduce(a map[string]uint, b map[string]uint) map[string]uint {
    m := make(map[string]uint)

    for i, x := range a {
	fmt.Println(i, string(i), x, a[string(i)], "s")
	m[string(i)] += a[string(i)]
    }

    for i, x := range b {
	fmt.Println(x, "x")
	m[string(i)] += b[string(i)]
    }

    return m
}

func main() {
    a := "satyam"
    b := "programmer"

    m := make(chan map[string]uint)
    x := make(chan map[string]uint)

    var done sync.WaitGroup
    done.Add(1)

    go func() {
	defer done.Done()
	mapx(a, m)
    }()
    fmt.Println("qw")
    var done1 sync.WaitGroup
    done1.Add(1)
    go func() {
	defer done1.Done()
	mapx(b, x)
    }()

    done.Wait()
    done1.Wait()

    f := <-m
    d := <-x

   fmt.Println(f, d) 

   // fmt.Println(reduce(m, x))
}
