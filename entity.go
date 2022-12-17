package main

// 抽象飞船、外星人、子弹
type Entity interface {
	Width() int
	Height() int
	X() float64
	Y() float64
}
