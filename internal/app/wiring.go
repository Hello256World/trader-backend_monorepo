package app

import (
	"github.com/Hello256World/trader-backend_monorepo/internal/bootstrap"
)

func (a *application) mapHandlers() {
	bootstrap.WireStrategy(a.httpEngine, a.mongoClient.Database(a.conf.MongoConfig.Database))
}
