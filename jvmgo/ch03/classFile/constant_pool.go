package classFile

type ConstantPool []ConstantInfo

func readConstantPool(reader *ClassReader) ConstantPool {
	cp_count := int(reader.readUint16())
	cp := make([]ConstantInfo, cp_count)
	//常量池索引0不起作用，从1到n-1，不包括n
	for i := 1; i < cp_count; i++ {
		cp[i] = readConstantInfo(reader, cp)
		switch cp[i].(type) {
		//ConstantLongInfo,ConstantDoubleInfo两个属性占两个位置
		case *ConstantLongInfo, *ConstantDoubleInfo:
			i++
		}
	}
	return cp
}

/*从constantPool根据传进来的index中获取一项constantInfo*/
func (cp ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	cpInfo := cp[index]
	if cpInfo != nil {
		return cpInfo
	}
	panic("Invalid constant pool index!")
}

/*获取到constantInfo后，转换为NameAndType实现类，再根据其中的nameIndex和descriptorIndex从
常量池中获取name和type*/
func (cp ConstantPool) getNameAndType(index uint16) (string, string) {
	ntInfo := cp.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	name := cp.getUtf8(ntInfo.nameIndex)
	_type := cp.getUtf8(ntInfo.descriptorIndex)
	return name, _type
}

/*根据index取出一项constantInfo，转换为constantClassInfo，根据classInfo的下标从常量池中获取信息*/
func (cp ConstantPool) getClassName(index uint16) string {
	classInfo := cp.getConstantInfo(index).(*ConstantClassInfo)
	return cp.getUtf8(classInfo.nameIndex)
}

/*根据index取出一项constantInfo，转换为ConstantUtf8Info，返回此结构体的str属性*/
func (cp ConstantPool) getUtf8(index uint16) string {
	utf8Info := cp.getConstantInfo(index).(*ConstantUtf8Info)
	return utf8Info.str
}
