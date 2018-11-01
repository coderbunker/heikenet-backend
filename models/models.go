package models

type (
	User struct {
		ID       uint   `json:"id"`
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password []byte `json:"-"`
	}

	NewUser struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password []byte `json:"password"`
	}
)
