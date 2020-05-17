/*
@Time : 2020/5/17 14:38
@Author : liangjiefan
*/
package adapter_function

import (
	"net/http"
	"testing"
)

func TestAdapterFunction1(t *testing.T) {
	http.HandleFunc("/tets", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello\n"))
	})

	http.ListenAndServe(":8080", nil)
}

type H struct{}

func (h *H) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("this is version 2"))
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello\n"))
}

func TestAdapterFunction2(t *testing.T) {
	mux := http.NewServeMux()
	mux.Handle("/", &H{})
	mux.HandleFunc("/bye", hello)
	http.ListenAndServe(":8080", mux)
}
