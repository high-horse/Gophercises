package main

import "bytes"

func Normalize(phone string) string {
	var buf  bytes.Buffer
	for _, ch := range phone  {
		if ch >= '0' && ch <= '9' {
			buf.WriteRune(ch)
		} 
	}
	return buf.String()
}

func main() {

}