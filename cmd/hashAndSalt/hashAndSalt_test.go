/*
 * 2848869
 * 8089098
 * 3861852
 */

package hashAndSalt

import (
	"crypto/sha512"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
)

/* Teste die Funktion GeneratePasswordAndSaltHash
- es soll das Passwort nicht im den sha512(password) sein
- und auch nicht im Klartext
*/

func getSalt(pwd []byte) []byte {
	return GenerateSalt(pwd)
}

func TestPasswordIsInPlainText(t *testing.T) {
	pwd := []byte("test")
	pwdHash := GeneratePasswordAndSaltHash(getSalt(pwd), pwd)

	assert.NotEqual(t, pwd, pwdHash, "Password is in plain text")
}

func TestPasswordIsInSHA512Hash(t *testing.T) {
	pwd := "test"
	pwdHash := GeneratePasswordAndSaltHash(getSalt([]byte(pwd)), []byte(pwd))

	hash := sha512.New()
	_, err := io.WriteString(hash, pwd)

	assert.NoError(t, err)
	assert.NotEqual(t, hash.Sum(nil), pwdHash, "Password is in SHA512 Hash")
}

func TestMatchWithRightPwd(t *testing.T) {
	userPwd := "test"
	salt := getSalt([]byte(userPwd))
	hashedPassword := GeneratePasswordAndSaltHash(salt, []byte(userPwd))

	assert.Equal(t, true, Match([]byte(userPwd), hashedPassword, salt))
}

func TestMatchWithWrongPwd(t *testing.T) {
	userPwd := "test"
	wrongPwd := "test2"

	salt := getSalt([]byte(userPwd))
	hashedPassword := GeneratePasswordAndSaltHash(salt, []byte(userPwd))

	assert.Equal(t, false, Match([]byte(wrongPwd), hashedPassword, salt))
}

func TestGenerateSaltIsNotNil(t *testing.T) {
	assert.NotEqual(t, nil, GenerateSalt([]byte("testPassword")))
}
