package out

var _ PasswordEncryptor = &PasswordEncryptorMock{}

type PasswordEncryptorMock struct {
	MockedEncrypt func(password string) (string, error)
	MockedCompare func(encrypted string, original string) bool
}

func (mock *PasswordEncryptorMock) Encrypt(password string) (string, error) {
	if mock.MockedEncrypt != nil {
		return mock.MockedEncrypt(password)
	}
	return "", nil
}

func (mock *PasswordEncryptorMock) Compare(encrypted string, original string) bool {
	if mock.MockedCompare != nil {
		return mock.MockedCompare(encrypted, original)
	}
	return false
}

func NewPasswordEncryptorMock() *PasswordEncryptorMock {
	return &PasswordEncryptorMock{}
}
