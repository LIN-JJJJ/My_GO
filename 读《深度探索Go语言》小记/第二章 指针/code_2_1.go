package main

func main() {
	n := 10
	println(read(&n))
}

// go:noinline
func read(p *int) (v int) {
	v = *p
	return
}
