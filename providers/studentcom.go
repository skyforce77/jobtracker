package providers

type studentcom struct {
	greenhouse
}

// NewStudentDotCom returns a new provider
func NewStudentDotCom() Provider {
	return &studentcom{
		greenhouse{
			"StudentDotCom",
			"studentcom",
		},
	}
}
