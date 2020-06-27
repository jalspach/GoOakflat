package main // Basic goReact program

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/rs/xid"
)

//New structs

// DynamicHW Currently all hardware is dymanic
type DynamicHW struct {
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
	Damage int // Damage is 0-100 to denote percentage of damage to the part (100 is total fail)
	RPM    int // Revs per min
	KWH    int // (KWH produced (positive) or consumed (negitive))
	ReqUse int // Requested Use in % of the max
	CurUse int // Current Use in % of max

	// Limits
	LimKWH int // LimKwH
	LimRPM int // LimRPM

	//order

	PartUUID string //unique ID for this part
	USPart   string //Previous part in the chain by name
	DSPart   string //Next part in the chain by name
}

//Interface

//Curstatus is the interface for current status
type Curstatus interface {
	Alarm() string
	Values() string
}

//Methods

//Alarm returns the color of the alarm based on .Damage
func (b DynamicHW) Alarm() string {
	switch {
	case b.Damage >= 75:
		return "Red Alarm"
	case b.Damage >= 50:
		return "Yellow Alarm"
	case b.Damage >= 15:
		return "C&W Alarm"
	default:
		return "no alarm"
	}
}

//start by defining two parts and giving them qualities. Pass data beteen them.
func main() {
	rand.Seed(time.Now().UnixNano())
	pcpguid := xid.New()
	PCP := DynamicHW{
		RAD:         0,
		DegC:        98,
		Bar:         50,
		Vol:         100,
		AgeHours:    20,
		PartUUID:    pcpguid.String(),
		LimRAD:      50,
		LimDegC:     50,
		LimBar:      50,
		LimAgeHours: 50,
		Damage:      rand.Intn(100),
		RPM:         0,
		KWH:         0,
		ReqUse:      0,
		CurUse:      0,
		LimKWH:      100,
		LimRPM:      500,
	}
	scpguid := xid.New()
	SCP := DynamicHW{
		RAD:         0,
		DegC:        98,
		Bar:         50,
		Vol:         100,
		AgeHours:    20,
		PartUUID:    scpguid.String(),
		LimRAD:      50,
		LimDegC:     50,
		LimBar:      50,
		LimAgeHours: 50,
		Damage:      rand.Intn(100),
		RPM:         0,
		KWH:         0,
		ReqUse:      0,
		CurUse:      0,
		LimKWH:      100,
		LimRPM:      500,
	}
	tower1guid := xid.New()
	Tower1 := DynamicHW{
		RAD:         0,
		DegC:        98,
		Bar:         50,
		Vol:         100,
		AgeHours:    20,
		PartUUID:    tower1guid.String(),
		LimRAD:      50,
		LimDegC:     50,
		LimBar:      50,
		LimAgeHours: 50,
		Damage:      rand.Intn(100),
		RPM:         0,
		KWH:         0,
		ReqUse:      0,
		CurUse:      0,
		LimKWH:      100,
		LimRPM:      500,
	}

	fmt.Println(PCP)
	fmt.Println("PCP UUID:", PCP.PartUUID)
	fmt.Println("SCP UUID:", SCP.PartUUID)
	fmt.Println("Tower 1 UUID:", Tower1.PartUUID)

	fmt.Println("Damage:", PCP.Damage)
	fmt.Println("PCP Status is:", PCP.Alarm())
	fmt.Println("Damage:", SCP.Damage)
	fmt.Println("SCP Status is:", SCP.Alarm())
	fmt.Println("Damage:", Tower1.Damage)
	fmt.Println("Tower 1 Status is:", Tower1.Alarm())
}
