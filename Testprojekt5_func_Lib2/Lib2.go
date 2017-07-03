package Lib2

import (
	_ "fmt"
	"math"
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
