package model

type BloodType struct {
	BloodID  uint   `gorm:"primaryKey; autoIncrement"`
	Name     string `gorm:"unique"`
	Capacity int
}
