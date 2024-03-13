package main

type BitWarden struct{}

func (bw *BitWarden) read(fileName string) Entry {
	return Entry{username: "bitwarden", password: "secret"}
}

func (bw *BitWarden) toString(e Entry) string {
	return "BitWarden format"
}
