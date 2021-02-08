package seed

import (
	"fmt"
	"frmgol/models"
	"gorm.io/gorm"
)

func userSeeder(db *gorm.DB)  {
	userTemp := [...][4]interface{}{
		{1, "eko@mail.com", "ADMIN", "password"},
		{2, "pebri@mail.com", "ENTRY", "password"},
		{3, "sulis@mail.com", "GUEST", "password"},
	}

	var user models.Users

	for _, value := range userTemp {
		user.ID = uint(value[0].(int))
		user.Email = fmt.Sprintf("%v", value[1])
		user.Role = fmt.Sprintf("%v", value[2])
		user.Password = fmt.Sprintf("%v", value[3])
		db.Create(&user)
	}
}
