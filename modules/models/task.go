package models

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model

	Title       string
	Description string
	Status      bool
	Advance     int
	UserId      uint
	CreatedAt   time.Time
    UpdatedAt   time.Time
}

//  no se puede instanciar un struct si no esta dentro de una funcion
// algo asi no se puede ahacer ts := Task{} si no esta dentro de una funcion
// func Test()  {
// 	ts := Task{}
// 	fmt.Println(ts.ID)
// }
