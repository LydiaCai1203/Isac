package db

import (
	"framework/db/cache"
	"framework/db/pg"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	pg.ProviderSet,
	cache.NewCache,
)
