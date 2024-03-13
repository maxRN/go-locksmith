package main

type OnePassword struct{}

func (bw *OnePassword) read(fileName string) Entry {
	return Entry{username: "onePassword", password: "secret"}
}

func (bw *OnePassword) toString(e Entry) string {
	return "OnePassword format"
}
