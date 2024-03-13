package main

type Entry struct {
	username, password string
}

type PasswordFormat interface {
	read(fileName string) Entry
	toString(e Entry) string
}
