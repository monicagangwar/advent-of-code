package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"

	"github.com/monicagangwar/advent-of-code-2022/input"
)

type dir struct {
	name      string
	files     []int64
	totalSize int64
	subdirs   []*dir
	parentDir *dir
}

func (d *dir) getPath() string {
	if d == nil {
		return ""
	}
	if d.name == "/" {
		return ""
	}
	return d.parentDir.getPath() + "/" + d.name
}

func dirFound(dirName string, curDir *dir, dirDiscovered map[string]*dir) *dir {
	if curDir == nil {
		fmt.Printf("\nfirst dir found : %s", dirName)
	} else {
		fmt.Printf("\ndir found : %s in curDir: %s", dirName, curDir.name)
	}

	dirDiscoveredAlready, found := dirDiscovered[curDir.getPath()+"/"+dirName]
	if found {
		fmt.Printf(" dir was already discovered nothing to do ")
		return dirDiscoveredAlready
	}
	dirFoundCurrently := &dir{
		name:      dirName,
		parentDir: curDir,
	}
	dirDiscovered[dirFoundCurrently.getPath()] = dirFoundCurrently
	if curDir != nil {
		if curDir.subdirs == nil {
			curDir.subdirs = make([]*dir, 0)
		}
		curDir.subdirs = append(curDir.subdirs, dirFoundCurrently)
	}

	return dirFoundCurrently
}

func computeTotalSize(root *dir) int64 {
	fmt.Printf("computing for dir %s\n", root.name)
	if root == nil {
		return 0
	}
	totalSize := int64(0)
	if root.subdirs != nil {
		for _, subdir := range root.subdirs {
			totalSize += computeTotalSize(subdir)
		}
	}
	if root.files != nil {
		for _, file := range root.files {
			totalSize += file
		}
	}
	root.totalSize = totalSize
	return totalSize
}

func findSumIfTotalSizeAtMost(root *dir, atMost int64) int64 {
	if root == nil {
		return 0
	}
	sum := int64(0)
	if root.totalSize <= atMost {
		sum += root.totalSize
	}
	if root.subdirs != nil {
		for _, subdir := range root.subdirs {
			sum += findSumIfTotalSizeAtMost(subdir, atMost)
		}
	}
	return sum
}

func findSmallestDirToFreeUp(root *dir, totalSpaceToFreeUp int64) int64 {
	if root == nil {
		return -1
	}
	spaceToDelete := root.totalSize
	if root.subdirs != nil {
		for _, subdir := range root.subdirs {
			spaceToDeleteFromSubdir := findSmallestDirToFreeUp(subdir, totalSpaceToFreeUp)
			if spaceToDeleteFromSubdir != -1 && spaceToDeleteFromSubdir >= totalSpaceToFreeUp && spaceToDeleteFromSubdir < spaceToDelete {
				spaceToDelete = spaceToDeleteFromSubdir
			}
		}
	}
	return spaceToDelete
}

func main() {
	_, currentFilePath, _, _ := runtime.Caller(0)
	content := input.ReadInput(currentFilePath)
	lines := strings.Split(string(content), "\n")
	lsCommandExecuted := false
	var rootDir *dir
	var curDir *dir
	dirDiscovered := make(map[string]*dir)
	filesDiscovered := make(map[string]struct{})
	for _, line := range lines {
		if line[0] == '$' {
			lsCommandExecuted = false
			switch line {
			case "$ cd /":
				fmt.Println("changing directory to /")
				rootDir = dirFound("/", curDir, dirDiscovered)
				curDir = rootDir
				break
			case "$ cd ..":
				fmt.Printf("\nchanging directory to %s", curDir.parentDir.name)
				curDir = curDir.parentDir
				break
			case "$ ls":
				lsCommandExecuted = true
				break
			default:
				// default is $ cd <dir>
				dirName := strings.ReplaceAll(line, "$ cd ", "")
				fmt.Printf("\nchanging directory to %s", dirName)
				dirCreated := dirFound(dirName, curDir, dirDiscovered)
				curDir = dirCreated
			}
		} else if lsCommandExecuted {
			if strings.HasPrefix(line, "dir") {
				dirName := strings.ReplaceAll(line, "dir ", "")
				dirFound(dirName, curDir, dirDiscovered)
			} else {
				fileDetails := strings.Split(line, " ")
				fileSize, _ := strconv.ParseInt(fileDetails[0], 10, 64)
				fileName := fileDetails[1]

				fileDirSha := curDir.getPath() + "/" + fileName
				if _, found := filesDiscovered[fileDirSha]; !found {
					filesDiscovered[fileDirSha] = struct{}{}
					if curDir.files == nil {
						curDir.files = make([]int64, 0)
					}
					curDir.files = append(curDir.files, fileSize)
				}
			}
		}
	}

	fmt.Println(computeTotalSize(rootDir))
	fmt.Println(findSumIfTotalSizeAtMost(rootDir, int64(100000)))
	totalSpaceNeeded := 30000000 - (70000000 - rootDir.totalSize)
	fmt.Println(totalSpaceNeeded)
	fmt.Println(findSmallestDirToFreeUp(rootDir, totalSpaceNeeded))
}
