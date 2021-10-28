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
	ego,
	Annoyingness,
	Creepiness VoteRoot
}

type Looks struct {
	Face,
	Fashion,
	Muscular,
	HairCoolness VoteRoot

	// Save in CM for wasteline
	Figure,
	// Save in CM!
	Height VoteRoot
}

type Political [2]float64

type VoteRoot struct {
	Length int16
	SureLength int16
	Total float64
	SureTotal float64
	// USER ID : Vote
	Votes map[string]Vote //prevent users from double voting
	/* USER ID : SCORE 
		Score is votes for - votes against. A user will be removed from here if their score goes above TODO:
		Vote score can be 0.5, so its a float32
	*/
	Unsure map[string]float32
}

type Vote struct {
	Value      float64
	ReportsNeg int16
	ReportsPos int16
	Extreme    bool
}