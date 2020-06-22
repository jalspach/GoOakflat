package main // Basic goReact program

import (
	"fmt"
	//"parts"
)

//New structs

// BasicStats contains all of the stats common to every part
type BasicStats struct {
	RAD      int // radiation absorbed dose (100 RAD = 1 Gy (gray))
	DegC     int // Temp in deg C
	Bar      int // Pressure in bar
	Vol      int // Current volume in cu m
	AgeHours int // Number of hours the part has been in opperation

	// Limits
	LimVolume   int // LimVolume - Amount of collant max
	LimRAD      int // LimRAD - Number of survivable RAD's
	LimDegC     int // LimDegC - Max temprature
	LimBar      int // LimBar - Max Pressure
	LimAgeHours int // LimAgeHours - Max Age

	// Status
	Damage     int    // Damage is 0-100 to denote percentage of damage to the part (100 is total fail)
	AlarmColor string // Red or Yellow alarm indicator
}

type BasicStatic struct {
	BasicStats // basic stats
}

// BasicDynamic adds dynamic stats to static parts
type BasicDynamic struct {
	BasicStats     // basic stats
	RPM        int // Revs per min
	KWH        int // (KWH produced (positive) or consumed (negitive))
	ReqUse     int // Requested Use in % of the max
	CurUse     int // Current Use in % of max

	// Limits
	LimKWH int // LimKwH
	LimRPM int // LimRPM
}

//Interface
type Curstatus interface {
	Alarm() string
	Values() string
}

//Methods
func (b BasicDynamic) Alarm() string {
	switch {
	case b.Damage >= 75:
		return "Red Alarm"
	case b.Damage >= 50:
		return "Yellow Alarm"
	case b.Damage >= 25:
		return "C&W Alarm"
	default:
		return "no alarm"
	}
}

//start by defining two parts and giving them qualities. Pass data beteen them.
func main() {
	PCP := BasicDynamic{
		BasicStats: BasicStats{0, 0, 0, 0, 0, 100, 100, 100, 100, 100, 0, "Green"},
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
	fmt.Println("PCP Status is", PCP.Alarm())
}
