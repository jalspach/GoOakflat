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

//Turbine is what turns the steam into rotational energy to turn the Generator.
type Turbine struct {
	FirstName, LastName string
	Location            *UserLocation
}

//Generator is rotated by the turbine to make electricity. Its output will be in Kwh
type Generator struct {
	FirstName, LastName string
	Location            *UserLocation
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

//HeatExchanger is where the heat from the primary coolant loop is used to heat the secondary coolant loop turning it to steam and turning the turbine
type HeatExchanger struct {
	FirstName, LastName string
	Location            *UserLocation
}

//Condenser is where the steam from the turbine is cooled by the tower coolant loop
type Condenser struct {
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

//PCP is the primary coolant pump. This pump circulates water from the reactorvessel through the heat HeatExchanger
type PCP struct {
	FirstName, LastName string
	Location            *UserLocation
}

//SCP is the secondary coolant pump. This pump circulates water from thge condencer back to the heat exchanger after the steam has turned the turbine.
type SCP struct {
	FirstName, LastName string
	Location            *UserLocation
}

//TCP is the tower coolant pump. this pump circulates water from the concencer through the cooling towers
type TCP struct {
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

// List of methods
func (u *User) Greeting() string {
	return fmt.Sprintf("Dear %s %s", u.FirstName, u.LastName)
}
