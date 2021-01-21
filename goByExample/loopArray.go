package main

import f "fmt"

func pr() {
    f.Println()
}

func main() {

    //I want to create a 2d array and print out a big X
    pr()
    var twod [10][10]int
    var counter int
    counter = 0
    for j := 0; j < len(twod); j++ {
        for i := 0; i < len(twod[j]); i++ {
            //f.Print(twod[j][i], " ")
            if (i == counter) || (9-i == counter) {
                f.Print("1", " ") 
            } else {
                f.Print("0", " ")
            }
        }
        pr()
        counter += 1
    }
    pr()

    //also want the user to create a 2d array and print it out
    var input1, input2 int
    f.Print("Please input a number between 1-10: ")
    f.Scanln(&input1)
    f.Print("Please input another number between 1-10: ")
    f.Scanln(&input2)
    
    //f.Println(input1, input2)
     arr := make([intput1][input2]int)
    f.Println("Here is your array: ", arr)

}
