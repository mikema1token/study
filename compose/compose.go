package main

type A struct {
	a string
}

type B struct {
	a string
	//b A
}

type C struct {
	a A
	b B
}

func main(){
	c :=C{}
	println(c.a)
}
