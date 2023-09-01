package testmod

import "rsc.io/quote/v3"

// Hello returns a string message
func Hello() string {
	return quote.HelloV3()
}

func Proverb() string {
	return quote.Concurrency()
}
