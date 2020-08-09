package main // Basic goReact program

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/jalspach/GoOakflat/pkg/parts"

	"github.com/rs/xid"
)

//start by defining some parts and giving them qualities. Pass data beteen them.
func main() {
	rand.Seed(time.Now().UnixNano())
	pcpguid := xid.New()
	PCP := parts.DynamicHW{
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
	SCP := parts.DynamicHW{
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
	Tower1 := parts.DynamicHW{
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
		DSPart:      "PCP",
	}

	//SCP.Vol = SCP.Vol + PCP.Vol
	// SCP.Update()

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
