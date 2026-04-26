package main

import "github.com/gogf/gf/v2/crypto/gmd5"

func main() {
	password := "root"
	md5Password, err := gmd5.Encrypt(password)
	if err != nil {
		panic(err)
	}
	println(md5Password)
}
