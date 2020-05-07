/*
@Time : 2020/5/7 22:35
@Author : liangjiefan
*/
package simple_factory

import "testing"

func TestEncryptFactory(t *testing.T) {
	encrypt := EncryptFactory("md5")
	t.Log(encrypt.Encrypt("big dog"))

	encrypt = EncryptFactory("sha1")
	t.Log(encrypt.Encrypt("big dog"))

	encrypt = EncryptFactory("crc32")
	t.Log(encrypt.Encrypt("big dog"))
}
