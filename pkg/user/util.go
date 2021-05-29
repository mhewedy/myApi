package user

import "golang.org/x/crypto/bcrypt"

func hash(password string) string {
	b, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(b)
}

func verifyHash(hashed, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password)) == nil
}

// Used to trick the user, so he don't know whether the username is invalid or the password is invalid
func verifyFakeHash() {
	_ = verifyHash("$2a$10$H5IpXKNIfTvrJSY5tr4NReLauM111sgxMNtdNH1vgMp9KJrdBsLca", "")
}
