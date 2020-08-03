package codec

import (
	"fmt"
	"io"
	"strings"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

func TestRegisterCodec(t *testing.T) {
	RegisterCodec("testCodec", nil)

	codec := GetCodec("testCodec")
	assert.Equal(t, codec, nil)
}

func TestDefaultCodec_Decode(t *testing.T) {

	// var b strings.Builder
	// s := []string{"123", "456"}
	// l := len(s)
	// for i := 0; i < l; i++ {
	//	b.WriteString(s[i])
	// }
	// t.Log(b.String())

	d := DefaultCodec
	bytes, err := d.Decode([]byte("123456789012345test"))
	if err != nil {
		panic(err)
	}
	t.Log(bytes)
	p := (*string)(unsafe.Pointer(&bytes))
	t.Log(p)
	t.Log(*p)
}

func TestDefaultCodec_Encode(t *testing.T) {

	var s strings.Builder
	s.WriteString("this is test string")

	d := DefaultCodec
	en, err := d.Encode([]byte(s.String()))
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%d", en)

	de, err := d.Decode(en)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%d", de)

	t.Logf("%[1]v %[1]T %[1]d %[1]X", 0x010A)
	fmt.Println()

	// s := strings.NewReader("HELLO WORLD!")
	// buf := make([]byte, s.Len()) // 创建缓冲区 buf
	// n, err := io.ReadFull(s, buf) // 将 s 中的数据读取到 buf 中
	// t.Logf("%s\n", buf) // HELLO WORLD!
	// t.Log(n, err)     // 12 <nil>
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
