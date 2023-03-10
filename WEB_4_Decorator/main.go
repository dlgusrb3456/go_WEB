package main

import (
	"fmt"

	"github.com/tuckersGo/goWeb/web9/cipher"
	"github.com/tuckersGo/goWeb/web9/lzw"
)

type Component interface {
	Operator(string)
}

var sentData string
var receiveData string

type SendComponent struct{}

func (self *SendComponent) Operator(data string) {
	// Send data
	sentData = data
}

type ZipComponent struct {
	com Component
}

func (self *ZipComponent) Operator(data string) {
	zipData, err := lzw.Write([]byte(data)) //데이터 압축하기
	if err != nil {
		panic(err)
	}
	self.com.Operator(string(zipData)) //압축한 데이터로 호출 (이래서 필요하구나)
}

type EncryptComponent struct {
	key string
	com Component
}

func (self *EncryptComponent) Operator(data string) {
	// Send data

	encryptData, err := cipher.Encrypt([]byte(data), self.key) //암호화
	if err != nil {
		panic(err)
	}
	self.com.Operator(string(encryptData))
}

type DecryptComponent struct {
	key string
	com Component
}

func (self *DecryptComponent) Operator(data string) {
	// Send data

	decryptData, err := cipher.Decrypt([]byte(data), self.key) //복호화
	if err != nil {
		panic(err)
	}
	self.com.Operator(string(decryptData))
}

type UnzipComponent struct {
	com Component
}

func (self *UnzipComponent) Operator(data string) {
	unzipData, err := lzw.Read([]byte(data)) //데이터 압축하기
	if err != nil {
		panic(err)
	}
	self.com.Operator(string(unzipData)) //압축한 데이터로 호출 (이래서 필요하구나)
}

type ReadComponent struct{}

func (self *ReadComponent) Operator(data string) {
	receiveData = data
}

func main() {
	sender := &EncryptComponent{key: "asdfe",
		com: &ZipComponent{
			com: &SendComponent{}}} //닫는 중괄호는 붙여줘야함. 혹은 , 찍어주기

	sender.Operator("Hello World")
	fmt.Println(sentData)

	receiver := &UnzipComponent{
		com: &DecryptComponent{key: "asdfe",
			com: &ReadComponent{}}}
	receiver.Operator(sentData)
	fmt.Println(receiveData)

	// 코드의 수정이 매우 간편해짐.
}
