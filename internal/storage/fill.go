package storage

func FillStudent(student *Student,
	programStorage ProgramStorage, facultyStorage FacultyStorage) {
	if student == nil {
		return
	}

	if student.ProgramID != 0 {
		if program, ok := programStorage.Program(student.ProgramID); ok {
			student.ProgramShort = program.Short
			student.ProgramLong = program.Long
		}
	} else if student.ProgramShort != "" {
		if program, ok := programStorage.ProgramByShort(student.ProgramShort); ok {
			student.ProgramID = program.ID
			student.ProgramLong = program.Long
		}
	}

	if student.FacultyID != 0 {
		if faculty, ok := facultyStorage.Faculty(student.FacultyID); ok {
			student.FacultyShort = faculty.Short
			student.FacultyLong = faculty.Long
		}
	} else if student.FacultyShort != "" {
		if faculty, ok := facultyStorage.FacultyByShort(student.FacultyShort); ok {
			student.FacultyID = faculty.ID
			student.FacultyLong = faculty.Long
		}
	}
}

func FillSubject(subject *Subject,
	programStorage ProgramStorage, facultyStorage FacultyStorage, typeSubStorage TypeSubStorage) {
	if subject == nil {
		return
	}

	if subject.ProgramID != 0 {
		if program, ok := programStorage.Program(subject.ProgramID); ok {
			subject.ProgramShort = program.Short
			subject.ProgramLong = program.Long
		}
	} else if subject.ProgramShort != "" {
		if program, ok := programStorage.ProgramByShort(subject.ProgramShort); ok {
			subject.ProgramID = program.ID
			subject.ProgramLong = program.Long
		}
	}

	if subject.FacultyID != 0 {
		if faculty, ok := facultyStorage.Faculty(subject.FacultyID); ok {
			subject.FacultyShort = faculty.Short
			subject.FacultyLong = faculty.Long
		}
	} else if subject.FacultyShort != "" {
		if faculty, ok := facultyStorage.FacultyByShort(subject.FacultyShort); ok {
			subject.FacultyID = faculty.ID
			subject.FacultyLong = faculty.Long
		}
	}

	if subject.TypeSubID != 0 {
		if typeSub, ok := typeSubStorage.TypeSub(subject.TypeSubID); ok {
			subject.TypeSubShort = typeSub.Short
			subject.TypeSubLong = typeSub.Long
		}
	} else if subject.TypeSubShort != "" {
		if typeSub, ok := typeSubStorage.TypeSubByShort(subject.TypeSubShort); ok {
			subject.TypeSubID = typeSub.ID
			subject.TypeSubLong = typeSub.Long
		}
	}
}
