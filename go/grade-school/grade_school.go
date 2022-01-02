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

func (grades ByLevel) Len() int {
	return len(grades)
}
func (grades ByLevel) Less(i, j int) bool {
	return grades[i].level < grades[j].level
}
func (grades ByLevel) Swap(i, j int) {
	grades[i], grades[j] = grades[j], grades[i]
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
