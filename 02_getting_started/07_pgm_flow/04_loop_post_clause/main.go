package main

func main() {
	i := 0
	for ; i < 5; i++ { // "i++" written after condition gets executed after every iteration of loop
		// you cannot write post clause loops with just condition and post clause "for i < 5; i++"
		// you'll need to add an initial clause which gets executed only one before the looping starts by checking with condition clause
		// if you don't have any initial clause, you can leave it blank as shown above but with a semicolon before condition clause
		println(i)
	}
	println(i) // 5

	for j := 5; j > -1; j-- { // you can combine looping variable creation and post clause along with the looping condition
		println(j)
	}
	/*
		println(j) // will throw compilation error
		// "j" was declared within the loop clause, hence its scope is valid only within the loop, and it gets deallocated after the loop finishes
	*/
}
