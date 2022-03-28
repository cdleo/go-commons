/*
Package e2hformat is the formatter's package of the Enhanced Error Handling module
*/
package formatter

import (
	"path/filepath"
	"strings"
)

type HidingMethod int8

//Allowed path treatments
const (
	HidingMethod_None HidingMethod = iota

	HidingMethod_FullBaseline

	HidingMethod_ToFolder
)

// This function format the sourcefile according to the provided params
func FormatSourceFile(file string, hidingMethod HidingMethod, hidingValue string) string {
	switch hidingMethod {
	case HidingMethod_FullBaseline:
		return removePathSegment(file, hidingValue)
	case HidingMethod_ToFolder:
		return removeBeforeFolder(file, hidingValue)
	default: //HidingMethod_None
		return file
	}
}

// Utility funtion that removes the first part of the filepath til the end of `baseline` path argument
func removePathSegment(file string, baseline string) string {

	file = filepath.Clean(file)
	prettyCaller := strings.ReplaceAll(file, baseline, "")
	if len(prettyCaller) > 0 {
		return prettyCaller
	}

	return file
}

// Utility funtion that removes the first part of the filepath til found the folder indicated in `newRootFolder` argument
func removeBeforeFolder(file string, newRootFolder string) string {

	fileParts := strings.Split(file, newRootFolder)
	if len(fileParts) < 2 {
		return file
	}
	return newRootFolder + fileParts[1]
}
