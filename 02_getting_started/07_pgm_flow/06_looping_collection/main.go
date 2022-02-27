package main

func main() {
	slice := []int{1, 2, 3}
	for i := 0; i < len(slice); i++ { // too ugly and too much code just to iterate
		println(slice[i])
	}

	for idx, ele := range slice { // "range" returns an index and element at that index at each iteration
		println(idx, ele)
	}

	wellKnownPorts := map[string]int{"http": 80, "https": 443}
	for k, v := range wellKnownPorts {
		println(k, v)
	}

	for k := range wellKnownPorts { // you can take only the first value returned by "range" every iteration and ignore the rest without using write only variable
		// however you cannot do the same with a function return for which you must use "_" to ignore the return value
		println(k)
	}

	for _, v := range wellKnownPorts { // write only variable "_" is used to ignore the first element of the return from "range"
		println(v)
	}
}
