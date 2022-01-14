package metrconv

import (
	"fmt"
)

type Metr float64
type Millimetr int64
type Kilometr float64

func (m Metr) String() string       { return fmt.Sprintf("%g m", m) }
func (mm Millimetr) String() string { return fmt.Sprintf("%d mm", mm) }
func (km Kilometr) String() string  { return fmt.Sprintf("%g km", km) }
