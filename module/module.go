package module

//Name returns the names of all the modules
//ENUM(bootstrap,instruct,max)
type Name int

type TypeHandleFunc = func(id string, data []byte) error

type Module interface {
	Start() error
	Name() Name
	HandleData(id string, data []byte) error
}
