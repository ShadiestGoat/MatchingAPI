package profile

type Profile struct {
	Data ProfileData
	Preferance ProfilePref
	Weights Weights
}

type ProfileData struct {
	BDSM        BDSM
	Personality Personality
	Traits      Traits
	Looks       Looks
	Political   Political
}

type ProfilePref struct {
	Personality Personality
	Traits      WeightTraits
	Looks       PrefLooks
	Political   Political
}

type PrefLooks struct {
	Figure,
	Height int16
}

type Personality struct {
	// Thinking & Feeling
	Coldness,
	// Turbulant vs Assertive
	Assertive,
	// Introversion to Extroversion
	Introversion,
	// Intuition vs Observation
	Intuitant float64
}

type BDSM struct {
	Vanilla,
	NonMonogamist,
	Experimentalist,
	Dominant,
	Switch,
	Submissive,
	Degrader,
	Brat,
	BratTamer,
	Owner,
	RopeBunny,
	Degradee,
	Voyeur,
	MasterMistress,
	Exhibitionist,
	Sadist,
	Slave,
	Hunter,
	Pet,
	DaddyMommy,
	Prey,
	Masochist,
	BoyGirl,
	Ageplayer,
	Rigger float64
}

type Traits struct {
	Funny,
	Cute,
	Kind,
	Intelligence,
	Rich,
	Creative,
	Maturity,

	Annoyingness,
	Creepiness []VoteFloat
}

type Looks struct {
	Face,
	Fashion,
	Muscular,
	HairCoolness []VoteFloat

	// Save in CM for wasteline
	Figure,
	// Save in CM!
	Height []VoteInt
}

type Political [2]float64

type VoteRoot struct {
	Length int16
	UnsureLength int16
	Total float64
	UnsureTotal float64
	// USER ID : Vote
	Votes map[string]Vote
	/* USER ID : SCORE 
		Score is votes for - votes against. A user will be removed from here if their score goes above TODO:
	*/
	Unsure map[string]float32
}

type Vote struct {
	Value      float64
	ReportsNeg int16
	ReportsPos int16
	Extreme    bool
}