package ttt

import "github.com/cosmos/cosmos-sdk/yyy"

type A struct {
	b yyy.B
}

func (a *A) BString() string {
	return a.b.String()
}

func NewA() A {
	return A{}
}
		