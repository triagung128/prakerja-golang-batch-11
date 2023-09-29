package user

type User struct {
	Id       uint   `json:"id" gorm:"type:int(11);primaryKey;autoIncrement"`
	Name     string `json:"name" gorm:"type:varchar(255);not null"`
	Email    string `json:"email" gorm:"type:varchar(255);uniqueIndex;not null"`
	Password string `json:"password" gorm:"type:varchar(255);not null"`
}
