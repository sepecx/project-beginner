package Lib2

import (
	"fmt"
	"math"
	"net"
	_ "os"
)

func Add(X, Y float64) float64 {
	return X + Y
}

func Sub(X, Y float64) float64 {
	return X - Y
}

func Mul(X, Y float64) float64 {
	return X * Y
}

func Div(X, Y float64) float64 {
	return X / Y
}

func Sqrt(X float64) float64 {
	return math.Sqrt(X)
}

func Pow(X, Y float64) float64 {
	return math.Pow(X, Y)
}

func LocalIP() string {

	addrs, err := net.InterfaceAddrs()

	if err != nil {
		fmt.Println(err)
	}

	for _, address := range addrs {

		// check the address type and if it is not a loopback the display it
		// = GET LOCAL IP ADDRESS
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()

			}
		}
	}
	return "ERROR"
}
