package jackanalyzer

var (
	classVarDecTypes = []string{"static, field"}
	varTypes         = []string{"int", "char", "boolean", "void"}
)

func valueInTable(value string, table []string) bool {
	for _, row := range table {
		if value == row {
			return true
		}
	}
	return false
}
