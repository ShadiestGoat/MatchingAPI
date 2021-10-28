package profile

import "math"

const HEIGHT_DIFF		float64 = 7
const WEIGHT_DIFF		float64 = 30
const POLITICAL_DIFF 	float64 = 55
const PERSONALITY_DIFF 	float64 = 0.55

type material struct {
	LooksRes,
	TraitRes,
	PersonalityRes,
	PoliticalRes,
	BDSMRes float64
}

type MatchInfo struct {
	Materials material
	Total     float64
}

func checkSexuality(user Gender, match Orientation) bool {
	switch user {
	case 0:
		return match.Male
	case 1:
		return match.Female
	case 2:
		return match.Fluid
	case 3:
		return match.Other
	default:
		return false
	}
}

func (user User) Match(match User) MatchInfo {
	var matchHere MatchInfo = MatchInfo{}
	matchHere.Materials.LooksRes = user.Profile.hotness(match.Profile)
	matchHere.Materials.TraitRes = user.Profile.traitMatch(match.Profile)

	if match.DONE_BDSM {
		matchHere.Materials.BDSMRes = user.Profile.Data.BDSM.bdsmMatch(match.Profile.Data.BDSM)
	} else {
		matchHere.Materials.BDSMRes = -1
	}
	if match.DONE_POLITICS {
		matchHere.Materials.PoliticalRes = user.Profile.politicalMatch(match.Profile)
	} else {
		matchHere.Materials.PersonalityRes = -1
	}
	if match.DONE_PERSONALITY {
		matchHere.Materials.PersonalityRes = user.Profile.personalityMatch(match.Profile)
	} else {
		matchHere.Materials.PersonalityRes = -1
	}
	if !checkSexuality(match.Gender, user.Orientation) {
		matchHere.Total = 0
	} else {
		matchHere.Total = matchHere.Materials.LooksRes*user.Profile.Weights.Major.Looks +
			matchHere.Materials.PersonalityRes*user.Profile.Weights.Major.Personality +
			matchHere.Materials.TraitRes*user.Profile.Weights.Major.Traits +
			matchHere.Materials.BDSMRes*user.Profile.Weights.Major.BDSM +
			matchHere.Materials.PoliticalRes*user.Profile.Weights.Major.Politics
	}
	return matchHere
}

func (user Profile) personalityMatch(match Profile) float64 {
	var res float64 = 0
	coldnessDiff := math.Abs(match.Data.Personality.Coldness - user.Preferance.Personality.Coldness)
	if coldnessDiff > PERSONALITY_DIFF {
		res += (1 - coldnessDiff) * user.Preferance.Personality.Coldness
	}
	assertiveDiff := math.Abs(match.Data.Personality.Assertive - user.Preferance.Personality.Assertive)
	if assertiveDiff > PERSONALITY_DIFF {
		res += (1 - assertiveDiff) * user.Preferance.Personality.Assertive
	}
	introversionDiff := math.Abs(match.Data.Personality.Introversion - user.Preferance.Personality.Introversion)
	if introversionDiff > PERSONALITY_DIFF {
		res += (1 - introversionDiff) * user.Preferance.Personality.Introversion
	}
	intuitantDiff := math.Abs(match.Data.Personality.Intuitant - user.Preferance.Personality.Intuitant)
	if intuitantDiff > PERSONALITY_DIFF {
		res += (1 - intuitantDiff) * user.Preferance.Personality.Intuitant
	}
	return res
}

func (user Profile) traitMatch(match Profile) float64 {
	var res float64
	res += math.Abs(match.Data.Traits.Funny			.SureAvg()-user.Preferance.Traits.Funny) * user.Weights.Traits.Funny
	res += math.Abs(match.Data.Traits.Cute			.SureAvg()-user.Preferance.Traits.Cute) * user.Weights.Traits.Cute
	res += math.Abs(match.Data.Traits.Kind			.SureAvg()-user.Preferance.Traits.Kind) * user.Weights.Traits.Kind
	res += math.Abs(match.Data.Traits.Intelligence	.SureAvg()-user.Preferance.Traits.Intelligence) * user.Weights.Traits.Intelligence
	res += math.Abs(match.Data.Traits.Rich			.SureAvg()-user.Preferance.Traits.Rich) * user.Weights.Traits.Rich
	res += math.Abs(match.Data.Traits.Creative		.SureAvg()-user.Preferance.Traits.Creative) * user.Weights.Traits.Creative
	res += math.Abs(match.Data.Traits.Maturity		.SureAvg()-user.Preferance.Traits.Maturity) * user.Weights.Traits.Maturity
	// TODO: 
	var bad float64 = math.Abs(Avg(match.Data.Traits.Annoyingness) - user.Preferance.Traits.Annoyingness)
	bad += math.Abs(Avg(match.Data.Traits.Creepiness) - user.Preferance.Traits.Creepiness)
	res -= bad * user.Weights.Major.BadTraits
	return res
}

