package jackanalyzer

var (
	classVarDecTypes  = []string{"static, field"}
	varTypes          = []string{"int", "char", "boolean", "void"}
	subroutineDecType = []string{"constructor", "function", "method"}
	statementsTypes   = []string{"let", "if", "while", "do", "return"}
)

func valueInTable(value string, table []string) bool {
	for _, row := range table {
		if value == row {
			return true
		}
	}
	return false
}
