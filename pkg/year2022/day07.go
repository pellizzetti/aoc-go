package year2022

import (
	"strconv"
	"strings"
)

type Day07 struct{}

// type Folder struct {
// 	parent  *Folder
// 	size    int
// 	dirName string
// }

type File struct {
	Name string
	Size int
}

type Dir struct {
	Name   string
	Parent *Dir
	// Children map[string]Dir
	// Files     []File
	Size int
	Root bool
}

func (p Day07) PartA(lines []string) any {
	directories := make(map[string]Dir)
	var path []string
	for _, line := range lines {
		if line == "" {
			break
		}
		input := strings.SplitN(line, " ", 3)
		if input[0] == "$" && input[1] == "cd" {
			dir := input[2]
			if dir == "/" {
				directories[dir] = Dir{Name: dir, Parent: nil, Root: true}
			}
			if dir == ".." {
				path[len(path)-1] = ""
				path = path[:len(path)-1]
				continue
			}
			path = append(path, dir)

		} else if input[0] == "dir" {
			parentDirName := path[len(path)-1]
			parentDir := directories[parentDirName]
			dirName := input[1]
			if _, ok := directories[dirName]; !ok {
				dirPath, _, _ := getPathString(path)
				if dirPath == "/" {
					dirPath = ""
				}
				dirPath = dirPath + "/" + dirName
				directories[dirPath] = Dir{Name: dirName, Parent: &parentDir}
			}
		} else {
			dirPath, parentPath, dirName := getPathString(path)
			size, _ := strconv.Atoi(strings.Split(line, " ")[0])
			parentDir := directories[parentPath]
			dir := directories[dirPath]
			dirSize := dir.Size + size
			dir = Dir{Name: dirName, Parent: &parentDir, Size: dirSize, Root: dir.Root}
			directories[dirPath] = dir
			if !dir.Root {
				parentDir.Size += size
				directories[parentPath] = parentDir
			}
		}
	}
	// fmt.Println(directories)
	sum := 0
	for _, v := range directories {
		if v.Parent.Name == "/" && v.Size <= 100000 {
			sum += v.Size
		}
	}
	return sum
}

// func (p Day07) PartA(lines []string) any {
// 	directories := make(map[string]Dir)
// 	var parentDirName string
// 	var dirName string
// 	for _, line := range lines {
// 		if line == "" {
// 			break
// 		}
// 		input := strings.SplitN(line, " ", 3)
// 		if input[0] == "$" {
// 			if input[1] == "cd" {
// 				dir := input[2]
// 				if dir == ".." {
// 					dirName = parentDirName
// 					parentDirName = directories[dirName].Parent.Name
// 					continue
// 				} else if dir == "/" {
// 					directories[dir] = Dir{Name: dir, Parent: nil, Root: true}
// 				}
// 				parentDirName = dirName
// 				dirName = dir
// 			}
// 		} else if input[0] == "dir" {
// 			parentDir := directories[dirName]
// 			childDirName := input[1]
// 			if _, ok := directories[childDirName]; !ok {
// 				directories[childDirName] = Dir{Name: childDirName, Parent: &parentDir}
// 			}
// 		} else {
// 			size, _ := strconv.Atoi(strings.Split(line, " ")[0])
// 			parentDir := directories[parentDirName]
// 			dir := directories[dirName]
// 			dirSize := dir.Size + size
// 			dir = Dir{Name: dirName, Parent: &parentDir, Size: dirSize, Root: dir.Root}
// 			directories[dirName] = dir
// 			if !dir.Root {
// 				parentDir.Size += size
// 				directories[parentDirName] = parentDir
// 			}
// 		}
// 	}
// 	sum := 0
// 	for _, v := range directories {
// 		if v.Parent.Name == "/" && v.Size <= 100000 {
// 			sum += v.Size
// 		}
// 	}
// 	return sum
// }

func (p Day07) PartB(lines []string) any {
	return "implement_me"
}

func getPathString(path []string) (string, string, string) {
	dirPath := strings.Join(path, "/")
	parentPath := strings.Join(path[:len(path)-1], "/")
	dirName := path[len(path)-1]
	if len(path) > 1 {
		dirPath = dirPath[1:]
	}
	if len(path) > 2 {
		parentPath = parentPath[1:]
	}
	return dirPath, parentPath, dirName
}
