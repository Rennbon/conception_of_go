package myInterface

type Distincter interface {
	Distinct()
}

type Person struct {
	Name   string
	Age    int
	PCount int
}

//实现了io.Writer结构
func (o *Person) Write(p []byte) (int, error) {
	o.PCount++
	o.Name = string(p) + "-" + o.Name
	return len(p), nil
}
func (*Person) Distinct() {
	return
}
