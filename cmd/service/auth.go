package services

import (
	"context"
	"math/big"

	"crypto/rand"
	authPb "tableLink/dto/authpb"

	"gorm.io/gorm"
)

type Auth struct {
	authPb.UnimplementedAuthServer
	Db *gorm.DB
}

func GenerateRandomKey(length int) (string, error) {
    const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
    result := make([]byte, length)

    for i := range result {
        // Use crypto/rand for cryptographically secure random numbers
        n, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
        if err != nil {
            return "", err
        }
        result[i] = charset[n.Int64()]
    }

    return string(result), nil
}

func (a *Auth) Login(ctx context.Context, req *authPb.LoginRequest) (authPb.ResponseRequestLogin,error) {
	var res authPb.ResponseRequestLogin
	var isExist bool
	query := `
	SELECT EXIST(
		select 1 
		from users
	FROM users.user_info 
	WHERE email = ? and password =? `

	err := a.Db.Raw(query, req.Email, req.Password).Scan(&isExist).Error
	if err != nil {
		return res,err
	}

	token ,err:= GenerateRandomKey(10)
	if err != nil{
		return res,err
	}
	if isExist {
		res.Status = true
		res.Message = "successfuly"
		res.Data.AccessToken = "Bearer " + token
		return res,nil
	}
	return  res,err
}
