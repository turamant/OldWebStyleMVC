package models

import (
	"errors"

	"askvart.com/goals/rand"
	"askvart.com/goals/hash"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"golang.org/x/crypto/bcrypt"
)

const hmacSecretKey = "secret-hmac-key" //временно константа

var (
	ErrNotFound = errors.New("models: resource not found")
	ErrInvalidID = errors.New("models: ID provided was invalid")
	ErrInvalidPassword = errors.New("models: incorrect password provided")
	userPwPepper = "this-is-salt-pepper"
)


type User struct {
	gorm.Model
	Name string
	Age uint
	Email string `gorm:"not null;unique_index"`
	Password string `gorm:"-"`
	PasswordHash string `gorm:"not null"`
	Remember string `gorm:"-"`
	RememberHash string `gorm:"not null;unique_index"`
}

type UserService struct {
	db *gorm.DB
	hmac hash.HMAC
}

func (us *UserService) Create(user *User) error {
	pwBytes := []byte(user.Password + userPwPepper)
	hashedBytes, err := bcrypt.GenerateFromPassword(pwBytes, bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.PasswordHash = string(hashedBytes)
	user.Password = ""
	if user.Remember == ""{
		token, err := rand.RememberToken()
		if err != nil{
			return err
		}
		user.Remember = token
	}
	user.RememberHash = us.hmac.Hash(user.Remember)
	return us.db.Create(user).Error
}

func (us *UserService) Update(user *User) error{
	if user.Remember != ""{
		user.RememberHash = us.hmac.Hash(user.Remember)
	}
	return us.db.Save(user).Error
}

func NewUserService(connectionInfo string) (*UserService, error) {
	db, err := gorm.Open("postgres", connectionInfo)
	if err != nil {
		return nil, err
	}
	db.LogMode(true)
	hmac := hash.NewHMAC(hmacSecretKey)
	return &UserService{
		db: db,
		hmac: hmac,
	},nil
}

func (us *UserService) Close() error{
	return us.db.Close()
}

func (us *UserService) ByID(id int) (*User, error){
	var user User
	db := us.db.Where("id=?",id)
	err := first(db, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (us *UserService) ByAge(age uint) (*User, error){
	var user User
	db := us.db.Where("age=?",age)
	err := first(db, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func(us *UserService) InAgeRange(minAge, maxAge uint) ([]User, error){
	var users []User
	db := us.db.Where("age <=? AND age >= ?", maxAge, minAge)
	err := find(db, &users)
	if err != nil{
		return nil, err
	}
	return users, err
}
func (us *UserService) ListUsers() ([]User, error){
	var users []User
	err := find(us.db, &users)
	if err != nil{
		return nil, err
	}
	return users, err

}


func (us *UserService) ByEmail(email string) (*User, error) {
	var user User
	db := us.db.Where("email = ?", email)
	err := first(db, &user)
	return &user, err
	}

// ByRemember ищет пользователя с заданным маркером запоминания
// и возвращает этого пользователя. Этот метод будет обрабатывать хешированный
// токен для нас.
// Ошибки такие же, как для ByEmail.

func (us *UserService) ByRemember(token string) (*User, error) {
	var user User
	rememberHash := us.hmac.Hash(token)
	err := first(us.db.Where("remember_hash = ?", rememberHash), &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (us *UserService) Delete(id uint) error {
	if id == 0 {
		return ErrInvalidID
	}
	user := User{Model: gorm.Model{ID: id}}
	return us.db.Delete(&user).Error
}

func (us *UserService) AutoMigrate() error {
	if err := us.db.AutoMigrate(&User{}).Error; err != nil {
		return err
	}
	return nil
}

func (us *UserService) DestructiveReset() error{
	err := us.db.DropTableIfExists(&User{}).Error
	if err != nil{
		return err
	}
	return us.AutoMigrate()
}


//Authenticate может использоваться для аутентификации пользователя с
// указанным адресом электронной почты и паролем.
// Если указанный адрес электронной почты недействителен, это вернет
// nil, ErrNotFound
// Если предоставленный пароль недействителен, это вернет
// nil, ErrInvalidPassword
// Если адрес электронной почты и пароль верны, это вернет
// пользователя, nil
// В противном случае, если будет обнаружена другая ошибка, будет возвращено
// nil, error
func (us *UserService) Authenticate(email, password string) (*User, error) {
	foundUser, err := us.ByEmail(email)
	if err != nil{
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword(
		[]byte(foundUser.PasswordHash),
		[]byte(password+userPwPepper))
	switch err {
	case nil:
		return foundUser, nil
	case bcrypt.ErrMismatchedHashAndPassword:
		return nil, ErrInvalidPassword
	default:
		return nil, err
	}	
}




//----------------------------------------------

func first(db *gorm.DB, dst interface{}) error{
	err := db.First(dst).Error
	if err == gorm.ErrRecordNotFound {
		return ErrNotFound
	}
	return err
}

func find(db *gorm.DB, dst interface{}) error{
	err := db.Find(dst).Error
	if err == gorm.ErrRecordNotFound {
		return ErrNotFound
	}
	return err
}



