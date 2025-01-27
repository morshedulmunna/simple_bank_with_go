package utils

import (
	"fmt"
	"math/rand"
	"time"
)

// RandomInt generates a random integer between min and max (inclusive)
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates a random string of length n using alphabetic characters
func RandomString(n int) string {
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// RandomOwner generates a random owner name of 6 characters
func RandomOwner() string {
	return RandomString(6)
}

// RandomMoney generates a random monetary amount between 0 and 1000
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// RandomCurrency returns a random currency code from a predefined list
func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "CAD"}
	return currencies[rand.Intn(len(currencies))]
}

// RandomEmail generates a random email address with a 6-character local part
func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}

// RandomURL generates a random URL with an 8-character domain name
func RandomURL() string {
	return fmt.Sprintf("https://%s.com", RandomString(8))
}

// RandomIPv4 generates a random IPv4 address in string format
func RandomIPv4() string {
	return fmt.Sprintf("%d.%d.%d.%d",
		rand.Intn(256), rand.Intn(256),
		rand.Intn(256), rand.Intn(256))
}

// RandomPort generates a random port number between 1024 and 65535
func RandomPort() int {
	return rand.Intn(65535-1024) + 1024
}

// RandomPhone generates a random US phone number in the format +1-XXX-XXX-XXXX
func RandomPhone() string {
	return fmt.Sprintf("+1-%d%d%d-%d%d%d-%d%d%d%d",
		rand.Intn(9)+1, rand.Intn(10), rand.Intn(10),
		rand.Intn(10), rand.Intn(10), rand.Intn(10),
		rand.Intn(10), rand.Intn(10), rand.Intn(10), rand.Intn(10))
}

// RandomDate generates a random date between January 1, 2000, and the current time
func RandomDate() time.Time {
	min := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Now().Unix()
	delta := max - min
	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0)
}

// RandomBool generates a random boolean value
func RandomBool() bool {
	return rand.Intn(2) == 1
}

// RandomFloat generates a random float64 between min and max
func RandomFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

// RandomColor generates a random color in hexadecimal format (#RRGGBB)
func RandomColor() string {
	return fmt.Sprintf("#%02x%02x%02x",
		rand.Intn(256), rand.Intn(256), rand.Intn(256))
}

// RandomUsername generates a random username in the format string_number
func RandomUsername() string {
	return fmt.Sprintf("%s_%d", RandomString(5), rand.Intn(1000))
}

// RandomPassword generates a random password between 8-16 characters
// including letters, numbers, and special characters
func RandomPassword() string {
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*"
	length := rand.Intn(8) + 8 // 8-16 characters
	b := make([]byte, length)
	for i := range b {
		b[i] = chars[rand.Intn(len(chars))]
	}
	return string(b)
}

// RandomCountry returns a random country name from a predefined list
func RandomCountry() string {
	countries := []string{"USA", "Canada", "UK", "France", "Germany", "Japan", "Australia", "Brazil", "India", "China"}
	return countries[rand.Intn(len(countries))]
}
