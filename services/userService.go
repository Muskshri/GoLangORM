package services
// buissness logic is written here.

import (
	"fmt"
	"gorm_demo/models"
	"gorm.io/gorm"
)

func CreateUserService(user models.User, db *gorm.DB) (models.User, error){
    //use validation error before snding to db
	if validationError := user.Validate(); validationError != nil{
		return user, validationError
	} 
   
	result:= db.Create(&user)
	if result.Error != nil{
		return user, fmt.Errorf("User can not be created...")
	}

	return user, nil
}

func GetUsersService(db *gorm.DB)([] models.User , error){
	//db connnection
	var user []models.User
	result:= db.Find(&user)
	if result.Error != nil{
		return user, fmt.Errorf("Can not fetch users from table...")
	}
	return user, nil
}

func GetUserByIdService(id int, db *gorm.DB)(models.User, error){
    var user models.User
	result:= db.First(&user, id)
	if result.Error != nil{
		return user, fmt.Errorf("User does not exist...")
	}
	return user, nil
}

func UpdateUserService(id int, user models.User, db *gorm.DB)(models.User, error){
	//check whether user exists or not
	var updatedUser models.User
	result:= db.First(&updatedUser, id)
	if result.Error != nil{
		return updatedUser, fmt.Errorf("User does not exist...")
	}

	// updating user
	// updatedUser recieves the updated info (so passing its address) the info to be updated comes into UpdatedServices as user.
	result= db.Model(&updatedUser).Updates(user) 
	if result.Error != nil{
		return user, fmt.Errorf("User Updation failed...")
	}
	return updatedUser, nil
}

func DeleteUserService(id int, db *gorm.DB) error{
	result:= db.Delete(&models.User{}, id)
	if result.Error != nil{
		return fmt.Errorf("User Delete failed...")
	}
	if result.RowsAffected == 0{
		return fmt.Errorf("User Not found...")
	}
	return nil
} 