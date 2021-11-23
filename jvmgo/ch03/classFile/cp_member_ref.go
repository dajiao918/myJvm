package classFile

type ConstantMemberrefInfo struct {
	cp               ConstantPool
	classIndex       uint16
	nameAndTypeIndex uint16
}

func (cmri *ConstantMemberrefInfo) readInfo(reader *ClassReader) {
	cmri.classIndex = reader.readUint16()
	cmri.nameAndTypeIndex = reader.readUint16()
}

func (cmri *ConstantMemberrefInfo) ClassName() string {
	return cmri.cp.getClassName(cmri.classIndex)
}

func (cmri *ConstantMemberrefInfo) NameAndType() (string, string) {
	return cmri.cp.getNameAndType(cmri.nameAndTypeIndex)
}

type ConstantFiledrefInfo struct {
	ConstantMemberrefInfo
}

type ConstantMethodrefInfo struct {
	ConstantMemberrefInfo
}

type ConstantInterfaceMethodrefInfo struct {
	ConstantMemberrefInfo
}
