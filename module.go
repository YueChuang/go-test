package go_test

func Add(i ,j int) int{
	return i+j
}

type Module struct {
	Name string
	Size int
}

func (this *Module)SetName(s string){
	this.Name = s
}

func (this *Module)GetName() string{
	return this.Name
}

func (this *Module)SetSize(i int){
	this.Size = i
}

func (this *Module)GetSize() int{
	return Add(this.Size,this.Size)
}
