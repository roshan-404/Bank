package test

import (
	"simpleBank/util"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
	password := util.RandomString(6)

	hashedPassword1, err := util.HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword1)

	err = util.CheckPassword(password, hashedPassword1)
	require.NoError(t, err)

	wrongPassword := util.RandomString(6)
    err = util.CheckPassword(wrongPassword, hashedPassword1)
    require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())

	hashedPassword2, err := util.HashPassword(password)
    require.NoError(t, err)
    require.NotEmpty(t, hashedPassword2)
    require.NotEqual(t, hashedPassword1, hashedPassword2)
}