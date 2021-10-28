package profile

type User struct {
	Name,
	ID string
	Gender      Gender
	Orientation Orientation
	DONE_BDSM,
	DONE_POLITICS,
	DONE_PERSONALITY bool
	Profile Profile
	Groups  map[string]GroupSettings
}

type GroupSettings struct {
	USE_BDSM bool
}

type Orientation struct {
	Male,
	Female,
	Fluid,
	Other bool
}

// M - 0 F-1 Fluid - 2 Other - 3
type Gender int8

const (
	Male Gender = iota
	Female
	Fluid
	Other
)