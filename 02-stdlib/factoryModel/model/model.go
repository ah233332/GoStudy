package model

type student struct {
	Name  string
	Age   int
	score float64
}

// student是小写，内部封装，外部暴露接口即可

func NewStudent(name string, age int, score float64) *student {
	return &student{
		Name:  name,
		Age:   age,
		score: score,
	}
}

func (s *student) GetScore() float64 {
	return s.score
}
