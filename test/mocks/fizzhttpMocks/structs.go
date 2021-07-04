package fizzhttpMocks

import (
	"github.com/slimaneakalie/fizzbuzz-golang/test/mocks/common"
)

type TestingEngineFactory struct {
	common.MockElement
	engine *TestingEngine
}

type TestingEngine struct {
	common.MockElement
	RouterGroups []*TestingRouterGroup
}

type TestingRouterGroup struct {
	common.MockElement
}
