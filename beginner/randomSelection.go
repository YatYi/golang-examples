package main

import (
"fmt"
"math/rand"
"time"
)

//return a random number
func random(min, max int) int {
    rand.Seed(time.Now().Unix())
    return rand.Intn(max - min) + min
}

func main() {
	//make a list of string with 12 things in it
	s:=make([]string, 12)
	s[0]="Apple Swift"	
	s[1]="C#"	
	s[2]="F#"	
	s[3]="Google Go"	
	s[4]="Julia"	
	s[5]="Kotlin"	
	s[6]="Nim"	
	s[7]="R"	
	s[8]="Ruby"
	s[9]="Rust"	
	s[10]="Scala"	
	s[11]="Typescript"	

	//create a dictionary that map integer as key and a boolean as values
	c:=map[int]bool{
		0:false,
		1:false,
		2:false,
		3:false,
		4:false,
		5:false,
		6:false,	
		7:false,
		8:false,
		9:false,
		10:false,
		11:false,
	}

	var counter int=0
	var i int=0
	//randomly print out the thing in the list without duplicate
	for{
		myrand:=random(0, 12)
		if c[myrand]{
			continue
		} else{
			c[myrand]=true
			fmt.Println(s)
    		fmt.Println(s[myrand])
    		i++
    		counter++
		}

		if counter>=12{
			break
		}

	}
}
