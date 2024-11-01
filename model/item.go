package model

type Item struct {
	ID string `gorm:"type:uuid;primaryKey"`

	Title string
}
