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

//Interface

//Curstatus is the interface for current status
type Curstatus interface {
	Alarm() string
	Values() string
	Update() int
}

//Methods

//Alarm returns the color of the alarm based on .Damage
// https://science.ksc.nasa.gov/shuttle/technology/sts-newsref/sts-caws.html
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

// Need interface and method to pass data from part to part based on other factors. i.e. based on pump speed, pass vol and temp.
//func (b DynamicHW) Update() int {
	// do I update the whole thing in one function or each part alone?
	//walk through the settings and update each in turn
	//next := b.DSPart
	//"b.DSPart".Vol = b.Vol + "b.DSPart".Vol
//}

//start by defining some parts and giving them qualities. Pass data beteen them.
func main() {
<<<<<<< HEAD
	PCP := BasicDynamic{
		BasicStats: BasicStats{0, 0, 0, 0, 0, 100, 100, 100, 100, 100, 65, "Green"},
		RPM:        0,
		KWH:        0,
		ReqUse:     0,
		CurUse:     0,
		LimKWH:     100,
		LimRPM:     500,
=======
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
		DSPart:      "SCP",
	}
	scpguid := xid.New()
	SCP := DynamicHW{
		RAD:         0,
		DegC:        98,
		Bar:         50,
		Vol:         50,
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
		DSPart:      "Tower1",
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
<<<<<<< HEAD
		DSPart:      "PCP",
=======
>>>>>>> 58ca78ebed8adc6cf0eab2f559a6a58816577beb
>>>>>>> 8ff863864e61e73f702a2afc6614b03fb94f9d9b
	}

	//SCP.Vol = SCP.Vol + PCP.Vol
	SCP.Update()

	fmt.Println("")
	fmt.Println("~~~~~~~~~~ UUID's ~~~~~~~~~")
	fmt.Println("PCP UUID:", PCP.PartUUID)
	fmt.Println("SCP UUID:", SCP.PartUUID)
	fmt.Println("Tower 1 UUID:", Tower1.PartUUID)
	fmt.Println("")
	fmt.Println("~~~~~~~~~~ PCP ~~~~~~~~~")
	fmt.Println(PCP)
	fmt.Println("Damage:", PCP.Damage)
	fmt.Println("PCP Status is:", PCP.Alarm())
	fmt.Println("")
	fmt.Println("~~~~~~~~~~ SCP ~~~~~~~~~")
	fmt.Println(SCP)
	fmt.Println("Damage:", SCP.Damage)
	fmt.Println("SCP Status is:", SCP.Alarm())
	fmt.Println("")
	fmt.Println("~~~~~~~~~~ Tower 1 ~~~~~~~~~")
	fmt.Println(Tower1)
	fmt.Println("Damage:", Tower1.Damage)
	fmt.Println("Tower 1 Status is:", Tower1.Alarm())
}