func (user Profile) hotness(match Profile) float64 {
	var res float64 = 0
	var heightDiff = math.Abs(float64(user.Preferance.Looks.Height) - match.Data.Looks.Height.SureAvg())
	var weightDiff = math.Abs(float64(user.Preferance.Looks.Figure) - match.Data.Looks.Figure.SureAvg())

	if heightDiff <= HEIGHT_DIFF {
		res += (HEIGHT_DIFF - heightDiff) / HEIGHT_DIFF * user.Weights.Looks.Height
	}
	if weightDiff <= WEIGHT_DIFF {
		res += (WEIGHT_DIFF - weightDiff) / WEIGHT_DIFF * user.Weights.Looks.Figure
	}
	res += match.Data.Looks.Face.SureAvg() * user.Weights.Looks.Face
	res += match.Data.Looks.Fashion.SureAvg() * user.Weights.Looks.Fashion
	res += match.Data.Looks.Muscular.SureAvg() * user.Weights.Looks.Muscular
	res += match.Data.Looks.HairCoolness.SureAvg() * user.Weights.Looks.HairCoolness
	return res
}

func (user Profile) politicalMatch(match Profile) float64 {
	var diffXY [2]float64 = [2]float64{
		match.Data.Political[0] - user.Preferance.Political[0],
		match.Data.Political[1] - user.Preferance.Political[1],
	}
	diff := math.Sqrt(diffXY[0]*diffXY[0] + diffXY[1]*diffXY[1])
	if diff <= POLITICAL_DIFF {
		return (POLITICAL_DIFF - diff) / POLITICAL_DIFF
	} else {
		return 0
	}
}

/*
	The function for the formula

	k  - The current kink value (ie. name's)

	d  - The desired kink value (ie. name's desired kink (for dominant it'd be the value of the current match's submissive))

	k2 - The current kink but of match name's. If its ==-1,then it will be ignored

	s  - Name's sum of all kink values
*/
func formula(k float64, d float64, k2 float64, s float64) float64 {
	var rV float64 = 0

	if k == 0 && d > 40 {
		return 1 + d*-1.2*d/10
	} //punish if the person absolutly hates it. If k=0,it must be absolute hate!

	kDiff := math.Abs(k - d)

	rV = (2 - kDiff) * (k / s) //scale with how important k is to the person. The absolute different basucally means that the smaller the big boy is,the better.

	if kDiff > 0.42 { // actively punish big differences
		if k < 0.13 { // if k<13,it must mean you really hate it. Punish it more!
			rV *= -0.45
		} else {
			rV *= -0.35
		}
	}

	if k2 != -1 { //ignore if ignore
		KsDiff := math.Abs(k - k2) // basically if 2 people are dominat,it wont work out,so we are punishing it ^^
		if KsDiff < 0.20 && k2 > 0.25 && k > 0.25 {
			v := 0.7
			if formula(k, d, -1, s) < 0 {
				v += 0.65
			}
			rV *= v
		}
	}
	return rV
}

