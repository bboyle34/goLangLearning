package main

import (
    "fmt"
    "time"
    "bytes"
    "regexp"
)

func main() {
    fmt.Println("Hello Brendan, it is", time.Now().Format("2006-01-02 3:4:5 PM"))
    fmt.Println()

    // this tests whether a pattern matches a string
    match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
    fmt.Println(match)
    fmt.Println()

    // above we used a string pattern directly, but for other
    // regexp tasks you'll need to Compile an optimized Regexp struct
    r, _ := regexp.Compile("p([a-z]+)ch")

    // many methods are available on these structs.
    // here's a match test like we saw earlier
    fmt.Println("--MatchString--")
    fmt.Println(r.MatchString("peach"))
    fmt.Println()

    // this finds the match for the regexp
    fmt.Println("--FindString--")
    fmt.Println(r.FindString("peach punch"))
    fmt.Println()

    // this also finds the first match but returns the start and
    // end indexes for the match instead of the matching text
    fmt.Println("--FindStringIndex--")
    fmt.Println(r.FindStringIndex("peach punch"))
    fmt.Println()

    // the Submatch variants include information about both the
    // whole pattern matches and the submatches within those matches.
    // for example this will return information for both 
    // p([a-z]+)ch and ([a-z]+)
    fmt.Println("--FindSringSubmatch--")
    fmt.Println(r.FindStringSubmatch("peach punch"))
    fmt.Println()

    // similarly this will return information about the indexes of
    // matches and submatches
    fmt.Println("--FindStringSubmatchIndex--")
    fmt.Println(r.FindStringSubmatchIndex("peach punch"))
    fmt.Println()

    // the All variants of these functions apply to all matches
    // in the input, not just the first.
    // for example to find all matches for a regexp
    fmt.Println("--FindAllString--")
    fmt.Println(r.FindAllString("peach punch pinch", -1))
    fmt.Println()

    // these All variants are available for the other functions we 
    // saw above as well
    fmt.Println("--FindAllStringSubmatchIndex--")
    fmt.Println(r.FindAllStringSubmatchIndex("peach pinch punch", -1))
    fmt.Println()

    // providing a non negative integer as the second argument
    // to these functions will limit the number of matches
    fmt.Println("--FindAllString--")
    fmt.Println(r.FindAllString("peach punch pinch", 2))
    fmt.Println()

    // our examples above had string arguments and used names like MatchString.
    // we can also provide []bye arguments and drop
    // String from the function name
    fmt.Println("--Byte Match--")
    fmt.Println(r.Match([]byte("peach")))
    fmt.Println()

    // when creating global variables with regular expressions
    // you can use the MustCompile variation of Compile.
    // MustCompile panics instaeadof returnign an error, which
    // makes it safer to use for global variables
    fmt.Println("--MustCompile--")
    r = regexp.MustCompile("p([a-z]+)ch")
    fmt.Println(r)
    fmt.Println()

    // the regexp package can also be used to replace subsets
    // of strings with other values
    fmt.Println("--ReplaceAllString--")
    fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))
    fmt.Println()

    // the Func variant allows you to transform mathed text
    // with a given function
    fmt.Println("--ReplaceAllFunc--")
    in := []byte("a peach")
    out := r.ReplaceAllFunc(in, bytes.ToUpper)
    fmt.Println(string(out))

}
