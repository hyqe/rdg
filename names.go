package rdg

func LastName() string {
	return lastNames[IntBetween(0, len(lastNames))]
}

func FirstName() string {
	return firstNames[IntBetween(0, len(firstNames))]
}

func FirstNameFemale() string {
	return femaleNames[IntBetween(0, len(femaleNames))]
}

func FirstNameMale() string {
	return maleNames[IntBetween(0, len(maleNames))]
}

func PrefixFemale() string {
	return femalePrefixes[IntBetween(0, len(femalePrefixes))]
}

func PrefixMale() string {
	return malePrefixes[IntBetween(0, len(malePrefixes))]
}

func Prefix() string {
	return prefixes[IntBetween(0, len(prefixes))]
}

func SuffixFemale() string {
	return femaleSuffixes[IntBetween(0, len(femaleSuffixes))]
}

func SuffixMale() string {
	return maleSuffixes[IntBetween(0, len(maleSuffixes))]
}

func Suffix() string {
	return suffixes[IntBetween(0, len(suffixes))]
}
