// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package homeopathic

import (
	"crypto/aes"
	"crypto/rand"
	"testing"
)

func BenchmarkEncrypt(b *testing.B) {
	in := make([]byte, 1024*1024)
	key := make([]byte, 16)
	rand.Read(in)
	rand.Read(key)
	for i := 0; i < b.N; i++ {
		Encrypt(in, key)
	}
}

func BenchmarkAES(b *testing.B) {
	in := make([]byte, 1024*1024)
	key := make([]byte, 16)
	dst := make([]byte, len(in))
	rand.Read(in)
	rand.Read(key)
	cipher, _ := aes.NewCipher(key)
	for i := 0; i < b.N; i++ {
		cipher.Encrypt(dst, in)
	}
}
