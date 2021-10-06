// Package tournament provides Tournament solution.
package tournament

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

// Tally reads football competition data from r and writes to w in table format.
func Tally(r io.Reader, w io.Writer) error {
	stats, err := readMatchesFrom(r)
	if err != nil {
		return err
	}

	standings := newStandings(stats)
	return writeStandingsTo(w, standings)
}

// team represents a team in a tournament.
type team struct {
	name                            string
	played, win, draw, lose, points int
}

// readMatchesFrom reads matches from r and returns the final team statistics.
func readMatchesFrom(r io.Reader) (map[string]team, error) {
	teams := make(map[string]team)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		match := strings.Split(line, ";")

		if match[0] == "" || match[0][0] == '#' {
			continue
		}

		if len(match) != 3 {
			return nil, fmt.Errorf("invalid line: want 3 fields, got %d field(s)", len(match))
		}

		home, away, result := match[0], match[1], match[2]
		homeTeam, awayTeam := teams[home], teams[away]
		homeTeam.name = home
		awayTeam.name = away
		homeTeam.played++
		awayTeam.played++
		switch result {
		case "win":
			homeTeam.win++
			homeTeam.points += 3
			awayTeam.lose++
		case "draw":
			homeTeam.draw++
			homeTeam.points++
			awayTeam.draw++
			awayTeam.points++
		case "loss":
			homeTeam.lose++
			awayTeam.win++
			awayTeam.points += 3
		default:
			return nil, fmt.Errorf("invalid line: %s is not a valid match result", result)
		}
		teams[home] = homeTeam
		teams[away] = awayTeam
	}
	return teams, nil
}

// standings represents a tournament standings with teams sorted by points in descending order.
type standings []team

// newStandings creates new tournament standings from team stats.
func newStandings(teams map[string]team) standings {
	var result standings
	for _, t := range teams {
		result = append(result, t)
	}
	sort.Sort(result)
	return result
}

func (s standings) Len() int {
	return len(s)
}

func (s standings) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s standings) Less(i, j int) bool {
	if s[i].points == s[j].points {
		return s[i].name < s[j].name
	}
	return s[i].points > s[j].points
}

const standingsHeader = "Team                           | MP |  W |  D |  L |  P\n"

// writeStandingsTo writes given standings into w in table format.
func writeStandingsTo(w io.Writer, standings []team) error {
	if _, err := fmt.Fprint(w, standingsHeader); err != nil {
		return err
	}

	for _, t := range standings {
		if _, err := fmt.Fprintf(w, "%-30s | %2d | %2d | %2d | %2d | %2d\n", t.name, t.played, t.win, t.draw, t.lose, t.points); err != nil {
			return err
		}
	}
	return nil
}
