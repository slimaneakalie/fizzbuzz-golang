package fizzhttpMocks

import "github.com/slimaneakalie/fizzbuzz-golang/test/mocks/common"

type TestingEngineFactory struct {
	common.MockElement
	engine *TestingEngine
}

type TestingEngine struct {
	common.MockElement
}

type TestingRouterGroup struct {
	common.MockElement
}
