package school

// Define the Grade and School types here.
type School struct {
	grades map[int]Grade
}

type Grade struct {
	level    int
	students []string
}

func New() *School {
	return &School{map[int]Grade{}}
}

func (s *School) Add(student string, level int) {
	grade, ok := s.grades[level]
	if !ok {
		s.grades[level] = Grade{level, []string{student}}
	} else {
		grade.students = append(grade.students, student)
		s.grades[level] = grade
	}
}

func (s *School) Grade(level int) []string {
	grade, ok := s.grades[level]
	if !ok {
		return []string{}
	} else {
		return grade.students
	}
}

func (s *School) Enrollment() (result []Grade) {
	for _, grade := range s.grades {
		result = append(result, grade)
	}
	return result
}
