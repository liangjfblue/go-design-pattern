/*
@Time : 2020/5/17 0:37
@Author : liangjiefan
*/
package wrapper_function

import (
	"io"
	"os"
	"strings"
	"testing"
)

func TestLimitReader(t *testing.T) {
	r := strings.NewReader("test1234 liangjf hello 111 11111 1111 1111")

	lr := LimitReader(r, 4)
	_, err := io.Copy(os.Stdout, lr)
	if err != nil {
		t.Fatal(err)
	}
}
