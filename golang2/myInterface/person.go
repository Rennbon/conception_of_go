package myInterface

type Distincter interface {
	Distinct() error
}

type Person struct {
	Name   string
	Age    int
	PCount int
	Speed  int
}

//实现了io.Writer结构
func (o *Person) Write(p []byte) (int, error) {
	o.PCount++
	o.Name = string(p) + "-" + o.Name
	return len(p), nil
}
