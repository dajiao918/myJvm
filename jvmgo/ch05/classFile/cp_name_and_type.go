package classFile

type ConstantNameAndTypeInfo struct {
	nameIndex       uint16
	descriptorIndex uint16
}

func (cnti *ConstantNameAndTypeInfo) readInfo(reader *ClassReader) {
	cnti.nameIndex = reader.readUint16()
	cnti.descriptorIndex = reader.readUint16()
}
