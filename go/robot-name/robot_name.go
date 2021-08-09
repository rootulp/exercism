package robotname

type Robot struct {
	name string
}

var robotNames = map[string]bool{}

func (r *Robot) Name() (name string, err error) {
	if r.name == "" {
		r.name = generateName()
		robotNames[r.name] = true
	}
	return r.name, nil
}

func (r *Robot) Reset() {
	delete(robotNames, r.name)
	r.name = generateName()
	robotNames[r.name] = true
}

func generateName() string {
	return generateNamePrefix(2) + generateNameSuffix(3)
}
