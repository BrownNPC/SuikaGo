package main

type CInput struct {
	Left  bool
	Right bool
	Drop  bool
}

func NewCInput() *CInput {
	return &CInput{}
}
