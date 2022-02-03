package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	//"gorm.io/gorm"
	//"github.com/jinzhu/gorm"
	//_ "github.com/jinzhu/gorm/dialects/postgres"
)

type App struct {
	db *gorm.DB
	r  *mux.Router
}

type User struct {
	Email    string `gorm:"primaryKey" json:"email"`
	Password string `json:"password"`
}

type Guide struct {
	Email    string `gorm:"primaryKey" json:"email"`
	Password string `json:"password"`
}

type Register struct {
	gorm.Model
	Email    string `gorm:"primaryKey" json:"email"`
	Password string `json:"password" binding:"required,min=8,max=20"`
	Role     string `json:"role"`
}

type Profile1 struct {
	Email string `gorm:"primaryKey" json:"email`
	Name  string `json:"name"`
	About string `json:"about"`
	Age   uint   `json:"age"`
	Fav1  string `json:"fav1"`
	Fav2  string `json:"fav2"`
	Fav3  string `json:"fav3"`
}

type Profile2 struct {
	Email    string `gorm:"primaryKey" json:"email`
	Name     string `json:"name"`
	About    string `json:"about"`
	Age      uint   `json:"age"`
	Address  string `json:"address"`
	Location string `json:"location"`
	Vehicle  string `json:"vehicle"`
}

type Comment struct {
	Comment  string   `json:"comment"`
	Email    string   `json:"email"`
	Location string   `json:"location"`
	User     User     `gorm:"foreignKey:email" json:"user_obj"`
	Profile1 Profile1 `gorm:"foreignKey:email" json:profile_obj"`
}

func setup(db *gorm.DB) {
	db.AutoMigrate(&User{}, &Guide{}, &Register{}, &Profile1{}, &Profile2{}, &Comment{})
	seed(db)
}

func seed(db *gorm.DB) {
	reg := []Register{
		{Email: "saduvishesha@gmail.com", Password: "vish", Role: "Tourist"},
		{Email: "visheshasadu@gmail.com", Password: "sadu", Role: "Guide"},
		{Email: "vish@gmail.com", Password: "vish", Role: "Guide"},
		{Email: "harshitha@gmail.com", Password: "harsh", Role: "Tourist"},
		{Email: "megh@gmail.com", Password: "megh", Role: "Tourist"},
		{Email: "arijit@gmail.com", Password: "arijit", Role: "Tourist"},
	}
	for _, c := range reg {
		db.Create(&c)
	}

	users := []User{
		{Email: "saduvishesha@gmail.com", Password: "vish"},
		{Email: "harshitha@gmail.com", Password: "harsh"},
		{Email: "megh@gmail.com", Password: "megh"},
		{Email: "arijit@gmail.com", Password: "arijit"},
	}
	for _, u := range users {
		db.Create(&u)
	}

	guides := []Guide{
		{Email: "visheshasadu@gmail.com", Password: "sadu"},
		{Email: "vish@gmail.com", Password: "vish"},
	}
	for _, gd := range guides {
		db.Create(&gd)
	}

	profiles := []Profile1{
		{Email: "saduvishesha@gmail.com", Name: "vishesha", About: "travel freak", Age: 23, Fav1: "Gainesville", Fav2: "Orlando", Fav3: "tampa"}, //User: sv, Register: sv_reg
		{Email: "megh@gmail.com", Name: "meghamala", About: "loves travel!", Age: 24, Fav1: "Miami", Fav2: "Gainesville", Fav3: "New york"},      // User: sv, Register: sv_reg
	}
	for _, m := range profiles {
		db.Create(&m)
	}

	profiles2 := []Profile2{
		{Email: "visheshasadu@gmail.com", Name: "visheshaSadu", About: "travel freak", Age: 23, Location: "Gainesville", Vehicle: "Ferrari car"}, //User: sv, Register: sv_reg
		{Email: "vish@gmail.com", Name: "vish", About: "loves travel!", Age: 24, Location: "Miami", Vehicle: "Range Rover XL car"},               // User: sv, Register: sv_reg
	}
	for _, pr2 := range profiles2 {
		db.Create(&pr2)
	}

}

