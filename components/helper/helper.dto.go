package helper

type TestStruct struct {
	Name    string `json:"name" `
	Surname string `json:"surname" validate:"required"`
}
