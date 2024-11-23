package models

import (
	"fmt"
	"strings"

	"gorm.io/gorm"
)

type User struct{
	gorm.Model 

	Id    int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name  string `json:"name" gorm:"not null"`
	Email string `json:"email" gorm:"unique; not null"`
	Age   int    `json:"age"`

}

// check if name and email are not empty.
func (u *User) Validate() error{
	if strings.Trim(u.Name ," ") == " " {
	   return fmt.Errorf("Invalid users name...")
	}
	if strings.Trim(u.Email," ") == " "{
		return fmt.Errorf("Invalid users emailid...")
	}

	return nil
} 

//data sanitisation and data validation.