package main

func main() {
	i := 0
	for i < 5 { // loop keeps running until "i" is >= 5
		println(i) // until now, we have used "fmt.println" which internally uses "println" and is way more sophisticated than it
		// for basic debugging, just "println" shall suffice
		i++ // if this didn't exist, loop would continue to run forever as "i" would be < 5 forever
	}

	for i < 10 { // "i" is 5 when it first reaches this line
		println(i)
		break // causes the loop that it was declared in to terminate immediately hence, "i" is not going to loop several times, but just one time
	}

	for i < 10 { // "i" is 5 when it first reaches this line
		println(i)
		i++
		continue       // continue causes whatever is below it to be skipped only during the iteration during which it was executed
		println(i * 2) // in every iteration "continue" above causing this to be skipped
	}
}
