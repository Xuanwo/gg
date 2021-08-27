package codegen

type Struct struct {
	name    string
	comment string
	fields  []*Field
	methods []*Function
}

func NewStruct(name string) *Struct {
	return &Struct{name: name}
}

func (s *Struct) WithComment(comment string) *Struct {
	s.comment = comment
	return s
}

func (s *Struct) NewField(name, ty string) *Field {
	f := NewField(name, ty)
	s.fields = append(s.fields, f)
	return f
}

type Field struct {
	comment string
	name    string
	ty      string
}

func NewField(name, ty string) *Field {
	return &Field{name: name, ty: ty}
}

func (s *Field) WithComment(comment string) *Field {
	s.comment = comment
	return s
}

func (s *Field) String() string {
	buf := pool.Get()
	defer buf.Free()

	if s.comment == "" {
		buf.AppendString(formatComment(s.comment))
	}
	buf.AppendString(s.name)
	buf.AppendString(" ")
	buf.AppendString(s.ty)
	return buf.String()
}
