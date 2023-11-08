package out

type PasswordEncryptor interface {
	Encrypt(password string) (string, error)
	Compare(password string, possiblePassword string) bool
}
