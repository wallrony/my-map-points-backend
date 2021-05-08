package models

type ApiError struct {
	Err string
}

func (a ApiError) AsMessage() map[string]string {
	return map[string]string{
		"Message": a.Err,
	}
}
