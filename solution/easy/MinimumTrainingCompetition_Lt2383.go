package easy

func minNumberOfHours(initialEnergy int, initialExperience int, energy []int, experience []int) int {
	sum := 0
	for _, e := range energy {
		sum += e
	}
	trainingHours := 0
	if initialEnergy <= sum {
		trainingHours = sum + 1 - initialEnergy
	}
	for _, e := range experience {
		if initialExperience <= e {
			trainingHours += 1 + (e - initialExperience)
			initialExperience = 2*e + 1
		} else {
			initialExperience += e
		}
	}
	return trainingHours
}
