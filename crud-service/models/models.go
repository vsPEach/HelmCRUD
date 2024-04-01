package models

type User struct {
	ID        uint64 `json:"id,omitempty" gorm:"primaryKey"`
	Username  string `json:"username"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}

type CustomError struct {
	Code int64  `json:"code"`
	Text string `json:"text"`
}
