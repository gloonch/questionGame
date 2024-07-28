package entity

type User struct {
	ID          uint
	PhoneNumber string
	Name        string
	// Password will always be kept hashed
	Password string
}
