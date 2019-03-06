package orggen

var (
	sizeClasses = []string{
  "small",
  "medium",
  "big",
	}

	sizeClassData = map[string]SizeClass{
  "small": {
  	"small",
  	6,
  	20,
  },
  "medium": {
  	"medium",
  	21,
  	50,
  },
  "big": {
  	"big",
  	51,
  	100,
  },
	}

	types = []string{
  "adventuring company",
  "church",
  "guild",
  "mercenary company",
	}

	typeData = map[string]OrganizationType{
  "adventuring company": {
  	"an adventuring company",
  	[]string{
    "reckless",
    "selfless",
    "aggressive",
    "avaricious",
    "curious",
    "forceful",
  	},
  	false,
  	"captain",
  	[]string{
    "Silver",
    "White",
    "Black",
    "Crimson",
    "Free",
    "Burning",
  	},
  	[]string{
    "Wolves",
    "Swords",
    "Axes",
    "Giants",
    "Lords",
    "Dragons",
    "Wyverns",
    "Company",
  	},
  	"{{.FirstPart}} {{.SecondPart}}",
  },
  "church": {
  	"a church",
  	[]string{
    "penitent",
    "helpful",
    "selfless",
    "controlling",
    "merciful",
    "merciless",
    "penniless",
    "proselytizing",
  	},
  	true,
  	"high priest",
  	[]string{
    "Holy",
    "Glorious",
    "Exalted",
    "Humble",
    "Penitent",
    "Righteous",
  	},
  	[]string{
    "Sword",
    "Flame",
    "Dove",
    "Forest",
    "Meek",
    "Truth",
    "Light",
  	},
  	"{{.FirstPart}} Church of the {{.SecondPart}}",
  },
  "guild": {
  	"a guild",
  	[]string{
    "avaricious",
    "ambitious",
    "productive",
    "lazy",
    "manipulative",
    "frugal",
    "observant",
  	},
  	true,
  	"guild leader",
  	[]string{
    "Incorporated",
    "August",
    "East Wind",
    "West Wind",
    "South Wind",
    "North Wind",
    "Global",
  	},
  	[]string{
    "Artist",
    "Furrier",
    "Carpenter",
    "Merchant",
    "Tanner",
    "Moneylender",
  	},
  	"{{.FirstPart}} {{.SecondPart}}'s Guild",
  },
  "mercenary company": {
  	"a mercenary company",
  	[]string{
    "aggressive",
    "belligerent",
    "avaricious",
    "neutral",
    "honorable",
    "dishonorable",
    "trustworthy",
  	},
  	true,
  	"captain",
  	[]string{
    "Silver",
    "White",
    "Black",
    "Crimson",
    "Free",
    "Burning",
  	},
  	[]string{
    "Wolves",
    "Swords",
    "Axes",
    "Giants",
    "Lords",
    "Dragons",
    "Wyverns",
    "Company",
  	},
  	"{{.FirstPart}} {{.SecondPart}}",
  },
	}
)
