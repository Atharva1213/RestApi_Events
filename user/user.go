package user

import (
	"Server_main/database"
	"Server_main/utilty"

	"errors"
)

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email" binding:"required"`
	Password string  `json:"password" binding:"required"`
}

func (user User) Save() error {
	query := `INSERT INTO user (email,password) VALUES (?,?);`
	stm, err := database.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stm.Close()
	hashedPassword, err := utilty.HashingPassword(user.Password)
	if err != nil {
		return err
	}

	_, err = stm.Exec(user.Email, hashedPassword)
	if err != nil {
		return errors.New("email is already registered")
	}

	return nil
}


func (user User) Login() (string,error) {
    query := `SELECT id,password FROM user WHERE email = ?`
    
	var temp User
    err := database.DB.QueryRow(query, user.Email).Scan(&temp.ID,&temp.Password)
    if err != nil {
        return "",errors.New("email is not registered")
    }

   // Verify the password
   ok, err := utilty.VerifiedPassword(temp.Password, user.Password)
   if err != nil {
	   return "", err
   }
   if !ok {
	   return "", errors.New("invalid credentials, try again")
   }

   // Generate token
   token, err := utilty.GenerateToken(user.Email,temp.ID)
   if err != nil {
	   return "", err
   }

   return token,nil
}


