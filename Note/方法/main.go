package main

func main()  {
	// 先打开文件对象
	f, _ := OpenFile("foo.dat")

	// 绑定到了 f 对象
	// func Close() error
	var Close = func Close() error {
		return (*File).Close(f)
	}

	// 绑定到了 f 对象
	// func Read(int64 offset, data []byte) int
	var Read = func Read(int64 offset, data []byte) int {
		return (*File).Read(f, offset, data)
	}

	// 文件处理
	Read(0, data)
	Close()
}

// 文件对象
type File struct {
	fd int
}

// Close 关闭文件
// todo CloseFile和ReadFile函数就成了File类型独有的方法了（而不是File对象方法）。它们也不再占用包级空间中的名字资源，
//同时File类型已经明确了它们操作对象，因此方法名字一般简化为Close和Read
func (f *File) Close() error {
	// ...
	return nil
}

// 读文件数据
func (f *File) Read(int64 offset, data []byte) int {
	// ...
	return 0
}
