package ds

import "strata/api/strata/api"

var EventBus = make(chan *api.StreamResponse, 1000)
