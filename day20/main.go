package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Tile struct {
	id   int
	data []string
}

func (tile *Tile) Border(dir int) string {
	switch dir {
	case 0:
		return tile.data[0]
	case 1:
		data := make([]byte, 10)
		for i := 0; i < 10; i++ {
			data[i] = tile.data[i][9]
		}
		return string(data)
	case 2:
		return reverse(tile.data[9])
	case 3:
		data := make([]byte, 10)
		for i := 0; i < 10; i++ {
			data[i] = tile.data[9-i][0]
		}
		return string(data)
	default:
		panic("invalid dir")
	}
}

func main() {
	tiles := make(map[int]Tile)

	lines := readLines("input.txt")
	idRegex := regexp.MustCompile(`Tile (\d+):`)
	for index := 0; index < len(lines); index += 12 {
		matches := idRegex.FindStringSubmatch(lines[index])
		id := toInt(matches[1])
		tiles[id] = Tile{id, lines[index+1 : index+11]}
	}

	borders := make(map[string][]int)
	addBorder := func(id int, border string) {
		borders[border] = append(borders[border], id)
	}
	for _, tile := range tiles {
		addBorder(tile.id, tile.Border(0))
		addBorder(tile.id, tile.Border(1))
		addBorder(tile.id, tile.Border(2))
		addBorder(tile.id, tile.Border(3))
		addBorder(tile.id, reverse(tile.Border(0)))
		addBorder(tile.id, reverse(tile.Border(1)))
		addBorder(tile.id, reverse(tile.Border(2)))
		addBorder(tile.id, reverse(tile.Border(3)))
	}

	var start Tile
	var cornerProduct int = 1

	for _, tile := range tiles {
		count := 0
		for i := 0; i < 4; i++ {
			count += len(borders[tile.Border(i)])
		}
		if count == 6 {
			start = tile
			cornerProduct *= tile.id
		}
	}

	{
		fmt.Println("--- Part One ---")
		fmt.Println(cornerProduct)
	}

	for i := 0; i < 4; i++ {
		if len(borders[start.Border(i)])+len(borders[start.Border((i+3)%4)]) == 2 {
			start.data = rotate(start.data, (4-i)%4)
			break
		}
	}

	tiledim := int(math.Sqrt(float64(len(tiles))))
	arrangement := make([]Tile, tiledim*tiledim)
	arrangement[0*tiledim+0] = start

	for y := 0; y < tiledim; y++ {
		if y != 0 {
			top := arrangement[(y-1)*tiledim+0]
			border := reverse(top.Border(2))

			var tile Tile
			for _, id := range borders[border] {
				if id != top.id {
					tile = tiles[id]
					break
				}
			}

			for i := 0; i < 4; i++ {
				if tile.Border(i) == border || reverse(tile.Border(i)) == border {
					tile.data = rotate(tile.data, (4-i)%4)
					break
				}
			}
			if tile.Border(0) != border {
				tile.data = flipHorizontally(tile.data)
			}

			arrangement[y*tiledim+0] = tile
		}

		for x := 1; x < tiledim; x++ {
			left := arrangement[y*tiledim+(x-1)]
			border := reverse(left.Border(1))

			var tile Tile
			for _, id := range borders[border] {
				if id != left.id {
					tile = tiles[id]
					break
				}
			}

			for i := 0; i < 4; i++ {
				if tile.Border((i+3)%4) == border || reverse(tile.Border((i+3)%4)) == border {
					tile.data = rotate(tile.data, (4-i)%4)
					break
				}
			}
			if tile.Border(3) != border {
				tile.data = flipVertically(tile.data)
			}

			arrangement[y*tiledim+x] = tile
		}
	}

	if false {
		// Print arrangement.
		for y := 0; y < tiledim; y++ {
			for i := 0; i < 10; i++ {
				for x := 0; x < tiledim; x++ {
					if arrangement[y*tiledim+x].id != 0 {
						fmt.Print(" ", arrangement[y*tiledim+x].data[i])
					}
				}
				fmt.Println()
			}
			fmt.Println()
		}
		fmt.Println()
	}

	imgdim := 8 * tiledim
	rawimage := make([][]byte, imgdim)

	for y := 0; y < tiledim; y++ {
		for i := 0; i < 8; i++ {
			rawimage[8*y+i] = make([]byte, imgdim)
			for x := 0; x < tiledim; x++ {
				for j := 0; j < 8; j++ {
					rawimage[(8*y + i)][(8*x + j)] = arrangement[y*tiledim+x].data[1+i][1+j]
				}
			}
		}
	}

	image := make([]string, imgdim)
	for i := range image {
		image[i] = string(rawimage[i])
	}

	if false {
		// Print image.
		for y := 0; y < imgdim; y++ {
			fmt.Println(image[y])
		}
	}

	kernel := []string{
		"..................#.",
		"#....##....##....###",
		".#..#..#..#..#..#...",
	}
	kheight, kwidth := len(kernel), len(kernel[0])

	findSeaMonsters := func(image []string) (count int) {
		for y := 0; y < len(image)-kheight; y++ {
			for x := 0; x < len(image[y])-kwidth; x++ {
				ok := true
				for dy := 0; ok && dy < kheight; dy++ {
					for dx := 0; ok && dx < kwidth; dx++ {
						if kernel[dy][dx] == '#' && image[(y + dy)][(x+dx)] != '#' {
							ok = false
						}
					}
				}
				if ok {
					count++
				}
			}
		}
		return
	}

	var monsters int
	for i := 0; i < 4; i++ {
		monsters += findSeaMonsters(rotate(image, i))
		monsters += findSeaMonsters(flipHorizontally(rotate(image, i)))
	}

	{
		fmt.Println("--- Part Two ---")
		fmt.Println(strings.Count(strings.Join(image, ""), "#") - monsters*strings.Count(strings.Join(kernel, ""), "#"))
	}
}

func rotate(data []string, dir int) []string {
	switch dir {
	case 0:
		return data
	case 1:
		result := make([]string, len(data[0]))
		for i := 0; i < len(result); i++ {
			temp := make([]byte, len(data))
			for j := 0; j < len(temp); j++ {
				temp[j] = data[len(data)-1-j][i]
			}
			result[i] = string(temp)
		}
		return result
	case 2:
		result := make([]string, len(data))
		for i := 0; i < len(result); i++ {
			result[i] = reverse(data[len(data)-1-i])
		}
		return result
	case 3:
		result := make([]string, len(data[0]))
		for i := 0; i < len(result); i++ {
			temp := make([]byte, len(data))
			for j := 0; j < len(temp); j++ {
				temp[j] = data[j][len(data[0])-1-i]
			}
			result[i] = string(temp)
		}
		return result
	default:
		panic("invalid dir")
	}
}

func flipHorizontally(data []string) []string {
	result := make([]string, len(data))
	for i := range result {
		result[i] = reverse(data[i])
	}
	return result
}

func flipVertically(data []string) []string {
	result := make([]string, len(data))
	for i := range result {
		result[i] = data[len(data)-1-i]
	}
	return result
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func readLines(filename string) []string {
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func toInt(s string) int {
	result, err := strconv.Atoi(s)
	check(err)
	return result
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
