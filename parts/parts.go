package parts

// list of packages to import

// list of constants
const (
	ConstExample = "const before vars"
)

// list of variables
var (
	ExportedVar    = 42
	nonExportedVar = "so say we all"
)

// Main type(s) for the file

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
}

//OLD Structs
/*
// try to keep the lowest amount of structs per file when possible.
type User struct {
	FirstName, LastName string
	Location            *UserLocation
}

type UserLocation struct {
	City    string
	Country string
}


type PipeStats struct {
	// Do we calc the capacty based on a size and length?
	Diameter int
	Length int
	// Coolant capacity in L is now a f
	CoolantCap = (Math.pi *((Diameter/2)*(Diameter / 2)) * Length)
	// Flowrate (m^3 / second) i.e. 1,000's of ltr / second
	FlowRate int
	// Alarm color (Red, Yellow)
}
func (pipe *PipeStats) Capacity() {
	(Math.pi *((PipeStats.Diameter / 2)*(PipeStats.Diameter / 2))* PipeStats.Length
}

// Turbine is what turns the steam into rotational energy to turn the Generator.
type Turbine struct {
	RPM   int
	Stats *BasicStats
}

// Generator is spun by the turbine to make electricity. Its output will be in Kwh
type Generator struct {
	KwH   float32
	RPM   int
	Stats *BasicStats
}

// ReactorVessel is the core of the reactor.
type ReactorVessel struct {
	FirstName, LastName string
	Location            *UserLocation
}

// ControlRods are used to control the reaction.
type ControlRods struct {
	FirstName, LastName string
	Location            *UserLocation
}

// ContainmentStructure is the outer building holding the reactorvessel and other parts
type ContainmentStructure struct {
	FirstName, LastName string
	Location            *UserLocation
}

// FuelRods are the actual uranium fuel
type FuelRods struct {
	FirstName, LastName string
	Location            *UserLocation
}

// Pressurizer is used to maintain the pressure in the Primary Coolant Loop
type Pressurizer struct {
	FirstName, LastName string
	Location            *UserLocation
}

// HeatExchanger is where the heat from the primary coolant loop is used to heat the secondary coolant loop turning it to steam and turning the turbine
type HeatExchanger struct {
	FirstName, LastName string
	Location            *UserLocation
}

// SteamGenerator is where the heat from the primary coolant loop is used to heat the secondary coolant loop turning it to steam and turning the turbine
type SteamGenerator struct {
	FirstName, LastName string
	Location            *UserLocation
}

// Condenser is where the steam from the turbine is cooled by the tower coolant loop
type Condenser struct {
	FirstName, LastName string
	Location            *UserLocation
}

// PCL is the Primary Coolant Loop - It represents the entire loop
type PCL struct {
	FirstName, LastName string
	Location            *UserLocation
}

// PCP is the primary coolant pump. This pump circulates water from the reactorvessel through the heat HeatExchanger
type PCP struct {
	FirstName, LastName string
	Location            *UserLocation
	MaxGPM              int
}

// PCV is the primary coolant Valve. This valve controls the flow of water from the reactorvessel through the heat HeatExchanger
type PCV struct {
	FirstName, LastName string
	Location            *UserLocation
	MaxGPM              int
}

// PPRL is the Primary Pressure Releif Valve. Pops off at a given pressure. Can be manually opened.
type PPRL struct {
	FirstName, LastName string
	Location            *UserLocation
}

// SCL is the secondary Coolant Loop - It represents the entire loop
type SCL struct {
	FirstName, LastName string
	Location            *UserLocation
}

// SCP is the secondary coolant pump. This pump circulates water from thge condencer back to the heat exchanger after the steam has turned the turbine.
type SCP struct {
	FirstName, LastName string
	Location            *UserLocation
}

// SCV is the secondary coolant Valve. This valve controls the flow of water through the secondary loop
type SCV struct {
	FirstName, LastName string
	Location            *UserLocation
	MaxGPM              int
}

// SPRL is the secondary Pressure Releif Valve. Pops off at a given pressure. Can be manually opened.
type SPRL struct {
	FirstName, LastName string
	Location            *UserLocation
}

// CoolingTower is the structure where the tower coolant loop is cooled
type CoolingTower struct {
	FirstName, LastName string
	Location            *UserLocation
}

// TowerFans are large fans in the towers which help cool the water
type TowerFans struct {
	FirstName, LastName string
	Location            *UserLocation
}

// TCL is the Tower Coolant Loop - It represents the entire loop
type TCL struct {
	FirstName, LastName string
	Location            *UserLocation
}

// TCP is the tower coolant pump. This pump circulates water from through the tower Loop
type TCP struct {
	FirstName, LastName string
	Location            *UserLocation
}

// TCV is the tower coolant Valve. This valve controls the flow of water through the tower loop
type TCV struct {
	FirstName, LastName string
	Location            *UserLocation
	MaxGPM              int
}

// PSIS Passive Safety Injection System provides emergency core cooling in LOCA (Loss-Of-Coolant Accident)
type PSIS struct {
	FirstName, LastName string
	Location            *UserLocation
}

// CVSump is the Containment Vessel Sump
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

// Convert GPM to LPM
func GPM2LPM(GPM, LPM) Float32 {
	if GPM > 0 {
		return GPM * 3.7854118
	} else if LPM > 0 {
		return LPM / 3.7854118
	} else {
		return "NA"
	}
}

// Convert flow to mass (in kg)
func Flow2Mass(LPM) float32 {
	return
}

// Compute the change in temp (in C) and mass (in Kg) (m^3/s @ 1000 kg / m^3)
func DeltaTemp(InFlowTemp, InFlowMass, ExistingTemp, ExistingMass, CoolentCapacity float32) {
	NewMass := 0.0
	NewTemp := 0.0
	NewMass = (ExistingMass - InFlowMass) + InFlowMass //this could be more accurate by taking into account the max capacity
	NewTemp = ((InFlowTemp * InflowMass) + (ExistingTemp*ExistingMaas)/NewMass)
	return NewMass, NewTemp
}

// List of methods
func (u *User) Greeting() string {
	return fmt.Sprintf("Dear %s %s", u.FirstName, u.LastName)
}
*/
