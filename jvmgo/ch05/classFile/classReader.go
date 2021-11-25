package classFile

import "encoding/binary"

//ClassReader 用于存储class文件的全部二进制数据
type ClassReader struct {
	data []byte
}

/*在classReader的data中读取1个字节，并后移*/
func (cr *ClassReader) readUint8() uint8 {
	val := cr.data[0]     //读取1字节数据
	cr.data = cr.data[1:] //后移1字节
	return val
}

/*在classReader的data中读取2个字节，并后移*/
func (cr *ClassReader) readUint16() uint16 {
	//字节码是大端字节序
	val := binary.BigEndian.Uint16(cr.data)
	cr.data = cr.data[2:]
	return val
}

/*在classReader的data中读取4个字节，并后移*/
func (cr *ClassReader) readUint32() uint32 {
	//字节码是大端字节序
	val := binary.BigEndian.Uint32(cr.data)
	cr.data = cr.data[4:]
	return val
}

/*在classReader的data中读取8个字节，并后移*/
func (cr *ClassReader) readUint64() uint64 {
	//字节码是大端字节序
	val := binary.BigEndian.Uint64(cr.data)
	cr.data = cr.data[8:]
	return val
}

/*读取classReader的data中的uint16数组，数组长度由cr的data开头两字节给出*/
func (cr *ClassReader) readUint16Table() []uint16 {
	len := cr.readUint16()
	u16table := make([]uint16, len)
	for i := range u16table {
		u16table[i] = cr.readUint16()
	}
	return u16table
}

/*在classReader的data中读取长度为n字节的数据*/
func (cr *ClassReader) readBytes(n uint32) []byte {
	bytes := cr.data[:n] //读取0-n-1
	cr.data = cr.data[n:]
	return bytes
}
