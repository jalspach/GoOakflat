package parts

import "math/rand"

// Need to build an initial populate tool - see note at the bottom of the page
// also include a way of looking at parts in various ways all the red alarms, all the

// DynamicHW contains all of the stats common to every part
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

//LoadItem auto populates the damage...probably not used in the final product
func LoadItem(id int) *DynamicHW {
	return &DynamicHW{
		Damage: rand.Intn(100)}
}

//Alarm types or "colors"
func (b DynamicHW) Alarm() string {
	switch {
	case b.Damage >= 75:
		return "Red Alarm"
	case b.Damage >= 50:
		return "Yellow Alarm"
	case b.Damage > -15:
		return "C&W Alarm"
	default:
		return "no alarm"
	}
}

//Curstatus is the interface for current status
type Curstatus interface {
	Alarm() string
	Values() string
	Update() int
}

// need to implement https://stackoverflow.com/questions/37135193/how-to-set-default-values-in-go-structs
