package graphs

import (
	"encoding/json"
	"epi/test"
	"errors"
	"strconv"
	"testing"
)

func PathElementIsFeasible(maze [][]Color, prev, cur Coordinate) bool {
	if !(0 <= cur.x && cur.x < len(maze) && 0 <= cur.y &&
		cur.y < len(maze[cur.x]) && maze[cur.x][cur.y] == 0) {
		return false
	}

	return cur.x == prev.x+1 && cur.y == prev.y ||
		cur.x == prev.x-1 && cur.y == prev.y ||
		cur.x == prev.x && cur.y == prev.y+1 ||
		cur.x == prev.x && cur.y == prev.y-1
}

func SearchMazeWrapper(maze [][]Color, s, e Coordinate) ([]Coordinate, bool, error) {
	path := SearchMaze(maze, s, e)
	if len(path) <= 0 {
		return path, s.Equals(e), nil
	}
	if !path[0].Equals(s) || !path[len(path)-1].Equals(e) {
		return path, false, errors.New("path doesn't lay between start and end points")
	}

	for i := 1; i < len(path); i++ {
		if !PathElementIsFeasible(maze, path[i-1], path[i]) {
			return path, false, errors.New("path contains invalid segments")
		}
	}
	return path, true, nil
}

func TestSearchMaze(t *testing.T) {
	testData := test.GetTestData("search_maze.tsv")[1:]
	type args struct {
		maze [][]Color
		s, e Coordinate
	}
	type test struct {
		name string
		args args
		want bool
	}
	var tests []test
	for i, datum := range testData {
		var input0 [][]Color
		var input1, input2 []int
		var output bool
		json.Unmarshal([]byte(datum[0]), &input0)
		json.Unmarshal([]byte(datum[1]), &input1)
		json.Unmarshal([]byte(datum[2]), &input2)
		json.Unmarshal([]byte(datum[3]), &output)

		tests = append(tests, test{
			name: strconv.Itoa(i),
			args: args{
				maze: input0,
				s:    Coordinate{input1[0], input1[1]},
				e:    Coordinate{input2[0], input2[1]},
			},
			want: output,
		})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if path, got, err := SearchMazeWrapper(tt.args.maze, tt.args.s, tt.args.e); err != nil || got != tt.want {
				t.Errorf("SearchMaze() = %v, got = %v, want = %v, err %v", path, got, tt.want, err)
			}
		})
	}
}
