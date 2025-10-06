package schema

import (
	"errors"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var MAIN_DATABASE = "fatcat.db"

func ConnectAndMigrate() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(MAIN_DATABASE), &gorm.Config{})
	if err != nil {
		error := errors.New("GetConnection: unable to open " + MAIN_DATABASE + err.Error())
		log.Fatal(error)
	}

	// TODO check production migration
	aErr := db.AutoMigrate(
		&User{}, &Cat{}, &WeightJounal{}, &ExerciseJounal{}, &MealJounal{}, &Membership{})

	if aErr != nil {
		error := errors.New("GetConnection: unable to migrate schema for" + MAIN_DATABASE + aErr.Error())
		log.Fatal(error)
	}
	return db
}

type User struct {
	gorm.Model
	Membership Membership `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Cat        []Cat      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Username   string     `gorm:"type:text;not null;default:guest" json:"username"`
	Email      string     `gorm:"type:text;unique;uniqueindex" json:"email"`
	Country    string     `gorm:"type:text;default:unknown" json:"country"`
}

type Cat struct {
	gorm.Model
	UserID         int
	WeightJounal   []WeightJounal   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ExerciseJounal []ExerciseJounal `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	MealJounal     []MealJounal     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Name           string           `gorm:"type:text;not null;default:loaf-and-paw" json:"name"`
	Age            float64          `gorm:"type:real;not null;default:1" json:"age"`
	Breed          string           `gorm:"type:text;not null;default:mixed" json:"breed"`
	Weight         float64          `gorm:"type:real;not null;default:1" json:"weight"`
	Picture        string           `gorm:"type:text" json:"picture"`
}

type WeightJounal struct {
	gorm.Model
	CatID       int
	Amount      float64 `gorm:"type:real;not null;default:1" json:"amount"`
	Description string  `gorm:"type:text;not null;default:okay-condition" json:"description"`
}

type ExerciseJounal struct {
	gorm.Model
	CatID       int
	Hour        float64 `gorm:"type:real;not null;default:0.5" json:"hour"`
	Location    string  `gorm:"type:text;not null;default:house" json:"location"`
	Description string  `gorm:"type:text;not null;default:cat wheel" json:"description"`
}

type MealJounal struct {
	gorm.Model
	CatID       int
	Hour        float64 `gorm:"type:real;not null;default:0.5" json:"hour"`
	Menu        string  `gorm:"type:text;not null;default:cat food, water, churu" json:"menu"`
	Description string  `gorm:"type:text;not null;default:breakfast" json:"description"`
}

type Membership struct {
	gorm.Model
	UserID   int
	Premium  bool   `gorm:"type:text;not null;default:false" json:"premium"`
	Price    int    `gorm:"type:integer;not null;default:10" json:"price"`
	Currency string `gorm:"type:text;not null;default:USD" json:"currency"`
}
