package db

import (
	"gorm.io/gorm"
)

func AutoMigrateDatabase(db *gorm.DB) {

	// var dbVersion model.DbVersion
	// err := db.First(&dbVersion).Error
	// if err != nil {
	// 	fmt.Println("error: ", err)
	// }
	// fmt.Println("db version is:", dbVersion.Version)
	// if dbVersion.Version < 1 {
	// 	err := db.AutoMigrate(&model.Player{}, &model.DailyReward{}, &model.PlayerCoins{}, &model.Session{}, &model.PowerUp{}, &model.PlayerCoins{}, model.PlayerPowerUps{}, model.ResetSession{})
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	db.Create(&model.DbVersion{
	// 		Version: 1,
	// 	})
	// 	dbVersion.Version = 1
	// }
	// if dbVersion.Version < 2 {
	// 	err := db.AutoMigrate(&model.Cart{}, &model.Payment{}, &model.CartItem{}, &model.PlayerPayment{})
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	db.Where("version=?", dbVersion.Version).Updates(&model.DbVersion{
	// 		Version: 2,
	// 	})
	// 	dbVersion.Version = 2
	// }
	// if dbVersion.Version < 3 {
	// 	err := db.AutoMigrate(&model.Avatar{}, model.PlayerAvatar{})
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	db.Where("version=?", dbVersion.Version).Updates(&model.DbVersion{
	// 		Version: 3,
	// 	})
	// 	dbVersion.Version = 3
	// }

}
