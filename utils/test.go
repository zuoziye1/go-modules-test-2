package utils

import "time"

func TestStr() string {
	return time.Now().String() + "：我变了"
}
