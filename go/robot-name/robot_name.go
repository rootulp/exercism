package robotname

type Robot struct {
	name string
}

func (r *Robot) Name() (name string, err error) {
	return r.name, nil
}

func (r *Robot) Reset() {
	r.name = generateName()
}

func generateName() string {
	return "AA111"
}
