package main

type Storage struct {
	class []Class
}

func NewStorage() *Storage {
	studentsFirstClass := []Student{
		{
			Name: "John",
			Id:   "649dbb648ab17759e1dd6546",
		},
		{
			Name: "Erik",
			Id:   "649dbb648ab17759e1dd6547",
		},
		{
			Name: "Helen",
			Id:   "649dbb648ab17759e1dd6548",
		},
	}

	studentsSecondClass := []Student{
		{
			Name: "Ann",
			Id:   "649dbb648ab17759e1dd6549",
		},
		{
			Name: "Bob",
			Id:   "649dbb648ab17759e1dd6555",
		},
		{
			Name: "Mark",
			Id:   "649dbb648ab17759e1dd6533",
		},
	}

	return &Storage{
		[]Class{
			{
				Students: studentsFirstClass,
				Teacher: Teacher{
					Password: "secret1",
					Username: "teacher1",
					ClassId:  "649dbb5cd7c795f157361bc4",
				},
				Id: "649dbb5cd7c795f157361bc4",
			},
			{
				Students: studentsSecondClass,
				Teacher: Teacher{
					Password: "secret2",
					Username: "teacher2",
					ClassId:  "649dbb5cd7c795f157361bc5",
				},
				Id: "649dbb5cd7c795f157361bc5",
			},
		},
	}
}

func (s *Storage) GetTeacherByUsername(username string) (*Teacher, bool) {
	for _, class := range s.class {
		if class.Teacher.Username == username {
			return &class.Teacher, true
		}
	}

	return nil, false
}

func (s *Storage) GetClassByStudentId(id string) (*Class, bool) {
	for _, class := range s.class {
		for _, student := range class.Students {
			if student.Id == id {
				return &class, true
			}
		}
	}

	return nil, false
}

func (s *Storage) GetStudentById(id string) (*Student, bool) {
	for _, class := range s.class {
		for _, student := range class.Students {
			if student.Id == id {
				return &student, true
			}
		}
	}

	return nil, false
}
