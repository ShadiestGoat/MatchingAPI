package profile

type Weights struct {
	Major struct {
		BDSM,
		Looks,
		Traits,
		Personality,
		BadTraits,
		Politics float64
	}
	Traits      WeightTraits
	Looks       WeightLooks
	Personality Personality
}

type WeightTraits struct {
	Funny,
	Cute,
	Kind,
	Intelligence,
	Rich,
	Creative,
	Maturity,

	Annoyingness,
	Creepiness float64
}

type WeightLooks struct {
	Face,
	Fashion,
	Muscular,
	HairCoolness,

	Figure,
	Height float64
}

