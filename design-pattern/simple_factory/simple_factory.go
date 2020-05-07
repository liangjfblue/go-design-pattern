/*
@Time : 2020/5/7 22:25
@Author : liangjiefan
*/
package simple_factory

import (
	"fmt"
	"log"
)

type IEncrypt interface {
	Encrypt(interface{}) string
}

type MD5 struct {}

func NewMD5() IEncrypt {
	return &MD5{}
}

func (m *MD5)Encrypt(v interface{}) string {
	return "MD5: "+fmt.Sprint(v)
}

type SHA1 struct {}

func NewSHA1() IEncrypt {
	return &SHA1{}
}

func (s *SHA1)Encrypt(v interface{}) string {
	return "SHA1: "+fmt.Sprint(v)
}

type CRC32 struct {}

func NewCRC32() IEncrypt {
	return &CRC32{}
}

func (c *CRC32)Encrypt(v interface{}) string {
	return "CRC32: "+fmt.Sprint(v)
}

func EncryptFactory(encrypt string) IEncrypt {
	switch encrypt {
	case "md5":
		return NewMD5()
	case "sha1":
		return NewSHA1()
	case "crc32":
		return NewCRC32()
	default:
		log.Fatal("not support encrypt")
	}
	return nil
}


