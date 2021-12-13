package numberutils

import (
	"fmt"
	"strconv"
)

func FormatInt64(value int64) string {
	return strconv.FormatInt(value, 10)
}

func FormatInt(value int) string {
	return strconv.Itoa(value)
}

func Pow1024(n int) int64 {
	if n <= 0 {
		return 1
	}
	if n == 1 {
		return 1024
	}
	return Pow1024(n-1) * 1024
}

func FormatBytes(bytes int64) string {
	if bytes < Pow1024(1) {
		return FormatInt64(bytes) + "B"
	} else if bytes < Pow1024(2) {
		return fmt.Sprintf("%.2fKB", float64(bytes)/float64(Pow1024(1)))
	} else if bytes < Pow1024(3) {
		return fmt.Sprintf("%.2fMB", float64(bytes)/float64(Pow1024(2)))
	} else if bytes < Pow1024(4) {
		return fmt.Sprintf("%.2fGB", float64(bytes)/float64(Pow1024(3)))
	} else if bytes < Pow1024(5) {
		return fmt.Sprintf("%.2fTB", float64(bytes)/float64(Pow1024(4)))
	} else if bytes < Pow1024(6) {
		return fmt.Sprintf("%.2fPB", float64(bytes)/float64(Pow1024(5)))
	} else {
		return fmt.Sprintf("%.2fEB", float64(bytes)/float64(Pow1024(6)))
	}
}
