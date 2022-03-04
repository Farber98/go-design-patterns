package flyweight

import "time"

/*
FLYWEIGHT:
- Allows sharing the state of a heavy object between many instances of some type.

OBJECTIVE:
- Share all posible states of objects in a single common object, minimizing object creation by using pointers to already created objects.

EXAMPLE: Simulate final match of European Championship.
- Store each team's information just once, and deliver references to them to the users.
- If we face a million users trying to access information about a match, we will actually just have two teams in memory with a million pointers to the same memory direction.

ACCEPTANCE CRITERIA:
- Must always reduce the amount of memory that is used, and must be focused primarly on this objective.
- We wiell create a Team struct with some basic information such as the team's name, players, historical results and an image depicting their shield.
- We must ensure correct team creation and not having duplicates
- When creating the same team twice, we must have two pointers pointing to the same memory address.

*/

const (
	TEAM_A = iota
	TEAM_B
)

type Player struct {
	Name         string
	Surname      string
	PreviousTeam uint64
	Photo        []byte
}

type Match struct {
	Date          time.Time
	VisitorID     uint64
	LocalID       uint64
	LocalScore    byte
	VisitorScore  byte
	LocalShoots   uint16
	VisitorShoots uint16
}

type HistoricalData struct {
	Year          uint8
	LeagueResults []Match
}

type Team struct {
	ID             uint64
	Name           string
	Shield         []byte
	Players        []Player
	HistoricalData []HistoricalData
}

func NewTeamFactory() TeamFlyweightFactory {
	return TeamFlyweightFactory{
		createdTeams: make(map[int]*Team, 0),
	}
}

func getTeamFactory(team int) Team {
	switch team {
	case TEAM_A:
		return Team{
			ID:   1,
			Name: "TEAM_A",
		}
	case TEAM_B:
		return Team{
			ID:   2,
			Name: "TEAM_B",
		}
	default:
		return Team{}
	}
}

type TeamFlyweightFactory struct {
	createdTeams map[int]*Team
}

func (t *TeamFlyweightFactory) GetTeam(teamID int) *Team {
	if t.createdTeams[teamID] != nil {
		return t.createdTeams[teamID]
	}
	team := getTeamFactory(teamID)
	t.createdTeams[teamID] = &team
	return t.createdTeams[teamID]
}

func (t *TeamFlyweightFactory) GetNumberOfObjects() int {
	return len(t.createdTeams)
}
