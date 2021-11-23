package classFile

//tag是常量池中每一种类型的标志，并且是类型的首位数据
const (
	ConstantClass              = 7
	ConstantFiledRef           = 9
	ConstantMethodRef          = 10
	ConstantInterfaceMethodRef = 11
	ConstantString             = 8
	ConstantInteger            = 3
	ConstantFloat              = 4
	ConstantLong               = 5
	ConstantDouble             = 6
	ConstantNameAndType        = 12
	ConstantUtf8               = 1

//	ConstantMethodHandle 	   = 15
//	ConstantMethodtype   	   = 16
//	ConstantInvokeDynamic 	   = 18
)

type ConstantInfo interface {
	readInfo(reader *ClassReader)
}

func readConstantInfo(reader *ClassReader, cp ConstantPool) ConstantInfo {
	tag := reader.readUint8()
	c := newConstantInfo(tag, cp)
	c.readInfo(reader)
	return c
}

func newConstantInfo(tag uint8, cp ConstantPool) ConstantInfo {
	switch tag {
	case ConstantInteger:
		return &ConstantIntegerInfo{}
	case ConstantDouble:
		return &ConstantDoubleInfo{}
	case ConstantClass:
		return &ConstantClassInfo{cp: cp}
	case ConstantLong:
		return &ConstantLongInfo{}
	case ConstantFloat:
		return &ConstantFloatInfo{}
	case ConstantNameAndType:
		return &ConstantNameAndTypeInfo{}
	case ConstantString:
		return &ConstantStringInfo{cp: cp}
	case ConstantUtf8:
		return &ConstantUtf8Info{}
	case ConstantFiledRef:
		return &ConstantFiledrefInfo{ConstantMemberrefInfo{cp: cp}}
	case ConstantMethodRef:
		return &ConstantMethodrefInfo{ConstantMemberrefInfo{cp: cp}}
	case ConstantInterfaceMethodRef:
		return &ConstantInterfaceMethodrefInfo{ConstantMemberrefInfo{cp: cp}}
	//case ConstantMethodtype:
	//	return &ConstantMethodTypeInfo{}
	//case ConstantMethodHandle:
	//	return &ConstantMethodHandleInfo{}
	//case ConstantInvokeDynamic:
	//	return &ConstantInvokeDynamicInfo{}
	default:
		panic("java.lang.ClassFormatError: constant pool tag!")
	}
}
