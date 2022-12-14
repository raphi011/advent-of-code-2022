package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"path"
	"strconv"
	"strings"
)

type dir struct {
	root  bool
	files map[string]*file
	dirs  map[string]*dir
}

type file struct {
	name string
	size int
}

func main() {
	exercise := flag.Int("e", 1, "exercise 1 or 2")
	flag.Parse()

	reader := bufio.NewReader(os.Stdin)

	cwd := "/"
	root := &dir{
		root:  true,
		files: make(map[string]*file),
		dirs:  make(map[string]*dir),
	}

	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}

		input := strings.Split(strings.TrimSpace(line), " ")

		if input[0] == "$" {
			input = input[1:]
			cmd := input[0]

			switch cmd {
			case "cd":
				if input[1] == ".." {
					if cwd != "/" {
						cwd = cwd[:strings.LastIndex(cwd, "/")]

						if cwd == "" {
							cwd = "/"
						}
					}
				} else if input[1] == "/" {
					cwd = "/"
				} else {
					if cwd == "/" {
						root.addDir(input[1])
					} else {
						root.dir(cwd[1:]).addDir(input[1])
					}

					cwd = path.Join(cwd, input[1])
				}
			case "ls":
			default:
				panic(fmt.Sprintf("unknown command: %q", cmd))
			}
		} else {
			currentDirectory := root

			if cwd != "/" {
				currentDirectory = root.dir(cwd[1:])
			}

			// output of `ls` command
			if input[0] == "dir" {
				currentDirectory.addDir(input[1])
			} else {
				size, err := strconv.Atoi(input[0])
				if err != nil {
					panic("filesize " + input[0] + " is not a number")
				}
				name := input[1]

				currentDirectory.addFile(name, size)
			}
		}
	}

	result := 0

	used := root.size()
	total := 70000000
	remaining := total - used
	needed := 30000000 - remaining

	root.walk(func(d *dir) {
		size := d.size()

		if *exercise == 1 {
			if !d.root && size <= 100000 {
				result += size
			}
		} else {
			if size >= needed && (result == 0 || size < result) {
				result = size
			}
		}
	})

	fmt.Printf("%d\n", result)
}

func (d *dir) size() int {
	s := 0

	for _, f := range d.files {
		s += f.size
	}

	for _, d := range d.dirs {
		s += d.size()
	}

	return s
}

func (d *dir) dir(relativePath string) *dir {
	i := strings.Index(relativePath, "/")
	if i == -1 {
		return d.dirs[relativePath]
	}

	childDir := d.dirs[relativePath[:i]]

	if childDir == nil {
		return nil
	}

	return childDir.dir(relativePath[i+1:])
}

func (d *dir) addDir(name string) {
	if d.dirs[name] == nil {
		d.dirs[name] = &dir{
			files: make(map[string]*file),
			dirs:  make(map[string]*dir),
			root:  false,
		}
	}
}

func (d *dir) walk(f func(d *dir)) {
	f(d)

	for _, d := range d.dirs {
		d.walk(f)
	}
}

func (d *dir) addFile(name string, size int) {
	d.files[name] = &file{
		name: name,
		size: size,
	}
}
