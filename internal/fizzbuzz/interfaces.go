package fizzbuzz

type StringListBuilder interface {
	BuildStringList(*StringListBuildInput) []string
}
