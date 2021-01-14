package context

// Params interface used to extract path params.
type Params interface {
	ByName(string) string
}
