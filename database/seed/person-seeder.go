package seed

import (
	"fmt"
	"frmgol/models"
	"gorm.io/gorm"
	"math/rand"
)

func personSeeder(db *gorm.DB) {
	userTemp := [...][8]interface{}{
		{1, 6, "C001", "Andrew Hasim", "Andrew", "Hasim", "19-12-1980", "Singapore"},
		{2, 6, "C002", "Tania Ahmad", "Tania", "Ahmad", "27-04-1990", "Jakarta"},
		{3, 6, "C003", "Babe Defera", "Babe", "Defera", "08-01-1994", "Jakarta"},
		{4, 6, "C001", "Ahmad Maliki", "Babe", "Defera", "08-01-1994", "Jakarta"},
		{5, 6, "C001", "Eko Pebri", "Babe", "Defera", "08-01-1994", "Jakarta"},
		{6, 6, "C001", "Pebri Eko", "Babe", "Defera", "08-01-1994", "Jakarta"},
		{7, 6, "C001", "Babe 1", "Babe", "Defera", "08-01-1994", "Jakarta"},
		{8, 6, "C001", "Babe 2", "Babe", "Defera", "08-01-1994", "Jakarta"},
		{9, 6, "C001", "Babe 3", "Babe", "Defera", "08-01-1994", "Jakarta"},
		{10, 6, "C001", "Babe 4", "Babe", "Defera", "08-01-1994", "Jakarta"},
		{11, 6, "C001", "Babe 5", "Babe", "Defera", "08-01-1994", "Jakarta"},
		{12, 6, "C001", "Babe 6", "Babe", "Defera", "08-01-1994", "Jakarta"},
		{13, 6, "C001", "Babe 7", "Babe", "Defera", "08-01-1994", "Jakarta"},
		{14, 6, "C001", "Babe 8", "Babe", "Defera", "08-01-1994", "Jakarta"},
		{15, 6, "C001", "Babe 9", "Babe", "Defera", "08-01-1994", "Jakarta"},
		{16, 6, "C001", "Babe Defer1a", "Babe", "Defera", "08-01-1994", "Jakarta"},
		{17, 6, "C0017", "Babe Defe3ra", "Babe", "Defera", "08-01-1994", "Jakarta"},
		{18, 6, "C0018", "Babe Def44era", "Babe", "Defera", "08-01-1994", "Jakarta"},
		{19, 6, "C0019", "Babe Defe1ra", "Babe", "Defera", "08-01-1994", "Jakarta"},
		{20, 6, "C0020", "Babe De1fe1ra", "Babe", "Defera", "08-01-1994", "Jakarta"},
	}

	var person models.Persons

	gender := [...]string{"M", "F"}
	zone := [...]string{"WIB", "WITA", "WIT"}

	for _, value := range userTemp {
		person.ID = uint(value[0].(int))
		person.SubDistrictID = uint(value[1].(int))
		person.Nip = fmt.Sprintf("%v", value[2])
		person.FullName = fmt.Sprintf("%v", value[3])
		person.FirstName = fmt.Sprintf("%v", value[4])
		person.LastName = fmt.Sprintf("%v", value[5])
		person.BirthDate = fmt.Sprintf("%v", value[6])
		person.BirthPlace = fmt.Sprintf("%v", value[7])
		person.Gender = gender[rand.Intn(1)]
		person.ZoneLocation = zone[rand.Intn(2)]
		db.Create(&person)
	}
}
