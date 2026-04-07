package utils

import (
	"crypto/rand"
	"encoding/hex"
	"time"
)

// NewUUIDv7 生成一个 UUIDv7 字符串（小写无连字符，32 位）
// 与 PostgreSQL 18 的 uuidv7() 函数格式一致
func NewUUIDv7() string {
	ms := uint64(time.Now().UnixMilli())

	var uuid [16]byte
	// 48-bit timestamp (big-endian)
	uuid[0] = byte(ms >> 40)
	uuid[1] = byte(ms >> 32)
	uuid[2] = byte(ms >> 24)
	uuid[3] = byte(ms >> 16)
	uuid[4] = byte(ms >> 8)
	uuid[5] = byte(ms)
	// random bytes
	_, _ = rand.Read(uuid[6:])
	// version = 0111
	uuid[6] = (uuid[6] & 0x0F) | 0x70
	// variant = 10
	uuid[8] = (uuid[8] & 0x3F) | 0x80

	return hex.EncodeToString(uuid[:])
}