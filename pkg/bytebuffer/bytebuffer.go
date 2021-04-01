package bytebuffer

type ByteBuf interface {
	//Read 从字节缓存中读取数据到指定字节数组
	Read([]byte) (int, error)
	//Write 把指定的字节数组数据读到字节缓存中
	Write([]byte) (int, error)
}

type DefaultByteBuf struct {
	//data 字节数组的元数据
	data     []byte
	//position 当前读取的位置下标
	position int
	//limit 结束标记的位置下标
	limit    int
	//capacity 当前元数据数组的
	capacity int
}


