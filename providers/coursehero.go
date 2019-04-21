package providers

type coursehero struct {
	greenhouse
}

// NewCourseHero returns a new provider
func NewCourseHero() Provider {
	return &coursehero{
		greenhouse{
			"CourseHero",
			"coursehero",
		},
	}
}
