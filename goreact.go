package main // Basic goReact program

import "fmt" //"/Users/jalspach/Documents/GitHub/simparts/parts/parts.go"

//"github.com/jalspach/simparts.git/parts.go"
//"/simparts/parts.go"
//"github.com/jalspach/simparts/parts.go"

//start by defining two parts and giving them qualities. Pass data beteen them.
func main() {
	PCP := BasicDynamic{
		BasicStats: BasicStats{0, 0, 0, 0, 100, 100, 100, 100, 100, 0, "Green"},
		RPM:        0,
		KWH:        0,
		ReqUse:     0,
		CurUse:     0,
		LimKWH:     100,
		LimRPM:     500,
	}
	fmt.Println("Temp", PCP.Bar)
	fmt.Println("AlarmColor", PCP.AlarmColor)
	fmt.Println("RPM", PCP.RPM)
	fmt.Println(PCP)
}

//New structs
// BasicStats contains all of the stats common to every part
type BasicStats struct {
	// radiation absorbed dose (100 RAD = 1 Gy (gray))
	RAD int
	// Temp in deg C
	DegC int
	// Pressure in bar
	Bar int
	// Number of hours the part has been in opperation
	AgeHours int

	// Limits
	// LimVolume - Amount of collant max
	LimVolume int
	// LimRAD - Number of survivable RAD's
	LimRAD int
	// LimDegC - Max temprature
	LimDegC int
	// LimBar - Max Pressure
	LimBar int
	// LimAgeHours - Max Age
	LimAgeHours int

	// Status
	// Damage is 0-100 to denote percentage of damage to the part (100 is total fail)
	Damage int
	// Red or Yellow alarm indicator
	AlarmColor string
}

type BasicStatic struct {
	// basic stats
	BasicStats
}
type BasicDynamic struct {
	// basic stats
	BasicStats
	// Revalutions per min (this can either be passive like a turbine or active like a pump)
	RPM int
	// (KWH produced (positive) or consumed (negitive))
	KWH int
	// Requested Use in % of the max
	ReqUse int
	// Current Use in % of max
	CurUse int
	// Limits
	// LimKwH
	LimKWH int
	// LimRPM
	LimRPM int

	//Functions
	func (p BasicDynamic) Alarm() bool {
		return r.Damage > 50
	}
}
