/**
 *2848869
 */

package cmd

type User struct {
	UserName string
	Password string
}

func newUser(name string, password string) User {
	return User{name, password}
}
