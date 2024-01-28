package model


type PatientBeds struct {
    ID        uint `gorm:"primaryKey"`
    All      int 
    Available     int
}