func main() {

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("can't connect to database")
	}

	defer db.Close()
	db.LogMode(true)
	setup(db)

	var users []User
	db.Find(&users)
	for _, u := range users {
		fmt.Println("Email:", u.Email, "Password:", u.Password)
	}

	a := App{
		db: db,
		r:  mux.NewRouter(),
	}

	a.r.HandleFunc("/profile", a.getTouristprofile).Methods("GET")
	a.r.HandleFunc("/students/{email}", a.deleteTourist).Methods("DELETE")
	a.r.HandleFunc("/profiles", a.getGuideprofile).Methods("GET")
	//r.HandleFunc("/profiles", a.getGuideProfile).Methods("GET")
	a.r.HandleFunc("/edit-profile/{email}", a.updateUserProfile).Methods("PUT")
	//r.HandleFunc("/create-profile", a.createProfile).Methods("POST")
	a.r.HandleFunc("/register", a.registerUser).Methods("POST")
	a.r.HandleFunc("/register", a.registerGuide).Methods("POST")
	a.r.HandleFunc("/comments/{location}", a.getComments).Methods("GET")
	a.r.HandleFunc("/comments", a.getAllComments).Methods("GET")
	a.r.HandleFunc("/comments", a.postComments).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", a.r))

}

func doError(db *gorm.DB) {
	var fred User
	if err := db.Where("email = ?", "saduvishesha@gmail.com").First(&fred).Error; err != nil {
		log.Fatalf("Error when loading user: %s", err)
	}
}

func (a *App) getTouristprofile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var ab []User
	err := a.db.Find(&ab).Error
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	err = json.NewEncoder(w).Encode(ab)
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}

func (a *App) getGuideprofile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var gb []Guide
	err := a.db.Find(&gb).Error
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	err = json.NewEncoder(w).Encode(gb)
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}

func (a *App) getAllComments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var cm []Comment
	err := a.db.Find(&cm).Error
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	err = json.NewEncoder(w).Encode(cm)
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}

func (a *App) getComments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var cm Comment
	err := a.db.Find(&cm).Error
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	err = json.NewEncoder(w).Encode(cm)
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}

func (a *App) registerUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var reg Register
	err := json.NewDecoder(r.Body).Decode(&reg)
	if err != nil {
		sendErr(w, http.StatusBadRequest, err.Error())
		return
	}
	err = a.db.Save(&reg).Error
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	} else {
		w.WriteHeader(http.StatusCreated)
	}

}

func (a *App) registerGuide(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var gui Guide
	err := json.NewDecoder(r.Body).Decode(&gui)
	if err != nil {
		sendErr(w, http.StatusBadRequest, err.Error())
		return
	}
	err = a.db.Save(&gui).Error
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	} else {
		w.WriteHeader(http.StatusCreated)
	}
}

func (a *App) postComments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var cmnts Comment
	err := json.NewDecoder(r.Body).Decode(&cmnts)
	if err != nil {
		sendErr(w, http.StatusBadRequest, err.Error())
		return
	}
	err = a.db.Save(&cmnts).Error
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	} else {
		w.WriteHeader(http.StatusCreated)
	}
}

func (a *App) updateUserProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var u User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		sendErr(w, http.StatusBadRequest, err.Error())
		return
	}
	u.Email = mux.Vars(r)["email"]
	err = a.db.Save(&u).Error
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}

func (a *App) deleteTourist(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := a.db.Unscoped().Delete(User{Email: mux.Vars(r)["email"]}).Error
	if err != nil {
		sendErr(w, http.StatusInternalServerError, err.Error())
	}
}

func sendErr(w http.ResponseWriter, code int, message string) {
	resp, _ := json.Marshal(map[string]string{"error": message})
	http.Error(w, string(resp), code)
}
