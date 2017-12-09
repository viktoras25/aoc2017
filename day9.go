package main

import "fmt"
import "io"
import "os"

const (
	T_GROUP_BEGIN   = "{"
	T_GROUP_END     = "}"
	T_GARBAGE_BEGIN = "<"
	T_GARBAGE_END   = ">"
	T_BANG          = "!"
)

func main() {
	f, _ := os.Open("input9.txt")

	buf := make([]byte, 1)

	depth := 0
	score := 0
	garbage := false
	removed := 0

	for {
		if _, err := f.Read(buf); err != io.EOF {
			switch string(buf) {
			case T_GROUP_BEGIN:
				if !garbage {
					depth++
				} else {
					removed++
				}

			case T_GROUP_END:
				if !garbage {
					score += depth
					depth--
				} else {
					removed++
				}

			case T_GARBAGE_BEGIN:
				if !garbage {
					garbage = true
				} else {
					removed++
				}

			case T_GARBAGE_END:
				if garbage {
					garbage = false
				} else {
					removed++
				}

			case T_BANG:
				f.Read(buf)

			default:
				if garbage {
					removed++
				}
			}

			continue
		}

		break
	}

	fmt.Println("Part 1", score)
	fmt.Println("Part 2", removed)
}
