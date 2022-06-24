package model

type StringError struct {
	ErrorString string `json:"error"`
}

func (s *StringError) Error() string {
	return s.ErrorString
}
