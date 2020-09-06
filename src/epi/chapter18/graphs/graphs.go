package graphs

type Coordinate struct {
	x, y int
}

func (c Coordinate) Equals(other Coordinate) bool {
	if c.x != other.x || c.y != other.y {
		return false
	}
	return true
}

type Color int

const (
	WHITE Color = 0
	BLACK Color = 1
)

func SearchMaze(maze [][]Color, s, e Coordinate) []Coordinate {
	visited := map[Coordinate]bool{}
	result := []Coordinate{}
	DFS(maze, s, e, visited, &result)
	return result
}

var direction = []Coordinate{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

func isValid(maze [][]Color, coor Coordinate) bool {
	return 0 <= coor.x &&
		coor.x < len(maze) &&
		0 <= coor.y &&
		coor.y < len(maze[coor.x]) &&
		maze[coor.x][coor.y] == WHITE
}

func DFS(maze [][]Color, s, e Coordinate, visited map[Coordinate]bool, path *[]Coordinate) bool {
	if !isValid(maze, s) {
		return false
	}
	*path = append(*path, s)
	visited[s] = true
	if s.Equals(e) {
		return true
	}
	for _, dir := range direction {
		nextPos := Coordinate{s.x + dir.x, s.y + dir.y}
		if !visited[nextPos] && DFS(maze, nextPos, e, visited, path) {
			return true
		}
	}
	*path = (*path)[0 : len(*path)-1]
	return false
}
