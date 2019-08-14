package utils

import (
	"crypto/md5"
	cryptoRand "crypto/rand"
	"encoding/binary"
	"encoding/hex"
	mathRand "math/rand"
	"sync/atomic"
	"time"
)

var globalMathRand = mathRand.New(mathRand.NewSource(time.Now().UnixNano()))

// 读取随机的字节到 p 指向的 []byte 里面.
func ReadRandomBytes(p []byte) {
	if len(p) <= 0 {
		return
	}

	// get from crypto/rand
	if _, err := cryptoRand.Read(p); err == nil {
		return
	}

	// get from math/rand
	for len(p) > 0 {
		n := globalMathRand.Int63()

		switch len(p) {
		case 8:
			p[0] = byte(n >> 56)
			p[1] = byte(n >> 48)
			p[2] = byte(n >> 40)
			p[3] = byte(n >> 32)
			p[4] = byte(n >> 24)
			p[5] = byte(n >> 16)
			p[6] = byte(n >> 8)
			p[7] = byte(n)
			return
		case 4:
			p[0] = byte(n >> 56)
			p[1] = byte(n >> 48)
			p[2] = byte(n >> 40)
			p[3] = byte(n >> 32)
			return
		case 1:
			p[0] = byte(n >> 56)
			return
		case 2:
			p[0] = byte(n >> 56)
			p[1] = byte(n >> 48)
			return
		case 3:
			p[0] = byte(n >> 56)
			p[1] = byte(n >> 48)
			p[2] = byte(n >> 40)
			return
		case 5:
			p[0] = byte(n >> 56)
			p[1] = byte(n >> 48)
			p[2] = byte(n >> 40)
			p[3] = byte(n >> 32)
			p[4] = byte(n >> 24)
			return
		case 6:
			p[0] = byte(n >> 56)
			p[1] = byte(n >> 48)
			p[2] = byte(n >> 40)
			p[3] = byte(n >> 32)
			p[4] = byte(n >> 24)
			p[5] = byte(n >> 16)
			return
		case 7:
			p[0] = byte(n >> 56)
			p[1] = byte(n >> 48)
			p[2] = byte(n >> 40)
			p[3] = byte(n >> 32)
			p[4] = byte(n >> 24)
			p[5] = byte(n >> 16)
			p[6] = byte(n >> 8)
			return
		default: // len(p) > 8
			p[0] = byte(n >> 56)
			p[1] = byte(n >> 48)
			p[2] = byte(n >> 40)
			p[3] = byte(n >> 32)
			p[4] = byte(n >> 24)
			p[5] = byte(n >> 16)
			p[6] = byte(n >> 8)
			p[7] = byte(n)
			p = p[8:]
		}
	}
}

// 获取一个随机的 uint32 整数.
func NewRandomUint32() uint32 {
	var x [4]byte
	ReadRandomBytes(x[:])
	return binary.BigEndian.Uint32(x[:])
}

// 获取一个随机的 uint64 整数.
func NewRandomUint64() uint64 {
	var x [8]byte
	ReadRandomBytes(x[:])
	return binary.BigEndian.Uint64(x[:])
}

const randomSaltLen = 45
const sessionIdSaltLen = 39

var (
	// 不同类型的 salt 切片 underlyingSalt 不同的部分,
	// 定期更新这个 underlyingSalt 来达到更新不同的 salt 的目的;
	// NOTE: 因为 salt 没有实际意义, 所以无需 lock.
	underlyingSalt [randomSaltLen + sessionIdSaltLen]byte

	pid                 uint16   // 进程号
	realMAC             [6]byte  // 本机的某一个网卡的 MAC 地址, 如果没有则取随机数
	mac                 [6]byte  // realMAC 混淆后的结果
	macSHA1HashSum      [20]byte // mac 的 SHA1 哈希码
	randomSalt          []byte   = underlyingSalt[:randomSaltLen]
	randomClockSequence uint32   = NewRandomUint32()
)

// NewRandom 返回一个随机字节数组.
//  NOTE: 返回的是原始数组, 不是可显示字符, 可以通过 hex, url_base64 等转换为可显示字符
func NewRandom() [16]byte {
	var src [8 + 2 + randomSaltLen]byte // 8+2+45 == 55

	nowUnixNano := time.Now().UnixNano()
	src[0] = byte(nowUnixNano >> 56)
	src[1] = byte(nowUnixNano >> 48)
	src[2] = byte(nowUnixNano >> 40)
	src[3] = byte(nowUnixNano >> 32)
	src[4] = byte(nowUnixNano >> 24)
	src[5] = byte(nowUnixNano >> 16)
	src[6] = byte(nowUnixNano >> 8)
	src[7] = byte(nowUnixNano)

	seq := atomic.AddUint32(&randomClockSequence, 1)
	src[8] = byte(seq >> 8)
	src[9] = byte(seq)

	copy(src[10:], randomSalt)

	return md5.Sum(src[:])
}

// NewToken 返回一个32字节的随机数.
//  NOTE: 返回的结果经过了 hex 编码.
func NewToken() (token []byte) {
	random := NewRandom()
	token = make([]byte, hex.EncodedLen(len(random)))
	hex.Encode(token, random[:])
	return
}
