package classFile

type MakerAttribute struct {
}

type DeprecatedAttribute struct {
	MakerAttribute
}

type SyntheticAttribute struct {
	MakerAttribute
}

func (self *MakerAttribute) readInfo(reader *ClassReader) {}
