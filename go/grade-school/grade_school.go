package school

import "sort"

// Define the Grade and School types here.
type School struct {
	grades map[int]Grade
}

type Grade struct {
	level    int
	students []string
}

// ByLevel implements sort.Interface based on level field.
type ByLevel []Grade

func (l ByLevel) Len() int {
	return len(l)
}
func (l ByLevel) Less(i, j int) bool {
	return l[i].level < l[j].level
}
func (l ByLevel) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
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
		sort.Strings(grade.students)
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
	sort.Sort(ByLevel(result))
	return result
}
