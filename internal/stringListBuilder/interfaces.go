package stringListBuilder

type StringListBuilder interface {
	BuildStringList(*StringListBuildInput) []string
}
