package models

// User defines a person which is MiBici's user
type User struct {
	Age    int32  `json:"age,omitempty"`
	Gender string `json:"gender,omitempty"`
}
