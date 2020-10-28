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
	point  int
}

// String returns string representation of a team.
func (t Team) String() string {
	return fmt.Sprintf("%-30s | %2d | %2d | %2d | %2d | %2d", t.name, t.played, t.win, t.draw, t.loss, t.point)
}

// Result represents a tournament standings.
type Result []Team

// Len returns length of a tournament result.
func (r Result) Len() int {
	return len(r)
}

// Less compares two elements of a tournament result.
func (r Result) Less(i, j int) bool {
	if r[i].point == r[j].point {
		return strings.Compare(r[i].name, r[j].name) < 0
	}
	return r[i].point > r[j].point
}

// Swap swaps two elements position in a tournament result.
func (r Result) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

// Teams represents all team data before sorted.
type Teams map[string]Team

// Tally read teams matches from a io.Reader and prints the result to the given io.Writer.
func Tally(r io.Reader, w io.Writer) error {
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
			return err
		}
		processMatch(teams, match)
	}
	result := generateResult(teams)
	return printResult(w, result)
}

// generateResult generates tournament result from slice of teams.
func generateResult(teams Teams) Result {
	var result Result
	for _, team := range teams {
		result = append(result, team)
	}
	sort.Sort(result)
	return result
}

// checkMatch checks if given match record is valid.
func checkMatch(match []string) error {
	if len(match) != 3 {
		return fmt.Errorf("wrong match format")
	}
	if match[2] != "win" && match[2] != "loss" && match[2] != "draw" {
		return fmt.Errorf("invalid result, must be win or loss")
	}
	return nil
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
		home.point += 3
		away.loss++
	case "loss":
		home.loss++
		away.win++
		away.point += 3
	case "draw":
		home.draw++
		home.point++
		away.draw++
		away.point++
	}
	teams[match[0]], teams[match[1]] = home, away
}

// printResult prints tournament result in table format to given writer.
func printResult(w io.Writer, result Result) error {
	if _, err := w.Write([]byte("Team                           | MP |  W |  D |  L |  P\n")); err != nil {
		return err
	}
	for _, team := range result {
		if _, err := w.Write([]byte(fmt.Sprintf("%s\n", team))); err != nil {
			return err
		}
	}
	return nil
}
