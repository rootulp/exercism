package robotname

import "errors"

type Robot struct {
	name string
}

const maxNumberOfNames = 26 * 26 * 10 * 10 * 10

var robotNames = map[string]bool{}

func (r *Robot) Name() (name string, err error) {
	if len(robotNames) == maxNumberOfNames {
		return "", errors.New("namespace exhausted")
	}
	if r.name == "" {
		r.name = generateName()
	}
	return r.name, nil
}

func (r *Robot) Reset() {
	// delete(robotNames, r.name)
	r.name = generateName()
}

func generateName() string {
	for {
		name := generateNamePrefix(2) + generateNameSuffix(3)
		if !robotNames[name] {
			robotNames[name] = true
			return name
		}
	}
}
