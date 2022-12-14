package main

import (
	"bufio"
	"fmt"
	"runtime"
	"strconv"
	"strings"

	"github.com/monicagangwar/advent-of-code/input"
)

/*
orientation

1  x,  y,  z      9   y,  z,  x    17  z,  y,  x
2  x,  y, -z      10  y,  z, -x    18  z,  y, -x
3  x, -y,  z      11  y, -z,  x    19  z, -y,  x
4  x, -y, -z      12  y, -z, -x    20  z, -y, -x
5 -x,  y,  z      13 -y,  z,  x    21 -z,  y,  x
6 -x,  y, -z      14 -y,  z, -x    22 -z,  y, -x
7 -x, -y,  z      15 -y, -z,  x    23 -z, -y,  x
8 -x, -y, -z      16 -y, -z, -x    24 -z, -y, -x
*/

type pos struct {
	x           int64
	y           int64
	z           int64
	orientation int
}

func main() {
	_, currentFilePath, _, _ := runtime.Caller(0)
	file := input.GetFileMarker(currentFilePath)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scannerAndBeacons := make([][]pos, 0)
	scannerIdx := -1

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		if strings.Contains(line, "scanner") {
			scannerIdx++
		} else {
			beaconPos := pos{}
			posMarkers := strings.Split(line, ",")
			beaconPos.x, _ = strconv.ParseInt(posMarkers[0], 10, 64)
			beaconPos.y, _ = strconv.ParseInt(posMarkers[1], 10, 64)
			beaconPos.z, _ = strconv.ParseInt(posMarkers[2], 10, 64)

			if len(scannerAndBeacons) == scannerIdx {
				scannerAndBeacons = append(scannerAndBeacons, []pos{beaconPos})
			} else {
				scannerAndBeacons[scannerIdx] = append(scannerAndBeacons[scannerIdx], beaconPos)
			}
		}
	}

	partOne(scannerAndBeacons)
}

func adjustPosToOrientation(beaconPos pos) pos {
	adjustedPos := beaconPos

	if beaconPos.orientation >= 9 && beaconPos.orientation <= 16 {
		adjustedPos.x = beaconPos.y
		adjustedPos.y = beaconPos.z
		adjustedPos.z = beaconPos.x
	} else if beaconPos.orientation >= 17 && beaconPos.orientation <= 24 {
		adjustedPos.x = beaconPos.z
		adjustedPos.y = beaconPos.x
		adjustedPos.z = beaconPos.y
	}

	orientationForBitConversion := beaconPos.orientation - 1
	if orientationForBitConversion&1 == 1 {
		adjustedPos.z *= -1
	}
	orientationForBitConversion = orientationForBitConversion >> 1
	if orientationForBitConversion&1 == 1 {
		adjustedPos.y *= -1
	}
	orientationForBitConversion = orientationForBitConversion >> 1
	if orientationForBitConversion&1 == 1 {
		adjustedPos.x *= -1
	}

	return adjustedPos
}

func adjustBeaconsToScannerPos(beaconPos pos, scannerPos pos) pos {
	newPos := beaconPos

	orientationForBitConversion := scannerPos.orientation - 1
	if orientationForBitConversion&1 == 1 {
		newPos.z *= -1
	}
	orientationForBitConversion = orientationForBitConversion >> 1
	if orientationForBitConversion&1 == 1 {
		newPos.y *= -1
	}
	orientationForBitConversion = orientationForBitConversion >> 1
	if orientationForBitConversion&1 == 1 {
		newPos.x *= -1
	}

	if scannerPos.orientation >= 9 && scannerPos.orientation <= 16 {
		newPos.x = beaconPos.z
		newPos.y = beaconPos.x
		newPos.z = beaconPos.y
	} else if scannerPos.orientation >= 17 && scannerPos.orientation <= 24 {
		newPos.x = beaconPos.y
		newPos.y = beaconPos.z
		newPos.z = beaconPos.x
	}

	newPos.x += scannerPos.x
	newPos.y += scannerPos.y
	newPos.z += scannerPos.z

	return newPos
}

func partOne(scannerAndBeacons [][]pos) {
	scannerActualPos := make(map[int]pos, 0)
	scannerPosFound := 1
	scannerActualPos[0] = pos{0, 0, 0, 1}
	scannerFoundList := make([]bool, len(scannerAndBeacons))
	scannerFoundList[0] = true
	scannerVisited := make([]bool, len(scannerAndBeacons))
	allScannerVisited := true
	for {
		for scannerFoundIdx, scannerFoundPos := range scannerActualPos {
			if scannerVisited[scannerFoundIdx] {
				allScannerVisited = true
				continue
			} else {
				scannerVisited[scannerFoundIdx] = true
				allScannerVisited = false
			}
			for scannerIdx := 0; scannerIdx < len(scannerAndBeacons); scannerIdx++ {
				if !scannerFoundList[scannerIdx] && scannerIdx != scannerFoundIdx {
					beaconsDelta := make(map[pos]int)
					for _, beacon1Pos := range scannerAndBeacons[scannerFoundIdx] {
						//adjustedBeacon1Pos := adjustBeaconsToScannerPos(beacon1Pos, scannerFoundPos)
						for _, beacon2Pos := range scannerAndBeacons[scannerIdx] {
							for orientation := 1; orientation <= 24; orientation++ {
								adjustedBeacon2Pos := beacon2Pos
								adjustedBeacon2Pos.orientation = orientation
								adjustedBeacon2Pos = adjustPosToOrientation(adjustedBeacon2Pos)
								//fmt.Printf("\n original: %+v, adjusted: %+v", beacon2Pos, adjustedBeacon2Pos)
								deltaPos := pos{
									x:           beacon1Pos.x - adjustedBeacon2Pos.x,
									y:           beacon1Pos.y - adjustedBeacon2Pos.y,
									z:           beacon1Pos.z - adjustedBeacon2Pos.z,
									orientation: orientation,
								}
								if freq, found := beaconsDelta[deltaPos]; found {
									beaconsDelta[deltaPos] = freq + 1
								} else {
									beaconsDelta[deltaPos] = 1
								}
							}
						}
					}
					fmt.Printf("\n Delta between scanner: %d and scanner: %d ", scannerFoundIdx, scannerIdx)
					for delta, freq := range beaconsDelta {
						//fmt.Printf("%d", freq)
						if freq >= 12 {
							fmt.Printf("%+v: %d", delta, freq)
							scannerFoundList[scannerIdx] = true
							scannerPosFound++

							scannerActualPos[scannerIdx] = adjustBeaconsToScannerPos(delta, scannerFoundPos)

							//scanner1Pos := scannerActualPos[scannerFoundIdx]
							//scannerActualPos[scannerIdx] = pos{
							//	x:           scanner1Pos.x + delta.x,
							//	y:           scanner1Pos.y + delta.y,
							//	z:           scanner1Pos.z + delta.z,
							//	orientation: delta.orientation,
							//}

						}
					}
				}
			}
		}
		if allScannerVisited {
			break
		}
	}

	allBeacons := make(map[pos]struct{})

	for scannerIdx, beacons := range scannerAndBeacons {
		scannerPos := scannerActualPos[scannerIdx]
		for _, beacon := range beacons {
			beacon.orientation = scannerPos.orientation
			newPos := adjustBeaconsToScannerPos(beacon, scannerPos)
			newPos.orientation = 0
			allBeacons[newPos] = struct{}{}
		}
	}

	fmt.Printf("\n%+v", scannerActualPos)
	fmt.Printf("\nall beacon count: %d", len(allBeacons))
}
