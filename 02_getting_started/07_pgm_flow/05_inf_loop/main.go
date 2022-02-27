package main

func main() {
	for i := 0; ; i++ { // you can skip the conditional clause
		println(i)
		break // commenting this will cause value of "i", 0 to be printed forever
	}

	for { // you can keep all the clauses empty and with just semicolons
		break
	}

	i := 0
	i++
	for { // you can just have a plain "for"
		println(i)
		break // commenting this will cause value of "i", 0 to be printed forever
	}

}
