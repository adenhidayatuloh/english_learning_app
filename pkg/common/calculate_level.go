package common

func CalculateLevel(totalExp int) (level int, nextLevelExp int) {
	level = 1
	xpForNextLevel := 50
	//totalExpAll := totalExp

	for totalExp >= xpForNextLevel {
		totalExp -= xpForNextLevel
		level++
		//xpForNextLevel = xpForNextLevel + 50
	}
	nextLevelExp = xpForNextLevel * level
	return
}
