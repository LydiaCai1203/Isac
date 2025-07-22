package cronjob

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewHelloWorldJob)
