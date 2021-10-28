package profile

func (vote VoteRoot) SureAvg() float64 {
	return vote.SureTotal/float64(vote.SureLength)
}

func (vote *VoteRoot) AddVote(Author string, Value float64, groupMembersLen int64) {
	var extreme bool = false //TODO: 
	vote.Votes[Author] = Vote{
		Value: Value,
		Extreme: extreme,
	}
	vote.Length++
	vote.Total += Value
	if !extreme {
		vote.SureLength++
		vote.SureTotal += Value
	} else {
		vote.Unsure[Author] = 0
	}
} 