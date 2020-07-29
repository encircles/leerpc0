package codec

import (
	"fmt"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegisterCodec(t *testing.T) {
	RegisterCodec("testCodec", nil)

	codec := GetCodec("testCodec")
	assert.Equal(t, codec, nil)
}

func TestDefaultCodec_Decode(t *testing.T) {
	c := DefaultCodec
	bytes, err := c.Encode([]byte("test"))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(bytes)
}

func TestDefaultCodec_Encode(t *testing.T) {
	s := NewUstr("Hello World!")   // 创建 Ustr 对象 s
	buf := make([]byte, s.Len()-3) // 创建缓冲区 buf

	n, err := io.ReadFull(s, buf) // 将 s 中的数据读取到 buf 中

	fmt.Printf("%s\n", buf) // HELLO WORLD!
	fmt.Println(n, err)     // 12 <nil>

	buf2 := make([]byte, 3)
	n, err = io.ReadFull(s, buf2) // 将 s 中的数据读取到 buf 中

	//n, err = io.ReadFull(s, buf) // 将
	//fmt.Printf("%s\n", buf) // HELLO WORLD!

	fmt.Printf("%s\n", buf2) // HELLO WORLD!
	fmt.Println(n, err)      // 12 <nil>
}

// 定义一个 Ustr 类型
type Ustr struct {
	s string // 数据流
	i int    // 读写位置
}

// 根据字符串创建 Ustr 对象
func NewUstr(s string) *Ustr {
	return &Ustr{s, 0}
}

// 获取未读取部分的数据长度
func (s *Ustr) Len() int {
	return len(s.s) - s.i
}

// 实现 Ustr 类型的 Read 方法
func (s *Ustr) Read(p []byte) (n int, err error) {
	for ; s.i < len(s.s) && n < len(p); s.i++ {
		c := s.s[s.i]
		// 将小写字母转换为大写字母，然后写入 p 中
		if 'a' <= c && c <= 'z' {
			p[n] = c + 'A' - 'a'
		} else {
			p[n] = c
		}
		n++
	}
	// 根据读取的字节数设置返回值
	if n == 0 {
		return n, io.EOF
	}
	return n, nil
}
