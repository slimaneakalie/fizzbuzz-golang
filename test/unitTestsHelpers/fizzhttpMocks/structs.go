package fizzhttpMocks

import (
	"github.com/slimaneakalie/fizzbuzz-golang/test/unitTestsHelpers/common"
)

type TestingEngineFactory struct {
	common.MockElement
	engine *TestingEngine
}

type TestingEngine struct {
	common.MockElement
	RouterGroups    []*TestingRouterGroup
	RunMethodOutput error
}

type TestingRouterGroup struct {
	common.MockElement
}
