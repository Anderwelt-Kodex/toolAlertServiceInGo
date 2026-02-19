package toolservicealert

import "errors"

type ToolInformation struct {
	toolId string
	//toolConditionStatus bool
	isServiceDue        bool
	minServiceThreshold int //warning
	maxServiceThreshold int //critical
	//toolScheduledService int
	toolCumulativeRuntime int
}

// pupblic function called from session
func CallingToolServiceAlert(toolId string, reportedRuntime int) (int, error) {

	//get the recorded data included in ToolInformation of the tool
	toolData, errToolData := GetToolInformation(toolId)

	//check for errors, can't work with the information
	if errToolData != nil {
		return 3, errToolData
	}
	//compare service thresholds with cumulative runtime, to get an int determine which kind of alert should be sent
	resultServiceAlert, errMsg := compareRuntimeToolToServiceThresholds(toolData.toolCumulativeRuntime, reportedRuntime, toolData.maxServiceThreshold, toolData.minServiceThreshold)
	return resultServiceAlert, errMsg

}

func GetToolInformation(toolIdent string) (ToolInformation, error) {

	var activeToolData ToolInformation

	//database information for testing

	switch toolIdent {
	case "tool001":
		activeToolData = ToolInformation{ //middle service needs
			toolId:                "tool001",
			isServiceDue:          false,
			minServiceThreshold:   5,
			maxServiceThreshold:   15,
			toolCumulativeRuntime: 3}

	case "tool002":
		activeToolData = ToolInformation{ //if you let it run a lot service
			toolId:                "tool002",
			isServiceDue:          false,
			minServiceThreshold:   20,
			maxServiceThreshold:   35,
			toolCumulativeRuntime: 10}

	case "tool003":
		activeToolData = ToolInformation{ //always service
			toolId:                "tool003",
			isServiceDue:          false,
			minServiceThreshold:   2,
			maxServiceThreshold:   5,
			toolCumulativeRuntime: 1}

	default:
		//activeToolData from database
	}

	//check if empty - no values
	if activeToolData == (ToolInformation{}) {

		return activeToolData, errors.New("Tool information not found")
	}

	return activeToolData, nil

}

func CalculateToolRuntime(toolCumulativeRuntime int, toolSessionRuntime int) (int, error) {

	toolRuntimeCounter := toolCumulativeRuntime + toolSessionRuntime

	//check value is realistic
	if toolRuntimeCounter <= 0 {
		return toolRuntimeCounter, errors.New("calculated runtime for tool <= 0")
	}

	return toolRuntimeCounter, nil
}

func compareRuntimeToolToServiceThresholds(toolCumulativeRuntime int, toolSessionRuntime int, maxServiceThreshold int, minServiceThreshold int) (int, error) {

	toolRuntime, err := CalculateToolRuntime(toolCumulativeRuntime, toolSessionRuntime)

	//the calculated runtime is 0 or negative
	if err != nil {
		return 4, err
	}
	//critical warning - tool will reach the maximum guaranteed runtime during session
	if toolRuntime >= maxServiceThreshold {
		return 2, nil
	}
	//warning - tool will reach the minimum guaranteed runtime during session
	if toolRuntime >= minServiceThreshold {
		return 1, nil
	}
	//default no service thresholds will be breeched, proceed
	return 0, nil

}
