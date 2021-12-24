package classFile

type MemberInfo struct {
	constantPool    ConstantPool
	accessFlags     uint16
	nameIndex       uint16
	descriptorIndex uint16
	attributes      []AttributeInfo
}

func readMembers(reader *ClassReader, cp ConstantPool) []*MemberInfo {
	memberCount := reader.readUint16()
	members := make([]*MemberInfo, memberCount)
	for i := range members {
		members[i] = readerMember(reader, cp)
	}
	return members
}

func readerMember(reader *ClassReader, cp ConstantPool) *MemberInfo {
	return &MemberInfo{
		constantPool:    cp,
		accessFlags:     reader.readUint16(),
		nameIndex:       reader.readUint16(),
		descriptorIndex: reader.readUint16(),
		attributes:      readAttributes(reader, cp),
	}
}

func (mem *MemberInfo) Name() string {
	return mem.constantPool.getUtf8(mem.nameIndex)
}

func (mem *MemberInfo) Descriptor() string {
	return mem.constantPool.getUtf8(mem.descriptorIndex)
}

func (mem *MemberInfo) CodeAttribute() *CodeAttribute {
	for _, attr := range mem.attributes {
		switch attr.(type) {
		case *CodeAttribute:
			return attr.(*CodeAttribute)
		}
	}
	return nil
}
