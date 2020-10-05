package user_test

import (
	"testing"

	"github.com/Yash-Handa/Trips/internal/db"
	"github.com/Yash-Handa/Trips/internal/db/user"
	"github.com/Yash-Handa/Trips/internal/gql/model"
)

func TestRegister_Login_GetUser(t *testing.T) {
	// setting up DB
	db.Connect()
	defer db.Close()

	// setting up cab repo
	ur := &user.Repo{DB: db.GetDB()}

	// register User

	u := model.RegisterInput{
		Email:           "yash@yahoo.com",
		FirstName:       "Yash",
		Password:        "ya12sh",
		ConfirmPassword: "ya12sh",
	}

	// check email verification
	u.Email = "yash@yash.com"
	_, err := ur.RegisterUser(u)
	if err == nil {
		t.Fatalf("created user with malicious email: %s", u.Email)
	}
	u.Email = "yash@yahoo.com"

	// check pass length
	u.Password = "yash"
	_, err = ur.RegisterUser(u)
	if err == nil {
		t.Fatalf("created user with short password: %s", u.Password)
	}
	u.Password = "ya12sh"

	// check pass == confirm pass
	u.ConfirmPassword = "ya43sh"
	_, err = ur.RegisterUser(u)
	if err == nil {
		t.Fatalf("created user passwords do not match: Pass %s, ConfirmPass %s", u.Password, u.ConfirmPassword)
	}
	u.ConfirmPassword = "ya12sh"

	// check First Name is not empty
	u.FirstName = ""
	_, err = ur.RegisterUser(u)
	if err == nil {
		t.Fatalf("created user with empty FirstName")
	}
	u.FirstName = "Yash"

	_, err = ur.DB.Model(&user.User{}).Where("email = ?", "yash@yahoo.com").Delete()
	if err != nil {
		panic(err)
	}

	// create the user
	_, err = ur.RegisterUser(u)
	if err != nil {
		t.Fatalf("The user could not be created: %v", err)
	}

	// checking user login
	userLoginInput := model.LoginInput{
		Email:    "yash@yahoo.com",
		Password: "ya12sh",
	}
	auth, err := ur.LoginUser(userLoginInput)
	if err != nil {
		t.Fatalf("The user could not login: %v", err)
	}

	// Get user by UserID
	want := user.User{
		Email:     "yash@yahoo.com",
		FirstName: "Yash",
	}
	userID := auth.User.ID
	got, err := ur.GetUserByID(userID)
	if err != nil {
		t.Fatalf("Cannot get user info: %v", err)
	}

	if got.Email != want.Email {
		t.Fatalf("Required %#v but got %#v", want.Email, (*got).Email)
	}

	if got.FirstName != want.FirstName {
		t.Fatalf("Required %#v but got %#v", want.Email, (*got).Email)
	}

}
