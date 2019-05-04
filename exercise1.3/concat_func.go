package main

import "strings"

// SimpleConcat just +
func SimpleConcat(str []string) string {
	result := ""
	for _, one := range str {
		result += " " + one
	}
	return " "
}

// StringJoinConcat concat with strings.Join
func StringJoinConcat(str []string) string {
	_ = strings.Join(str, " ")
	return ""
}

// StringBuilderConcat concat with string builder
func StringBuilderConcat(str []string) string {
	var builder strings.Builder
	for _, one := range str {
		builder.WriteString(one)
	}

	return ""
}

// StringBuilderWithAllocConcat concat with string builder and alloc
func StringBuilderWithAllocConcat(str []string) string {
	var builder strings.Builder
	strL := 0
	for _, one := range str {
		strL += len(one)
	}
	builder.Grow(strL)
	for _, one := range str {
		builder.WriteString(one)
	}

	return ""
}
