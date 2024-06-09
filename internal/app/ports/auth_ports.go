package ports

type IAuth interface {
	CreateToken(username string) (string, error)
	VerifyToken(tokenString string) error
}
