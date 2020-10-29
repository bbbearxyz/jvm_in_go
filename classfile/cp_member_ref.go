package classfile

type ConstantMemberInfo struct {
	cp					ConstantPool
	classIndex			uint16
	nameAndTypeIndex	uint16
}

func (self *ConstantMemberInfo) readInfo(reader *ClassReader) {
	self.classIndex = reader.readUint16()
	self.nameAndTypeIndex = reader.readUint16()
}

func (self *ConstantMemberInfo) ClassName() string {
	return self.cp.getClassName(self.classIndex)
}

func (self *ConstantMemberInfo) NameAndTypeDescriptor() (string, string) {
	return self.cp.getNameAndType(self.nameAndTypeIndex)
}

type ConstantFieldRefInfo struct {
	ConstantMemberInfo
}

type ConstantMethodRefInfo struct {
	ConstantMemberInfo
}

type ConstantInterfaceMethodRefInfo struct {
	ConstantMemberInfo
}


