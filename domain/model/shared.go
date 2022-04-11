package model

type Meta map[string]interface{}
type NamedArguments map[string]interface{}

type DbModelReader interface {
	Scan(dest ...interface{}) error
}

type DbModelProvider interface {
	// ReadModel создает новый экземпляр модели, заполняет его данными из reader-а.
	// Вовзращает ссылку на созданную модель (interface{} = *SampleModel)
	ReadModel(reader DbModelReader) (interface{}, error)
}

type SimpleModelProvider struct {
	value     interface{}
	readModel *func(reader DbModelReader) (interface{}, error)
}

func NewSimpleModelProvider(readModel func(reader DbModelReader) (interface{}, error)) *SimpleModelProvider {
	return &SimpleModelProvider{readModel: &readModel}
}

func (s SimpleModelProvider) ReadModel(reader DbModelReader) (interface{}, error) {
	value, err := (*s.readModel)(reader)
	if err != nil {
		return nil, err
	}

	return &SimpleModelProvider{value: value, readModel: s.readModel}, nil
}

func (s SimpleModelProvider) Value() interface{} {
	return s.value
}

type CodeNameStruct struct {
	Code string
	Name string
}
