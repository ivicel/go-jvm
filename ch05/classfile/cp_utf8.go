package classfile

import (
	"fmt"
	"unicode/utf16"
)

// MUTF-8 编码的字符串, 和标准 utf8 并不完全相同
// 一是 null 字符(U+0000)会被编码成两字节, 0xc0, 0x80
// 二是补充字符(Supplementary Characters, 代码点大于 U+FFFF 是把 utf16 代理对)
// 再编码成为 utf8, 也就是 6 个字节
//
// CONSTANT_Utf8_info {
// 	u1 tag;
// 	u2 length;
// 	u1 bytes[length];
// }
type ConstantUtf8Info struct {
	str string
}

func (c *ConstantUtf8Info) readInfo(reader *ClassReader) {
	length := reader.readUint16()
	bytes := reader.readBytes(uint32(length))
	c.str = decodeMUTF8(bytes)
}

// decodeMUTF8 解析 mutf8 字符
//
// mutf8 -> utf16 -> utf32 -> string
// see java.io.DataInputStream.readUTF(DataInput)
func decodeMUTF8(bytes []byte) string {
	utflen := len(bytes)
	chararr := make([]uint16, utflen)

	var c, char2, char3 uint16
	count := 0
	chararr_count := 0

	//					   0xxx xxxx -> 0x1 - 0x7F
	//           110x xxxx 10xx xxxx -> 0xC080 - 0xDFBF
	// 1110 xxxx 10xx xxxx 10xx xxxx ->
	// 单字节: 0001 - 0111 (1->7)
	// 双字节: 1100 - 1101 (12->13)
	// 三字节: 1110 - 1110 (14->14)

	for count < utflen {
		c = uint16(bytes[count])
		switch c >> 4 { // 右移 4 位后
		case 0, 1, 2, 3, 4, 5, 6, 7: // 单字节下只会出现 0b0001(1) 到 0b0111(7)
			// 单字节 utf8: 0xxx xxxx
			// 单字节 utf8 和 utf16 低 8 位码是一样了
			count++
			chararr[chararr_count] = c
			chararr_count++

		case 12, 13: // 双字节只会出现 0b1100(12) 和 0b1101(13)
			// 双字节: 110x xxxx 10xx xxxx
			count += 2
			if count > utflen {
				panic("malformed input: partial character at end")
			}

			char2 = uint16(bytes[count-1])
			// 双字节 utf8, 第二个字节肯定是 0b10xx 开头
			if (char2 & 0xc0) != 0x80 {
				panic(fmt.Errorf("malformed input around byte %v", count))
			}

			// 取出 utf8 中的有效位
			chararr[chararr_count] = (c & 0x1f << 6) | (char2 & 0x3f)
			chararr_count++

		case 14:	// 三字节只会出现  0b1110(14)
			// 三字节: 1110 xxxx 110x xxxx 10xx xxxx
			count += 3
			if count > utflen {
				panic("malformed input: partial character at end")
			}

			char2 = uint16(bytes[count-2])
			char3 = uint16(bytes[count-1])

			// 三字节 utf8 第二字, 第三字节肯定也是以 0b10xx 开头
			if (char2&0xc0) != 0x80 || (char3&0xc0) != 0x80 {
				panic(fmt.Errorf("malformed input around byte %v", (count - 1)))
			}

			// 取出 utf8 中的有效位
			chararr[chararr_count] = (c&0x0F)<<12 | (char2&0x3F)<<6 | (char3&0x3F)<<0
			chararr_count++

		default:
			// 10xx xxxx, 1111 xxxx
			panic(fmt.Errorf("malformed input around byte %v", count))
		}
	}

	chararr = chararr[0:chararr_count]
	runes := utf16.Decode(chararr)
	return string(runes)
}
