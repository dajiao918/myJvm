package classFile

type AttributeInfo interface {
	readInfo(reader *ClassReader)
}

func readAttributes(reader *ClassReader, cp ConstantPool) []AttributeInfo {
	attributeLen := reader.readUint16()
	attributes := make([]AttributeInfo, attributeLen)
	for i := range attributes {
		attributes[i] = readAttribute(reader, cp)
	}
	return attributes
}

func readAttribute(reader *ClassReader, cp ConstantPool) AttributeInfo {
	//获取attribute_NameIndex
	attrNameIndex := reader.readUint16()
	//获取attribute_len
	attrLen := reader.readUint32()
	//获取attribute_name
	attrName := cp.getUtf8(attrNameIndex)
	//通过名字创建对应的属性结构体
	attribute := newAttributeInfo(attrName, attrLen, cp)
	attribute.readInfo(reader)
	//调用对应的结构体方法
	return attribute
}

func newAttributeInfo(name string, attrLen uint32, cp ConstantPool) AttributeInfo {
	switch name {
	case "Code":
		return &CodeAttribute{cp: cp}
	case "ConstantValue":
		return &ConstantValueAttribute{}
	case "Deprecated":
		return &DeprecatedAttribute{}
	case "Exceptions":
		return &ExceptionsAttribute{}
	case "LineNumberTable":
		return &LineNumberTableAttribute{}
	case "LocalVariableTable":
		return &LocalVariableTableAttribute{}
	case "SourceFile":
		return &SourceFileAttribute{cp: cp}
	case "Synthetic":
		return &SyntheticAttribute{}
	default:
		return &UnparsedAttribute{name, attrLen, nil}
	}
}
