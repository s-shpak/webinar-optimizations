package main

import "context"

func main() {
	s := &SomeStruct{}
	_ = doSomethingUseful(s)
}

func doSomethingUseful(i any) error {
	return i.(Foo).DoSomethingUseful(context.Background())
}
