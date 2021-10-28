package group

import "shadygoat.eu/MatchingAPI/profile"

type Question struct {
	// UserID
	User string
	// 0 - Empty field, 1 - Sus response, 
	Reason int8
	Item string
	Index int16
}

type Group struct {
	InNeed []Question
	Members []string
	// {ID: {MID: INDIVIDUAL_SCORE}}
	Matches map[string]map[string]int8
}

func (group *Group) AddCache(user profile.User) {//TODO:
	_, ok := group.Matches[user.ID]
	group.Matches[user.ID] = map[string]int8{}
	if !ok {
		group.Members = append(group.Members, user.ID)
		group.IN_NEED = append(group.InNeed, 
			Question{User: user.ID, Type: 0, Reason: 0, Item: }
		)
	}
	for _, match := range group.Members {
		_, member := group.GetMember(match)
		matchInfo := user.Match(member)
		group.Matches[user.ID][match] = int8(matchInfo.Total*100)
		group.Matches[match][user.ID] = int8(member.Match(user).Total*100)
	}
}

func (group Group) GetMember(id string) (error, profile.User) {
	
}