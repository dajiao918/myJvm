package classFile

type ClassFile struct {
	magic        uint32          //魔数CAFEBABY
	minorVersion uint16          //字节码的次要版本
	majorVersion uint16          //字节码的主要版本
	constantPool ConstantPool    //常量池结构体
	accessFlag   uint16          //类的权限修饰符，是否是final类，enum，接口等
	thisClass    uint16          //当前类名常量池索引
	superClass   uint16          //父类名常量池索引
	interfaces   []uint16        //实现的接口数量
	fields       []*MemberInfo   //实例信息结构体
	methods      []*MemberInfo   //方法信息结构体
	attributes   []AttributeInfo //其他属性结构体
}

func Parse(classData []byte) (cf *ClassFile, err error) {
	defer func() {
		r := recover()
		if r != nil {
			err, ok := r.(error)
			if !ok {
				panic(err)
			}
		}
	}()
	classReader := &ClassReader{classData}
	cf = &ClassFile{}
	cf.read(classReader)
	return
}

func (cf *ClassFile) read(reader *ClassReader) {
	cf.readAndCheckMagic(reader)                            //读取魔数
	cf.readAndCheckVersion(reader)                          //读取版本
	cf.constantPool = readConstantPool(reader)              //读取常量池
	cf.accessFlag = reader.readUint16()                     //读取访问标识符
	cf.thisClass = reader.readUint16()                      //读取当前类名
	cf.superClass = reader.readUint16()                     //读取父类类名
	cf.interfaces = reader.readUint16Table()                //读取接口数量
	cf.fields = readMembers(reader, cf.constantPool)        //读取字段信息
	cf.methods = readMembers(reader, cf.constantPool)       //读取方法信息
	cf.attributes = readAttributes(reader, cf.constantPool) //读取其他属性信息
}

func (cf *ClassFile) MajorVersion() uint16 {
	return cf.majorVersion
}

func (cf *ClassFile) MinorVersion() uint16 {
	return cf.minorVersion
}

func (cf *ClassFile) ConstantPool() ConstantPool {
	return cf.constantPool
}

func (cf *ClassFile) AccessFlags() uint16 {
	return cf.accessFlag
}

func (cf *ClassFile) Fields() []*MemberInfo {
	return cf.fields
}

func (cf *ClassFile) Methods() []*MemberInfo {
	return cf.methods
}

func (cf *ClassFile) ClassName() string {
	return cf.constantPool.getClassName(cf.thisClass)
}

func (cf *ClassFile) SuperClassName() string {
	if cf.superClass > 0 {
		return cf.constantPool.getClassName(cf.superClass)
	}
	return ""
}

func (cf *ClassFile) InterfacesName() []string {
	strs := make([]string, len(cf.interfaces))
	for i, cpIndex := range cf.interfaces {
		strs[i] = cf.constantPool.getClassName(cpIndex)
	}
	return strs
}

func (cf *ClassFile) readAndCheckMagic(reader *ClassReader) {
	cf.magic = reader.readUint32()
	if cf.magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError: magic!")
	}
}

func (cf *ClassFile) readAndCheckVersion(reader *ClassReader) {
	cf.minorVersion = reader.readUint16()
	cf.majorVersion = reader.readUint16()
	//主版本45支持此版本为0-65535
	//而主版本在46及46之上，次版本只能为0
	switch cf.majorVersion {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if cf.minorVersion == 0 {
			return
		}
	}
	panic("java.lang.UnsupportedClassVersionError!")
}
