package metrconv

// KmToMM convert distance from metr to millimetrs
func KmToM(km Kilometr) Metr { return Metr(km * 1000) }

// KmToMM convert distance from kilometrs to millimetrs
func KmToMM(km Kilometr) Millimetr { return Millimetr(km * 1000000) }
