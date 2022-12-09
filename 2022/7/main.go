package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

type Folder struct {
	Name      string
	Parent    *Folder
	SubFolder []*Folder
	Files     []*File
}

type File struct {
	Name string
	Size int
}

func (folder *Folder) FolderSizeMap() map[string]int {
	sizeMap := make(map[string]int)
	for _, subFolder := range folder.SubFolder {
		for name, size := range subFolder.FolderSizeMap() {
			sizeMap[name] = size
		}
	}
	sizeMap[folder.Name] = folder.Size()
	return sizeMap
}

func (folder *Folder) FolderLessThanSize(size int) []*Folder {
	var folders []*Folder
	for _, sub := range folder.SubFolder {
		folders = append(folders, sub.FolderLessThanSize(size)...)
		if sub.Size() <= size {
			folders = append(folders, sub)
		}
	}
	return folders
}

func (folder *Folder) Size() int {
	var size int
	for _, file := range folder.Files {
		size += file.Size
	}
	for _, folder := range folder.SubFolder {
		size += folder.Size()
	}
	return size
}

func (folder *Folder) FindFolder(path string) *Folder {
	if path == "/" {
		return folder
	}
	if strings.HasPrefix(path, "/") {
		path = strings.TrimPrefix(path, "/")
	}
	if strings.HasSuffix(path, "/") {
		path = strings.TrimSuffix(path, "/")
	}
	if strings.Contains(path, "/") {
		tmp := strings.Split(path, "/")
		for _, sub := range folder.SubFolder {
			if sub.Name == tmp[0] {
				return sub.FindFolder(strings.Join(tmp[1:], "/"))
			}
		}
	} else {
		for _, sub := range folder.SubFolder {
			if sub.Name == path {
				return sub
			}
		}
	}
	return nil
}

func Parse(terminal string) *Folder {
	root := &Folder{}
	current := root
	for _, line := range strings.Split(terminal, "\r\n") {
		if strings.HasPrefix(line, "$ cd ") {
			path := strings.TrimPrefix(line, "$ cd ")
			if path == "/" {
				current = root
			} else if path == ".." {
				current = current.Parent
			} else {
				current = current.FindFolder(path)
			}
		} else if strings.HasPrefix(line, "$ ls") {
			//Do nothing
		} else {
			if strings.HasPrefix(line, "dir ") {
				current.SubFolder = append(current.SubFolder, &Folder{
					Name:   strings.TrimPrefix(line, "dir "),
					Parent: current,
				})
			} else {
				tmp := strings.Split(line, " ")
				size, _ := strconv.Atoi(tmp[0])
				current.Files = append(current.Files, &File{
					Name: tmp[1],
					Size: size,
				})
			}
		}
	}
	return root
}

func main() {
	input, err := ioutil.ReadFile("./2022/7/input.txt")
	if err != nil {
		panic(err)
	}

	root := Parse(string(input))
	folders := root.FolderLessThanSize(100000)
	var size int
	for _, folder := range folders {
		size += folder.Size()
	}
	println("Part 1:", size)

	//The total disk space available to the filesystem is 70000000.
	//To run the update, you need unused space of at least 30000000.
	//You need to find a directory you can delete that will free up enough space to run the update.
	//Find the smallest directory that, if deleted, would free up enough space on the filesystem to run the update.
	unusedSpace := 70000000 - root.Size()
	spaceNeedToFree := 30000000 - unusedSpace
	sizeMap := root.FolderSizeMap()
	//sort the sizeMap by size
	var sizes []int
	for _, size := range sizeMap {
		sizes = append(sizes, size)
	}
	for i := 0; i < len(sizes); i++ {
		for j := i + 1; j < len(sizes); j++ {
			if sizes[i] > sizes[j] {
				sizes[i], sizes[j] = sizes[j], sizes[i]
			}
		}
	}

	//find the smallest folder that can free up enough space
	for _, size := range sizes {
		if size >= spaceNeedToFree {
			println("Part 2:", size)
			break
		}
	}
}
