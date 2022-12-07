package main

import (
	"advent/util"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	Input            string
	changeDirRegex   = regexp.MustCompile(`\$ cd (.*)`)
	lsRegex          = regexp.MustCompile(`\$ ls`)
	fileRegex        = regexp.MustCompile(`(\d+) (.*)`)
	dirRegex         = regexp.MustCompile(`dir (.*)`)
	rootDirectory    *Directory
	total            = 70000000
	desiredFreeSpace = 30000000
)

func main() {
	Input = util.ReadFile("./input.txt")
	fmt.Println(part1())
	fmt.Println(part2())
}

func part1() (sum int) {
	parseInput()
	for _, d := range rootDirectory.findDirsWithMaxSize(100000) {
		sum += d.Size
	}
	return
}

func part2() (sum int) {
	parseInput()
	smallestDir := rootDirectory

	for _, d := range rootDirectory.findDirsWithMinSize(desiredFreeSpace - (total - rootDirectory.Size)) {
		if d.Size < smallestDir.Size {
			smallestDir = d
		}
	}
	return smallestDir.Size
}

func parseInput() {
	rootDirectory = &Directory{
		Name:        "/",
		Directories: make(map[string]*Directory),
		Files:       make(map[string]*File),
		Parent:      nil,
	}
	currentDir := rootDirectory

	for _, s := range strings.Split(Input, "\n") {
		if s == "$ cd /" {
			continue
		}

		switch true {
		case changeDirRegex.MatchString(s): // cd into subdir
			matches := changeDirRegex.FindStringSubmatch(s)
			if matches[1] == ".." {
				currentDir = currentDir.Parent
			} else {
				currentDir = currentDir.Directories[matches[1]]
			}
		case lsRegex.MatchString(s): // idk why i used regex for this but w/e
			continue
		case fileRegex.MatchString(s): // add file to current directory
			matches := fileRegex.FindStringSubmatch(s)
			size, _ := strconv.Atoi(matches[1])
			name := matches[2]

			currentDir.Files[name] = &File{
				Name: name,
				Size: size,
			}
		case dirRegex.MatchString(s): // add directory to current directory
			var (
				matches = dirRegex.FindStringSubmatch(s)
				dirName = matches[1]
			)
			currentDir.Directories[dirName] = &Directory{
				Name:        dirName,
				Directories: make(map[string]*Directory),
				Files:       make(map[string]*File),
				Parent:      currentDir,
			}
		}
	}
	rootDirectory.calculateSize()
}

type Directory struct {
	Name        string
	Directories map[string]*Directory
	Files       map[string]*File
	Parent      *Directory
	Size        int
}

type File struct {
	Size int
	Name string
}

func (d *Directory) findDirsWithMaxSize(maxSize int) (directories []*Directory) {
	for _, dir := range d.Directories {
		directories = append(directories, dir.findDirsWithMaxSize(maxSize)...)
	}

	if d.Size <= maxSize {
		directories = append(directories, d)
	}
	return
}

func (d *Directory) findDirsWithMinSize(min int) (directories []*Directory) {
	for _, dir := range d.Directories {
		directories = append(directories, dir.findDirsWithMinSize(min)...)
	}

	if d.Size >= min {
		directories = append(directories, d)
	}
	return
}

func (d *Directory) calculateSize() {
	for _, f := range d.Files {
		d.Size += f.Size
	}

	for _, dir := range d.Directories {
		dir.calculateSize()
		d.Size += dir.Size
	}
}
