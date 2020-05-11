package parts

// list of packages to import
import (
	"fmt"
)

// list of constants
const (
	ConstExample = "const before vars"
)

// list of variables
var (
	ExportedVar    = 42
	nonExportedVar = "so say we all"
	temp           = 0
)

// Main type(s) for the file,
// try to keep the lowest amount of structs per file when possible.
type User struct {
	FirstName, LastName string
	Location            *UserLocation
}

type UserLocation struct {
	City    string
	Country string
}

//BasicStats contains all of the stats common to each part
type BasicStats struct {
	// radiation absorbed dose (100 RAD = 1 Gy (gray))
	RAD     int
	//Temp in deg C
	DegC    float32
	//Pressure in bar
	Bar		float32
	//Number of hours the part has been in opperation
	AgeHours int
	//DamageStatus is 0-100 to denote percentage of damage to the part (100% is total fail)
	DamageStatus int
	//Coolant capacity
	CoolantCap int
}

//Turbine is what turns the steam into rotational energy to turn the Generator.
type Turbine struct {
	RPM   int
	Stats *BasicStats
}

//Generator is spun by the turbine to make electricity. Its output will be in Kwh
type Generator struct {
	KwH   float32
	RPM   int
	Stats *BasicStats
}

//ReactorVessel is the core of the reactor.
type ReactorVessel struct {
	FirstName, LastName string
	Location            *UserLocation
}

//ControlRods are used to control the reaction.
type ControlRods struct {
	FirstName, LastName string
	Location            *UserLocation
}

//ContainmentStructure is the outer building holding the reactorvessel and other parts
type ContainmentStructure struct {
	FirstName, LastName string
	Location            *UserLocation
}

//FuelRods are the actual uranium fuel
type FuelRods struct {
	FirstName, LastName string
	Location            *UserLocation
}

//Pressurizer is used to maintain the pressure in the Primary Coolant Loop
type Pressurizer struct {
	FirstName, LastName string
	Location            *UserLocation
}

//HeatExchanger is where the heat from the primary coolant loop is used to heat the secondary coolant loop turning it to steam and turning the turbine
type HeatExchanger struct {
	FirstName, LastName string
	Location            *UserLocation
}

//SteamGenerator is where the heat from the primary coolant loop is used to heat the secondary coolant loop turning it to steam and turning the turbine
type SteamGenerator struct {
	FirstName, LastName string
	Location            *UserLocation
}

//Condenser is where the steam from the turbine is cooled by the tower coolant loop
type Condenser struct {
	FirstName, LastName string
	Location            *UserLocation
}

//PCL is the Primary Coolant Loop - It represents the entire loop
type PCL struct {
	FirstName, LastName string
	Location            *UserLocation
}

//PCP is the primary coolant pump. This pump circulates water from the reactorvessel through the heat HeatExchanger
type PCP struct {
	FirstName, LastName string
	Location            *UserLocation
	MaxGPM              int
}

//PCV is the primary coolant Valve. This valve controls the flow of water from the reactorvessel through the heat HeatExchanger
type PCV struct {
	FirstName, LastName string
	Location            *UserLocation
	MaxGPM              int
}

//PPRL is the Primary Pressure Releif Valve. Pops off at a given pressure. Can be manually opened.
type PPRL struct {
	FirstName, LastName string
	Location            *UserLocation
}

//SCL is the secondary Coolant Loop - It represents the entire loop
type SCL struct {
	FirstName, LastName string
	Location            *UserLocation
}

//SCP is the secondary coolant pump. This pump circulates water from thge condencer back to the heat exchanger after the steam has turned the turbine.
type SCP struct {
	FirstName, LastName string
	Location            *UserLocation
}

//SCV is the secondary coolant Valve. This valve controls the flow of water through the secondary loop
type SCV struct {
	FirstName, LastName string
	Location            *UserLocation
	MaxGPM              int
}

//SPRL is the secondary Pressure Releif Valve. Pops off at a given pressure. Can be manually opened.
type SPRL struct {
	FirstName, LastName string
	Location            *UserLocation
}

//CoolingTower is the structure where the tower coolant loop is cooled
type CoolingTower struct {
	FirstName, LastName string
	Location            *UserLocation
}

//TowerFans are large fans in the towers which help cool the water
type TowerFans struct {
	FirstName, LastName string
	Location            *UserLocation
}

//TCL is the Tower Coolant Loop - It represents the entire loop
type TCL struct {
	FirstName, LastName string
	Location            *UserLocation
}

//TCP is the tower coolant pump. This pump circulates water from through the tower Loop
type TCP struct {
	FirstName, LastName string
	Location            *UserLocation
}

//TCV is the tower coolant Valve. This valve controls the flow of water through the tower loop
type TCV struct {
	FirstName, LastName string
	Location            *UserLocation
	MaxGPM              int
}

//PSIS Passive Safety Injection System provides emergency core cooling in LOCA (Loss-Of-Coolant Accident)
type PSIS struct {
	FirstName, LastName string
	Location            *UserLocation
}

//CVSump is the Containment Vessel Sump
type CVSump struct {
	FirstName, LastName string
	Location            *UserLocation
}

// List of functions
func NewUser(firstName, lastName string) *User {
	return &User{FirstName: firstName,
		LastName: lastName,
		Location: &UserLocation{
			City:    "Santa Monica",
			Country: "USA",
		},
	}
}
//Convert GPM to LPM
func GPM2LPM(GPM, LPM) Float32 {
	if GPM > 0 {
	return GPM * 3.7854118
	}
	else if LPM > 0 {
		return LPM / 3.7854118
	}
}
//Convert flow to mass (in kg)
func Flow2Mass (LPM) float32 {
	return 
}
//Compute the change in temp (in C) and mass (in Kg) (m^3/s @ 1000 kg / m^3)
func DeltaTemp(InFlowTemp, InFlowMass, ExistingTemp, ExistingMass, CoolentCapacity float32) {
	NewMass := 0.0
	NewTemp := 0.0
	NewMass = (InFlowMass + ExistingMass) - OutFlowMass
	return &Temp{newTemp: newTemp,ExistingMass
		newMass: newMass,
		},
		((Mass1 * Temp1)+(Mass2 * Temp2))/(Mass1+Mass2)
	}
}
// List of methods
func (u *User) Greeting() string {
	return fmt.Sprintf("Dear %s %s", u.FirstName, u.LastName)
}
