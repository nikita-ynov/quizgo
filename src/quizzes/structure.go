package quizzes

type Quiz struct {
	Title     string
	Questions []Question
}

type Question struct {
	Text    string
	Choices []string
	Answer  int
}
