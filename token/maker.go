package token

import "time"

type Maker interface {
	//Create token for user
	CreateToken(username string, duration time.Duration) (string, error)

	//Check if token is valid
	VerifyToken(token string) (*Payload, error)
}
