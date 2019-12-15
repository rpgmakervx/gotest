package user

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"strconv"
	"time"
)

type User struct {
	password string
	username string
}

func (user User) GetPassword() string {
	return user.password
}

func (user *User) SetPassword(password string) {
	user.password = "ssh-" + password
}

func (user *User) Init() {
	rand.Seed(time.Now().Unix())
	user.password = hash(strconv.Itoa(rand.Int()))
	user.username = "Mr init"
}

func hash(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
