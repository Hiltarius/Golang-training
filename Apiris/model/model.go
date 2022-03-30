package model

type Order struct {
	ID int64 `gorm:"type:int" json:"id`
	Description string `gorm:"type:varchar" json:"description`
	Ts int64 `gorm:"type:int" json:"ts`


}