package profile

import "math"

func validatePerc(items []*float64) {
	var total float64 = 0
	for _, k := range items {
		total += math.Abs(*k)
	}
	var factor = 1 / total
	for i, k := range items {
		*items[i] = math.Abs(*k * factor)
	}
}
func (user *User) Validate() {
	if !user.DONE_BDSM {
		user.Profile.Weights.Major.BDSM = 0
	}
	if !user.DONE_PERSONALITY {
		user.Profile.Weights.Major.Personality = 0
	}
	if !user.DONE_POLITICS {
		user.Profile.Weights.Major.Politics = 0
	}
	user.Profile.Weights.Validate()
	user.Profile.Data.BDSM.Validate()
}

func (weights *Weights) Validate() {
	validatePerc([]*float64{
		&weights.Major.BDSM,
		&weights.Major.Looks,
		&weights.Major.Personality,
		&weights.Major.Politics,
		&weights.Major.Traits,
	})
	if math.Abs(weights.Major.BadTraits) > 1 {
		weights.Major.BadTraits = 1
	} else if weights.Major.BadTraits < 0 {
		weights.Major.BadTraits *= -1
	}

	validatePerc([]*float64{
		&weights.Looks.Face,
		&weights.Looks.Fashion,
		&weights.Looks.Figure,
		&weights.Looks.HairCoolness,
		&weights.Looks.Height,
		&weights.Looks.Muscular,
	})

	validatePerc([]*float64{
		&weights.Personality.Assertive,
		&weights.Personality.Coldness,
		&weights.Personality.Introversion,
		&weights.Personality.Intuitant,
	})

	validatePerc([]*float64{
		&weights.Traits.Creative,
		&weights.Traits.Cute,
		&weights.Traits.Funny,
		&weights.Traits.Intelligence,
		&weights.Traits.Kind,
		&weights.Traits.Maturity,
		&weights.Traits.Rich,
	})
}

func (res *BDSM) Validate() {
	validatePerc([]*float64{
		&res.Vanilla,
		&res.NonMonogamist,
		&res.Experimentalist,
		&res.Dominant,
		&res.Switch,
		&res.Submissive,
		&res.Degrader,
		&res.Brat,
		&res.BratTamer,
		&res.Owner,
		&res.RopeBunny,
		&res.Degradee,
		&res.Voyeur,
		&res.MasterMistress,
		&res.Exhibitionist,
		&res.Sadist,
		&res.Slave,
		&res.Hunter,
		&res.Pet,
		&res.DaddyMommy,
		&res.Prey,
		&res.Masochist,
		&res.BoyGirl,
		&res.Ageplayer,
		&res.Rigger,
	})
}

