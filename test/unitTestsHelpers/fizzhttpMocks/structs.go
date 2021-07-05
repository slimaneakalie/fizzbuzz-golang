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
	RouterGroups             []*TestingRouterGroup
	runMethodOutput          error
	formatBindingErrorOutput error
}

type TestingRouterGroup struct {
	common.MockElement
}

type TestingRequestContext struct {
	common.MockElement
	shouldBindBodyWithOutput error
}
