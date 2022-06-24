package vdfparser

type vpkError struct {
	s string
}

func (error *vpkError) Error() string {
	return error.s
}

var ErrInvalidDescriptorFile = vpkError{"invalid descriptor file"}
