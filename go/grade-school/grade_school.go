package school

import (
	"sort"
)

type School struct {
	grades   map[int][]string
	students map[string]bool
}

type Grade struct {
	level    int
	students []string
}

func New() *School {
	return &School{
		grades:   make(map[int][]string),
		students: make(map[string]bool),
	}
}

func (s *School) Add(student string, grade int) {
	if _, exists := s.students[student]; exists {
		return
	}
	s.students[student] = true
	s.grades[grade] = append(s.grades[grade], student)
}

func (s *School) Grade(level int) []string {
	if students, exists := s.grades[level]; exists {
		return students
	}
	return nil
}

func (s *School) Enrollment() []Grade {
	grades := make([]Grade, 0, len(s.grades))
	for level, students := range s.grades {
		sort.Slice(students, func(i, j int) bool {
			return students[i] < students[j]
		})
		grades = append(grades, Grade{level, students})
	}
	sort.Slice(grades, func(i, j int) bool {
		return grades[i].level < grades[j].level
	})
	return grades
}
