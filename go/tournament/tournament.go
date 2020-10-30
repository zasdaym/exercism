package tournament

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

// Team represents a team.
type Team struct {
	name   string
	played int
	win    int
	loss   int
	draw   int
	points int
}

// String returns string representation of a team.
func (t Team) String() string {
	return fmt.Sprintf("%-30s | %2d | %2d | %2d | %2d | %2d", t.name, t.played, t.win, t.draw, t.loss, t.points)
}

// Result represents a tournament standings.
type Result []Team

// Teams represents all team data before sorted.
type Teams map[string]Team

// Tally read teams matches from a io.Reader and prints the result to the given io.Writer.
func Tally(r io.Reader, w io.Writer) error {
	teams, err := readMatchDataFrom(r)
	if err != nil {
		return err
	}
	result := generateResult(teams)
	return writeResultTo(w, result)
}

// readMatchDataFrom reads match records from an io.Reader and returns teams stats from it.
func readMatchDataFrom(r io.Reader) (Teams, error) {
	teams := make(Teams)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		match := strings.Split(line, ";")
		if match[0] == "" {
			continue
		}
		if match[0][0] == '#' {
			continue
		}
		if len(match) != 3 {
			return nil, fmt.Errorf("wrong match format")
		}
		a, b := teams[match[0]], teams[match[1]]
		a.name = match[0]
		b.name = match[1]
		a.played++
		b.played++
		switch match[2] {
		case "win":
			a.win++
			a.points += 3
			b.loss++
		case "loss":
			a.loss++
			b.win++
			b.points += 3
		case "draw":
			a.draw++
			a.points++
			b.draw++
			b.points++
		default:
			return nil, fmt.Errorf("match result must be win, lose, or draw")
		}
		teams[match[0]], teams[match[1]] = a, b
	}
	return teams, nil
}

// generateResult generates tournament result from slice of teams.
func generateResult(teams Teams) Result {
	var result Result
	for _, team := range teams {
		result = append(result, team)
	}
	sort.Slice(result, func(i, j int) bool {
		if result[i].points == result[j].points {
			return result[i].name < result[j].name
		}
		return result[i].points > result[j].points
	})
	return result
}

// writeResultTo writes tournament result table to given writer.
func writeResultTo(w io.Writer, result Result) error {
	if _, err := fmt.Fprintf(w, "Team                           | MP |  W |  D |  L |  P\n"); err != nil {
		return err
	}
	for _, team := range result {
		if _, err := fmt.Fprintf(w, fmt.Sprintf("%s\n", team)); err != nil {
			return err
		}
	}
	return nil
}
