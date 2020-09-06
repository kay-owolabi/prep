package graphs

type MatchResult struct {
	winningTeam string
	losingTeam  string
}

type voidStruct struct{}
type stringSet map[string]voidStruct

var void voidStruct

func (m MatchResult) New(winningTeam, losingTeam string) *MatchResult {
	return &MatchResult{winningTeam, losingTeam}
}

func canTeamABeatTeamB(matches []*MatchResult, teamA, teamB string) bool {
	return isReachableDfs(BuildGraph(matches), teamA, teamB, make(stringSet))
}

func BuildGraph(matches []*MatchResult) map[string]stringSet {
	graph := map[string]stringSet{}
	for _, match := range matches {
		graph[match.winningTeam][match.losingTeam] = void
	}
	return graph
}

func isReachableDfs(graph map[string]stringSet, curr, dest string, visitedSet stringSet) bool {
	_, visited := visitedSet[curr]
	beatenTeams, containsCurr := graph[curr]
	if curr == dest {
		return true
	} else if visited || !containsCurr {
		return false
	}
	visitedSet[curr] = void

	for team, _ := range beatenTeams {
		if isReachableDfs(graph, team, dest, visitedSet) {
			return true
		}
	}
	return false
}
