package mock

import (
	"fmt"
	"math/rand"

	conf "github.com/mt1976/crt/config"
	lang "github.com/mt1976/crt/language"
)

var config = conf.Configuration

// The randomIP function generates a random IP address in IPv4 format.
func RandomIP() string {
	// Generate a random IP address in ipv4 format
	//
	// Example: 123.456.789.012
	//
	// Returns:
	// 	string: A random IP address
	//
	// Usage:
	// 	ip := randomIP()
	// 	fmt.Println(ip)
	//
	ip1 := RandomNumber(config.DefaultRandomIPMin, config.DefaultRandomIPMax)
	ip2 := RandomNumber(config.DefaultRandomIPMin, config.DefaultRandomIPMax)
	ip3 := RandomNumber(config.DefaultRandomIPMin, config.DefaultRandomIPMax)
	ip4 := RandomNumber(config.DefaultRandomIPMin, config.DefaultRandomIPMax)

	return fmt.Sprintf(lang.IPAddressConstructor.Text(), ip1, ip2, ip3, ip4)
}

// The randomMAC function generates a random MAC address in the format of 00:00:00:00:00:00.
func RandomMAC() string {
	// Generate a random MAC address in the format of 00:00:00:00:00:00
	//
	// Returns:
	// 	string: A random MAC address
	//
	// Usage:
	// 	mac := randomMAC()
	// 	fmt.Println(mac)
	//
	mac1 := fmt.Sprintf("%02x", RandomNumber(config.DefaultRandomMACMin, config.DefaultRandomMACMax))
	mac2 := fmt.Sprintf("%02x", RandomNumber(config.DefaultRandomMACMin, config.DefaultRandomMACMax))
	mac3 := fmt.Sprintf("%02x", RandomNumber(config.DefaultRandomMACMin, config.DefaultRandomMACMax))
	mac4 := fmt.Sprintf("%02x", RandomNumber(config.DefaultRandomMACMin, config.DefaultRandomMACMax))
	mac5 := fmt.Sprintf("%02x", RandomNumber(config.DefaultRandomMACMin, config.DefaultRandomMACMax))
	mac6 := fmt.Sprintf("%02x", RandomNumber(config.DefaultRandomMACMin, config.DefaultRandomMACMax))

	return fmt.Sprintf(lang.MACAddressConstructor.Text(), mac1, mac2, mac3, mac4, mac5, mac6)
}

// The randomPort function generates a random port number between 1 and 65535.
func RandomPort() int {
	// Generate a random port number between 1 and 65535
	//
	// Returns:
	// 	int: A random port number
	//
	// Usage:
	// 	port := randomPort()
	// 	fmt.Println(port)
	//
	return RandomNumber(config.DefaultRandomPortMin, config.DefaultRandomPortMax)
}

// The randomNumber function generates a random number within a given range.
func RandomNumber(min int, max int) int {
	// Generate a random number between the given range
	//
	xx := rand.Intn(max-min+1) + min

	return xx
}

func RandomFloat(min int, max int) float64 {
	// Generate a random number between the given range
	//

	minFloat := float64(min)
	maxFloat := float64(max)

	xx := minFloat + rand.Float64()*(maxFloat-minFloat)

	return xx
}
