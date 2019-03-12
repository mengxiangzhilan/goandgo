package Security

import (
	"crypto/md5"
	"fmt"
)

/*一个Hash结构体拥有一个 io.Writer接口，你可以通过writer方法写入被 hash 的数据.你可以
通过Size方法获取hash 值的长度，Sum方法返回 hash 值。*/
func Md5main() {
	hash := md5.New()
	bytes := []byte("hello\n")
	hash.Write(bytes)
	hashValue := hash.Sum(nil)
	hashSize := hash.Size()
	for n := 0; n < hashSize; n += 4 {
		var val uint32
		val = uint32(hashValue[n])<<24 + uint32(hashValue[n+1])<<16 +
			uint32(hashValue[n+2])<<8 + uint32(hashValue[n+3])
		fmt.Printf("%x", val)
	}
	fmt.Println()
}
