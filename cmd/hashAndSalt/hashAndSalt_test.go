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

/*
Test the GeneratePasswordAndSaltHash function
- The password should not be in the sha512-hash for the given password => Rainbow-Table
- and also not in plain text

Test the generateSalt function
- the salt should be different every time
- a salt should be generated that is not nil

Test the match function
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

func TestGenerateSalt(t *testing.T) {
	salt1 := getSalt([]byte("test"))
	salt2 := getSalt([]byte("test"))

	assert.NotEqual(t, salt1, salt2, "Salt is equal!")
}
