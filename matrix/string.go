package matrix

import (
	"fmt"
	"strings"
)

func (m Matrix) String() string {
	var sb strings.Builder
	rows := len(m)
	cols := len(m[0])
	sb.WriteString("matrix([")
	for i := 0; i < rows; i++ {
		if i > 0 {
			sb.WriteString("        [")
		} else {
			sb.WriteString("[")
		}

		for j := 0; j < cols; j++ {
			sb.WriteString(fmt.Sprintf("%v", m[i][j]))
			if j < cols-1 {
				sb.WriteString("\t")
			}
		}
		sb.WriteString("\t]")
		if i < rows-1 {
			sb.WriteString(",\n")
		}

	}
	sb.WriteString("])")
	return sb.String()
}
