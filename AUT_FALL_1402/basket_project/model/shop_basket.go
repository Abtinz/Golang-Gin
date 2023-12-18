package model

//we need to track user and basket modifications
import "time"

type Basket struct {
    ID        uint `gorm:"primaryKey"`
    CreatedAt time.Time
    UpdatedAt time.Time
    Data      string `gorm:"size:2048"`
    State     string `gorm:"size:10"`
    UserID   uint
}