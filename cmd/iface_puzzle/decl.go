package main

import "context"

type Foo interface {
	DoSomethingUseful(ctx context.Context) error
}

var _ Foo = new(SomeStruct)

type SomeStruct struct{}

func (s *SomeStruct) DoSomethingUseful(ctx context.Context) error {
	return nil
}
