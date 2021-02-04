package main

import (
    "fmt"
    "testing"
)

// we'll be testing this simple implementation of an integer
// minimum. typically, the code we're testing would be in a
// source file named something like intutils.go, and the test
// file for it would then be named intutils_test.go
func IntMin(a, b int) int {
    if a < b {
        return a
    }
    return b
}

// a test is created by writing a function with a name 
// beggingin with Test
func TestIntMinBasic(t *testing.T) {
    ans := IntMin(2, -2)
    if ans != -2 {
        // t.Error* will report test failures but continue executing
        // the test. t.Fatal* will report test failures and 
        // stop the test immediately
        t.Errorf("IntMin(2, -2) = %d; want -2", ans)
    }
}

// writing tests can be repretitive, so it's idiomatic ot use a
// table-drive style, where test inputs and expected outputs
// are listed in a table and a single loop walks over them and 
// performs the test logic
func TestIntMinTableDriven(t *testing.T) {
    var tests = [] struct {
        a, b int
        want int
    }{
        {0, 1, 0},
        {1, 0, 0},
        {2, -2, -2},
        {0, -1, -1},
        {-1, 0, -1},
    }

    // t.Run enables running "subtests", one for each table entry
    // these are shown separately when executing go test -v
    for _, tt := range tests {
        testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
        t.Run(testname, func(t *testing.T) {
            ans := IntMin(tt.a, tt.b)
            if ans != tt.want {
                t.Errorf("got %d, want %d", ans, tt.want)
            }
        })
    }
}
