// Copyright 2023 SphereEx Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package random

import "crypto/rand"

const (
	normal = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
)

// StringN returns a random string with length n
func StringN(n int) string {
	return randomStringWithLen(n, []byte(normal))
}

// StringCustom returns a random string with length n and customized charset s
func StringCustom(n int, s []byte) string {
	return randomStringWithLen(n, s)
}

// randomStringWithLen returns a new random string of the provided length, consisting of the provided byte slice of allowed characters(maximum 256).
func randomStringWithLen(length int, chars []byte) string {
	if length == 0 {
		return ""
	}
	clen := len(chars)
	if clen < 2 || clen > 256 {
		return ""
	}

	bs := make([]byte, length)
	for i := 0; i < length; i++ {
		_, _ = rand.Read(bs[i : i+1])
	}
	for k, v := range bs {
		bs[k] = chars[v%byte(clen)]
	}
	return string(bs)
}
