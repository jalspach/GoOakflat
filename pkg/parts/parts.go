package parts

import "math/rand"

//New structs
// BasicStats contains all of the stats common to every part
type DynamicHW struct {
	RAD      int // radiation absorbed dose (100 RAD = 1 Gy (gray))
	DegC     int // Temp in deg C
	Bar      int // Pressure in bar
	Vol      int // Current volume in cu m
	AgeHours int // Number of hours the part has been in opperation
	Building string

	// Limits
	LimVolume   int // LimVolume - Amount of collant max
	LimRAD      int // LimRAD - Number of survivable RAD's
	LimDegC     int // LimDegC - Max temprature
	LimBar      int // LimBar - Max Pressure
	LimAgeHours int // LimAgeHours - Max Age

	// Status
	Damage int // Damage is 0-100 to denote percentage of damage to the part (100 is total fail)
	RPM    int // Revs per min
	KWH    int // (KWH produced (positive) or consumed (negitive))
	ReqUse int // Requested Use in % of the max
	CurUse int // Current Use in % of max

	// Limits
	LimKWH int // LimKwH
	LimRPM int // LimRPM

	//order of operations

	//PartUUID string //unique ID for this part
	PartUUID string
	USPart   string //Previous part in the chain by name
	DSPart   string //Next part in the chain by name
}

func LoadItem(id int) *DynamicHW {
	return &DynamicHW{
		Damage: rand.Intn(100)}
}
