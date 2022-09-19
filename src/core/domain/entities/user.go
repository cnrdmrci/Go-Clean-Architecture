package entities

type User struct {
	ID     int    `json:"id" gorm:"column:Id"`
	Name   string `json:"name" gorm:"column:Name"`
	Gender string `json:"gender" gorm:"column:Gender"`
	Age    int    `json:"age" gorm:"column:Age"`
}
