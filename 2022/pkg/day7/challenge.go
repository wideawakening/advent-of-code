package day7

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Star1(inputFileName string) int {

	//-----------------------------------------------
	data, err := os.ReadFile(inputFileName)
	if err != nil {
		fmt.Print(err)
		return 0
	}

	commands := []string{}
	for _, dataLine := range strings.Split(string(data), "\n") {
		//fmt.Printf("%s", dataLine)
		commands = append(commands, dataLine)
	}
	folderSizes := GetFolderSizes(commands)
	//-----------------------------------------------

	lowerThan100KFolderSizes := GetOnlyLowerThan100000(folderSizes)

	totalSize := 0
	for _, v := range lowerThan100KFolderSizes {
		totalSize += v
	}
	return totalSize
}

func Star2(inputFileName string) int {

	//-----------------------------------------------
	data, err := os.ReadFile(inputFileName)
	if err != nil {
		fmt.Print(err)
		return 0
	}

	commands := []string{}
	for _, dataLine := range strings.Split(string(data), "\n") {
		//fmt.Printf("%s", dataLine)
		commands = append(commands, dataLine)
	}
	folderSizes := GetFolderSizes(commands)
	//-----------------------------------------------

	const totalSize = 70000000
	usedSpace := folderSizes["/"]
	availableSpace := totalSize - usedSpace
	const requiredSpaceForUppate = 30000000
	needToDelete := requiredSpaceForUppate - availableSpace

	// find lowestValue required for update
	lowestValue := requiredSpaceForUppate
	for _, v := range folderSizes {
		if v > needToDelete {
			if v < lowestValue {
				lowestValue = v
			}
		}
	}

	return lowestValue
}

func GetFolderSizes(commands []string) map[string]int {

	folderSizes := make(map[string]int)
	var pwd string
	for i := 0; i < len(commands); i++ {
		commandLine := commands[i]
		if !IsCommand(commandLine) {
			continue
		}
		command, argument := ExtractCommand(commandLine)
		switch command {
		case "cd":
			pwd = ProcessCwd(pwd, argument)
		case "ls":
			keepLS := true
			folderSize := 0
			for ; keepLS; i++ {
				commandLine = commands[i]
				folderSize += ProcessFileSize(commandLine)

				// end of LS
				if i+1 >= len(commands) || IsCommand(commands[i+1]) {
					keepLS = false
					i--
				}
			}
			SaveFolderSize(folderSizes, pwd, folderSize)
		}
	}

	return folderSizes
}

func GetOnlyLowerThan100000(folders map[string]int) map[string]int {
	folders1000K := make(map[string]int)
	for k, v := range folders {
		if v <= 100000 {
			folders1000K[k] = v
		}
	}

	return folders1000K
}

func IsCommand(commandLine string) bool {
	return len(commandLine) > 0 && commandLine[0] == '$'
}

func ExtractCommand(commandLine string) (string, string) {
	arguments := strings.Split(commandLine, " ")
	if len(arguments) == 2 {
		return arguments[1], ""
	}
	return arguments[1], arguments[2]
}

func ProcessCwd(cwd string, argument string) string {
	switch argument {
	case "/":
		return "/"

	case "..":
		if strings.Contains(cwd, "/") {
			lastDir := strings.LastIndex(cwd, "/")
			if lastDir == 0 {
				return "/"
			}
			return cwd[0:lastDir]
		}

	}
	if cwd == "/" {
		return cwd + argument
	}

	return cwd + "/" + argument
}

func ProcessFileSize(commandLine string) int {
	total, err := strconv.Atoi(strings.Split(commandLine, " ")[0])
	if err != nil {
		return 0
	}

	return total
}

func SaveFolderSize(folderSizes map[string]int, fullFolderPath string, size int) map[string]int {
	for _, subfolder := range GetSubFolders(fullFolderPath) {
		subfolderSize := folderSizes[subfolder]
		if subfolderSize == 0 {
			folderSizes[subfolder] = size
		} else {
			folderSizes[subfolder] = subfolderSize + size
		}
	}

	return folderSizes
}

func GetSubFolders(fullFolderPath string) []string {

	if string(fullFolderPath[0]) != "/" {
		return []string{}
	}
	if fullFolderPath == "/" {
		return []string{"/"}
	}

	var tempFullFolderPath string = fullFolderPath
	if string(fullFolderPath[len(fullFolderPath)-1]) == "/" {
		tempFullFolderPath = fullFolderPath[0 : len(fullFolderPath)-1]
	}

	allPaths := []string{"/"}

	tempFullFolderPath = tempFullFolderPath[1:] + "/"

	nextPathIndex := strings.Index(tempFullFolderPath, "/")
	if nextPathIndex == -1 {
		return []string{"/" + tempFullFolderPath}
	}

	lastPath := ""
	for nextPathIndex != -1 {
		path := tempFullFolderPath[:nextPathIndex]
		lastPath = lastPath + "/" + path
		allPaths = append(allPaths, lastPath)
		tempFullFolderPath = tempFullFolderPath[nextPathIndex+1:]
		nextPathIndex = strings.Index(tempFullFolderPath, "/")
	}

	return allPaths

}
