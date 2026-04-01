package templates

import "strings"

func Section(title string, lines ...string) string {
	var builder strings.Builder
	builder.WriteString("## ")
	builder.WriteString(title)
	builder.WriteString("\n\n")
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed == "" {
			continue
		}
		builder.WriteString(trimmed)
		builder.WriteString("\n\n")
	}
	return builder.String()
}

func BulletList(items []string) string {
	if len(items) == 0 {
		return "- none"
	}
	var builder strings.Builder
	for _, item := range items {
		trimmed := strings.TrimSpace(item)
		if trimmed == "" {
			continue
		}
		builder.WriteString("- ")
		builder.WriteString(trimmed)
		builder.WriteString("\n")
	}
	return strings.TrimSpace(builder.String())
}

func Table(headers []string, rows [][]string) string {
	var builder strings.Builder
	builder.WriteString("| ")
	builder.WriteString(strings.Join(headers, " | "))
	builder.WriteString(" |\n| ")
	separator := make([]string, len(headers))
	for index := range headers {
		separator[index] = "---"
	}
	builder.WriteString(strings.Join(separator, " | "))
	builder.WriteString(" |\n")
	for _, row := range rows {
		cells := make([]string, len(headers))
		copy(cells, row)
		for index := range cells {
			cells[index] = escapeTableCell(cells[index])
		}
		builder.WriteString("| ")
		builder.WriteString(strings.Join(cells, " | "))
		builder.WriteString(" |\n")
	}
	return strings.TrimSpace(builder.String())
}

func escapeTableCell(value string) string {
	return strings.ReplaceAll(strings.TrimSpace(value), "|", "\\|")
}
