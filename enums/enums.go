package enums

import "github.com/Vector-ops/battleships/types"

// ship size enum
const (
	Small  types.ShipSize = iota + 1
	Medium types.ShipSize = iota + 1
	Large  types.ShipSize = iota + 1
)

// ship direction enum
const (
	Vertical   types.Direction = "v"
	Horizontal types.Direction = "h"
)
