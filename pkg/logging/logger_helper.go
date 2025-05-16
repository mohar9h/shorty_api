package logging

func mapToZapParams(extra map[ExtraKey]any) []any {
	params := make([]any, 0)
	for k, v := range extra {
		params = append(params, string(k))
		params = append(params, v)
	}
	return params
}
