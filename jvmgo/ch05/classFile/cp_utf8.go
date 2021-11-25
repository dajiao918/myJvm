package classFile

type ConstantUtf8Info struct {
	str string
}

func (cui *ConstantUtf8Info) readInfo(reader *ClassReader) {
	length := uint32(reader.readUint16())
	bytes := reader.readBytes(length)
	cui.str = decode(bytes)
}

func decode(bytes []byte) string {
	str := string(bytes)
	return str
}