func (user BDSM) bdsmMatch(match BDSM) float64 {
	var matchNum, tot float64 = 0, 0
	selfS := user.Vanilla +
		user.NonMonogamist +
		user.Experimentalist +
		user.Dominant +
		user.Switch +
		user.Submissive +
		user.Degrader +
		user.Brat +
		user.BratTamer +
		user.Owner +
		user.RopeBunny +
		user.Degradee +
		user.Voyeur +
		user.MasterMistress +
		user.Exhibitionist +
		user.Sadist +
		user.Slave +
		user.Hunter +
		user.Pet +
		user.DaddyMommy +
		user.Prey +
		user.Masochist +
		user.BoyGirl +
		user.Ageplayer +
		user.Rigger

	matchNum += formula(user.Submissive, match.Dominant, match.Submissive, selfS)
	tot += formula(user.Submissive, user.Submissive, -1, selfS)
	matchNum += formula(user.RopeBunny, match.Rigger, match.RopeBunny, selfS)
	tot += formula(user.RopeBunny, user.RopeBunny, -1, selfS)
	matchNum += formula(user.Masochist, match.Sadist, match.Masochist, selfS)
	tot += formula(user.Masochist, user.Masochist, -1, selfS)
	matchNum += formula(user.Brat, match.BratTamer, match.Brat, selfS)
	tot += formula(user.Brat, user.Brat, -1, selfS)
	matchNum += formula(user.Degradee, match.Degrader, match.Degradee, selfS)
	tot += formula(user.Degradee, user.Degradee, -1, selfS)
	matchNum += formula(user.Slave, match.MasterMistress, match.Slave, selfS)
	tot += formula(user.Slave, user.Slave, -1, selfS)
	matchNum += formula(user.Experimentalist, match.Experimentalist, -1, selfS)
	tot += formula(user.Experimentalist, user.Experimentalist, -1, selfS)
	matchNum += formula(user.Prey, match.Hunter, match.Prey, selfS)
	tot += formula(user.Prey, user.Prey, -1, selfS)
	matchNum += formula(user.NonMonogamist, match.NonMonogamist, -1, selfS)
	tot += formula(user.NonMonogamist, user.NonMonogamist, -1, selfS)
	matchNum += formula(user.Vanilla, match.Vanilla, -1, selfS)
	tot += formula(user.Vanilla, user.Vanilla, -1, selfS)
	matchNum += formula(user.Pet, match.Owner, match.Pet, selfS)
	tot += formula(user.Pet, user.Pet, -1, selfS)
	matchNum += formula(user.Exhibitionist, match.Voyeur, match.Exhibitionist, selfS)
	tot += formula(user.Exhibitionist, user.Exhibitionist, -1, selfS)
	matchNum += formula(user.Voyeur, match.Exhibitionist, match.Voyeur, selfS)
	tot += formula(user.Voyeur, user.Voyeur, -1, selfS)
	matchNum += formula(user.Hunter, match.Prey, match.Hunter, selfS)
	tot += formula(user.Hunter, user.Hunter, -1, selfS)
	matchNum += formula(user.DaddyMommy, match.BoyGirl, match.DaddyMommy, selfS)
	tot += formula(user.DaddyMommy, user.DaddyMommy, -1, selfS)
	matchNum += formula(user.Rigger, match.RopeBunny, match.Rigger, selfS)
	tot += formula(user.Rigger, user.Rigger, -1, selfS)
	matchNum += formula(user.Owner, match.Pet, match.Owner, selfS)
	tot += formula(user.Owner, user.Owner, -1, selfS)
	matchNum += formula(user.MasterMistress, match.Slave, match.MasterMistress, selfS)
	tot += formula(user.MasterMistress, user.MasterMistress, -1, selfS)
	matchNum += formula(user.Sadist, match.Masochist, match.Sadist, selfS)
	tot += formula(user.Sadist, user.Sadist, -1, selfS)
	matchNum += formula(user.BoyGirl, match.DaddyMommy, match.BoyGirl, selfS)
	tot += formula(user.BoyGirl, user.BoyGirl, -1, selfS)
	matchNum += formula(user.Dominant, match.Submissive, match.Dominant, selfS)
	tot += formula(user.Dominant, user.Dominant, -1, selfS)
	matchNum += formula(user.Ageplayer, match.Ageplayer, -1, selfS)
	tot += formula(user.Ageplayer, user.Ageplayer, -1, selfS)
	matchNum += formula(user.BratTamer, match.Brat, match.BratTamer, selfS)
	tot += formula(user.BratTamer, user.BratTamer, -1, selfS)
	matchNum += formula(user.Degrader, match.Degradee, match.Degrader, selfS)
	tot += formula(user.Degrader, user.Degrader, -1, selfS)
	matchNum += formula(user.Switch, match.Switch, -1, selfS)
	tot += formula(user.Switch, user.Switch, -1, selfS)

	if matchNum < 0 {
		return 0
	} else {
		return (matchNum / tot) * 100
	}
}