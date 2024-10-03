package internal

import (
	"encoding/json"
)

type ApiObject interface {
	ToJson() (string, error)
	FromJson(string) (ApiObject, error)
}

type ApiUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	UserId   string `json:"id"`
	ImageSrc string `json:"image"`
	// SessionId string `json:"sessionId"`
	// Username string `json:"username"`
	// ProfileImageSrc string `json:"imageUrl"`
}

func (a *ApiUser) ToJson() (string, error) {
	res, err := json.Marshal(a)
	if err != nil {
		return "", err
	}
	return string(res), nil
}
