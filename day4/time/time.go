package time

import (
	"time"
)

var waktu string

func init()  {
	waktu = time.Now().String()
}

func GetTime() string {
	return waktu
}