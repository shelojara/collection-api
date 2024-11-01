package model

type List struct {
	ID string `gorm:"type:uuid;primaryKey"`

	Name        string
	Description string

	Items    []ListItem   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Statuses []ListStatus `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type ListStatus struct {
	ID     string `gorm:"type:uuid;primaryKey"`
	ListID string `gorm:"type:uuid"`

	Name string
}

type ListItem struct {
	ListID string `gorm:"type:uuid;primaryKey"`
	ItemID string `gorm:"type:uuid;primaryKey"`

	StatusID string `gorm:"type:uuid"`
	Status   ListStatus
}
