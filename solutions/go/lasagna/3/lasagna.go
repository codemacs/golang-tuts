package lasagna 
const (
    OvenTime = 40 
    minsTakenPerLayer = 2 
)

func RemainingOvenTime(actualMinutesInOven int) int {
    return OvenTime - actualMinutesInOven
}

func PreparationTime(numberOfLayers int ) int {
    return numberOfLayers*minsTakenPerLayer
}

func ElapsedTime(numberOfLayers int, actualMinutesInOven int) int {
    return (PreparationTime(numberOfLayers) + actualMinutesInOven)
}
