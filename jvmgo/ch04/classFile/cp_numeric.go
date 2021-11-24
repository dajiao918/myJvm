package classFile

type ConstantIntegerInfo struct {
	val int32
}

type ConstantFloatInfo struct {
	val float32
}

type ConstantLongInfo struct {
	val int64
}

type ConstantDoubleInfo struct {
	val float64
}

func (cii *ConstantIntegerInfo) readInfo(reader *ClassReader) {
	temp := reader.readUint32()
	cii.val = int32(temp)
}

func (cfi *ConstantFloatInfo) readInfo(reader *ClassReader) {
	temp := reader.readUint32()
	cfi.val = float32(temp)
}

func (cli *ConstantLongInfo) readInfo(reader *ClassReader) {
	temp := reader.readUint64()
	cli.val = int64(temp)
}

func (cdi *ConstantDoubleInfo) readInfo(reader *ClassReader) {
	temp := reader.readUint64()
	cdi.val = float64(temp)
}
