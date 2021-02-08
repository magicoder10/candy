package screen

import (
    "candy/observability"
)

type screen struct {
    name   string
    logger *observability.Logger
}

func (s screen) Init() {
    s.logger.Infof("%s screen initialized\n", s.name)
}

func (s screen) Destroy() {
    s.logger.Infof("%s screen destroyed\n", s.name)
}
