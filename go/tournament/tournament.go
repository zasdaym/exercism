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
	points  int
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
	teams, err := readMatches(r)
	if err != nil {
		return err
	}
	result := generateResult(teams)
	return printResult(w, result)
}

// readMatches reads match records from a io.Reader and returns teams stats from it.
func readMatches(r io.Reader) (Teams, error) {
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
		if err := checkMatch(match); err != nil {
			return nil, err
		}
		processMatch(teams, match)
	}
	return teams, nil
}

// checkMatch checks if given match record is valid.
func checkMatch(match []string) error {
	if len(match) != 3 {
		return fmt.Errorf("wrong match format")
	}
	switch match[2] {
	case "win", "loss", "draw":
		return nil
	}
	return fmt.Errorf("match result must be win, loss, or draw")
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

// processMatch processes a match record.
func processMatch(teams Teams, match []string) {
	home, away := teams[match[0]], teams[match[1]]
	home.name = match[0]
	away.name = match[1]
	home.played++
	away.played++
	switch match[2] {
	case "win":
		home.win++
		home.points += 3
		away.loss++
	case "loss":
		home.loss++
		away.win++
		away.points += 3
	case "draw":
		home.draw++
		home.points++
		away.draw++
		away.points++
	}
	teams[match[0]], teams[match[1]] = home, away
}

// printResult prints tournament result in table format to given writer.
func printResult(w io.Writer, result Result) error {
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
