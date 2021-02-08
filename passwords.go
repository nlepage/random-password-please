package main

import "math/rand"

const (
	minPasswordLength = 8
	maxPasswordLength = 30
)

var passwords chan (string)

func generatePasswords() {
	// Create a buffer of passwords so requests don't have to wait for a password to be generated.
	passwords = make(chan string, 10)

	// Derived from https://docs.djangoproject.com/en/dev/topics/auth/#django.contrib.auth.models.UserManager.make_random_password
	alphabet := "abcdefghjkmnpqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ23456789"
	password := make([]byte, maxPasswordLength)
	for {
		for i := 0; i < len(password); i++ {
			password[i] = alphabet[rand.Int()%len(alphabet)]
		}
		passwords <- string(password)
	}
}

func getPassword() string {
	counterLock.Lock()
	defer counterLock.Unlock()
	counter++
	if counterFile != nil && counter%100 == 0 {
		go saveCounter()
	}
	return <-passwords
}
