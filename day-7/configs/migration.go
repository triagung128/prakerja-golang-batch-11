package configs

import usermodel "day-7/models/user"

func InitMigrate() {
	DB.AutoMigrate(&usermodel.User{})
}
