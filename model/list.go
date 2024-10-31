package model

type List struct {
	ID   string `gorm:"type:uuid;primaryKey"`
	Name string
}
