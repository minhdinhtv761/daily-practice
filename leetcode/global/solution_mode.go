package global

const (
	SolutionModeDefault   = "default"
	SolutionModeIterative = "iterative"
	SolutionModeRecursive = "recursive"
)

var solutionModeValidator = map[string]struct{}{
	SolutionModeDefault:   {},
	SolutionModeIterative: {},
	SolutionModeRecursive: {},
}

var solutionMode = SolutionModeDefault

func GetSolutionMode() string {
	return solutionMode
}

func SetSolutionMode(mode string) {
	if _, ok := solutionModeValidator[mode]; !ok {
		solutionMode = SolutionModeDefault
		return
	}
	solutionMode = mode
}
