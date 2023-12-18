package model

type User struct {
	UserID   uint   `gorm:"primaryKey; autoIncrement"`
	Username string `gorm:"unique"`
	Password string
}

//why i did not put token here?
//for high security principals in local server(my other dimension always follows hacking this poor dimension)