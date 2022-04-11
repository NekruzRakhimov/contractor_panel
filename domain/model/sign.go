package model

import "golang.org/x/crypto/bcrypt"

type Credentials struct {
	Id           int64
	UserLogin    string
	UserPassword string
}

func (c Credentials) ReadModel(reader DbModelReader) (interface{}, error) {
	tmp := Credentials{}
	err := reader.Scan(&tmp.Id, &tmp.UserLogin, &tmp.UserPassword)
	if err != nil {
		return nil, err
	}

	return &tmp, nil
}

func (c Credentials) Generate() (string, error) {
	saltedBytes := []byte(c.UserPassword)
	hashedBytes, err := bcrypt.GenerateFromPassword(saltedBytes, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	hash := string(hashedBytes[:])
	return hash, nil
}

func (c *Credentials) Compare(s string) error {
	incoming := []byte(s)
	existing := []byte(c.UserPassword)
	return bcrypt.CompareHashAndPassword(existing, incoming)
}

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUuid   string
	RefreshUuid  string
	AtExpires    int64
	RtExpires    int64
}

type AccessDetails struct {
	AccessUuid string
	UserId     int64
}
