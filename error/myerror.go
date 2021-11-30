package main

import (
	"errors"
	"fmt"
)

func main(){
	makeError()
	fmt.Println("abcd")
	callMakePanic()
	fmt.Println("finish ok")
}

func makeError()error{
	return errors.New("hi,wo shi error")
}

func makePanic(){
	panic("haha,wo shi panic")
}

func callMakePanic(){
	defer func() {
		if err:=recover();err!=nil{
			fmt.Println("catch panic")
		}
	}()
	makePanic()
}