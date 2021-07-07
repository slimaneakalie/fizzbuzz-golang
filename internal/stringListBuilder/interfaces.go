package stringListBuilder

type StringListBuilder interface {
	BuildStringList(*BuildInput) []string
}
