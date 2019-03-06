package orggen

import (
	"bytes"
	"html/template"
	"math/rand"

	"github.com/ironarachne/utility"
)

// NameData is name data
type NameData struct {
	FirstPart  string
	SecondPart string
}

// Organization is an organization
type Organization struct {
	Name         string
	Type         OrganizationType
	LeaderType   string
	SizeClass    SizeClass
	Size         int
	PrimaryTrait string
}

// OrganizationType is a type of organization
type OrganizationType struct {
	Name            string
	PossibleTraits  []string
	CanBeLedByGroup bool
	LeaderTitle     string
	NameFirstParts  []string
	NameSecondParts []string
	NameTemplate    string
}

// SizeClass is a size range
type SizeClass struct {
	Name    string
	MinSize int
	MaxSize int
}

func (org Organization) setLeaderType() {
	org.LeaderType = "individual"

	if org.Type.CanBeLedByGroup {
		x := rand.Intn(10)

		if x >= 9 {
			org.LeaderType = "group"
		}
	}
}

func (org Organization) setName() {
	var tplOutput bytes.Buffer

	firstPart := utility.RandomItem(org.Type.NameFirstParts)
	secondPart := utility.RandomItem(org.Type.NameSecondParts)
	nameData := NameData{firstPart, secondPart}

	tmpl, err := template.New("orgname").Parse(org.Type.NameTemplate)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(&tplOutput, nameData)
	if err != nil {
		panic(err)
	}
	name := tplOutput.String()

	org.Name = name
}

func (org Organization) setTrait() {
	org.PrimaryTrait = utility.RandomItem(org.Type.PossibleTraits)
}

// Generate generates a org
func Generate() Organization {
	org := Organization{}

	org.Type = typeData[utility.RandomItem(types)]
	org.SizeClass = sizeClassData[utility.RandomItem(sizeClasses)]
	org.Size = rand.Intn(org.SizeClass.MaxSize) + org.SizeClass.MinSize
	org.setName()
	org.setLeaderType()
	org.setTrait()

	return org
}
