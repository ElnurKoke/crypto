package crypto

import (
	"encoding/binary"
	"fmt"
)

func Sha1(message string) string {
	h0 := uint32(0x67452301)
	h1 := uint32(0xEFCDAB89)
	h2 := uint32(0x98BADCFE)
	h3 := uint32(0x10325476)
	h4 := uint32(0xC3D2E1F0)
	data := []byte(message)
	data = append(data, 0x80)
	for (len(data)+8)%64 != 0 {
		data = append(data, 0)
	}
	fmt.Println("data2:", data)

	length := uint64(len(message) * 8)
	data = append(data, make([]byte, 8)...)
	binary.BigEndian.PutUint64(data[len(data)-8:], length)

	for i := 0; i < len(data); i += 64 {
		chunk := data[i : i+64]
		var w [80]uint32
		for j := 0; j < 16; j++ {
			w[j] = binary.BigEndian.Uint32(chunk[j*4 : (j+1)*4])
		}
		for j := 16; j < 80; j++ {
			w[j] = leftRotate(w[j-3]^w[j-8]^w[j-14]^w[j-16], 1)
		}
		a, b, c, d, e := h0, h1, h2, h3, h4
		for j := 0; j < 80; j++ {
			var f, k uint32
			if j < 20 {
				f = (b & c) | ((^b) & d)
				k = 0x5A827999
			} else if j < 40 {
				f = b ^ c ^ d
				k = 0x6ED9EBA1
			} else if j < 60 {
				f = (b & c) | (b & d) | (c & d)
				k = 0x8F1BBCDC
			} else {
				f = b ^ c ^ d
				k = 0xCA62C1D6
			}

			temp := leftRotate(a, 5) + f + e + k + w[j]
			e = d
			d = c
			c = leftRotate(b, 30)
			b = a
			a = temp
		}

		h0 += a
		h1 += b
		h2 += c
		h3 += d
		h4 += e
	}

	digest := fmt.Sprintf("%08x%08x%08x%08x%08x", h0, h1, h2, h3, h4)
	fmt.Println("SHA-1:", digest)
	return digest
}
func leftRotate(n uint32, bits int) uint32 {
	return (n << bits) | (n >> (32 - bits))
}
