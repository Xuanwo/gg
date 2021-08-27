package codegen

import "github.com/Xuanwo/go-bufferpool"

var pool = bufferpool.New(1024)

type File struct {
	comment     string
	packageName string
	imports     *ImportGroup
	functions   []*Function
}

func NewFile(packageName string) *File {
	return &File{
		packageName: packageName,
		imports:     NewImportGroup(),
	}
}

func (f *File) WithComment(comment string) *File {
	f.comment = comment
	return f
}

func (f *File) AddImports(ims ...*Import) *File {
	f.imports.Add(ims...)
	return f
}

func (f *File) AddFunction(fn *Function) *File {
	f.functions = append(f.functions, fn)
	return f
}

func (f *File) String() string {
	buf := pool.Get()
	defer buf.Free()

	if len(f.comment) > 0 {
		buf.AppendString(formatComment(f.comment))
	}
	buf.AppendString("package ")
	buf.AppendString(f.packageName)
	buf.AppendString("\n")

	buf.AppendString(f.imports.String())

	for _, v := range f.functions {
		buf.AppendString(v.String())
	}

	return buf.String()
}
