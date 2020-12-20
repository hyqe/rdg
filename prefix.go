package rdg

import (
	"github.com/LukeOchoa/usefulThings"
)

var prefixes = []string{
	"1st Lt",
	"Adm",
	"Atty",
	"Brother",
	"Capt",
	"Chief",
	"Cmdr",
	"Col",
	"Dean",
	"Dr",
	"Elder",
	"Father",
	"Gen",
	"Gov",
	"Hon",
	"Lt Col",
	"Maj",
	"MSgt",
	"Mr",
	"Mrs",
	"Ms",
	"Prince",
	"Prof",
	"Rabbi",
	"Rev",
	"Sister",
	"II",
	"III",
	"IV",
	"CPA",
	"DDS",
	"Esq",
	"JD",
	"Jr",
	"LLD",
	"MD",
	"PhD",
	"Ret",
	"RN",
	"Sr",
	"DO",
}

func Prefix() string {
	return prefixes[usefulThings.NewInt32Between(0, int32(len(prefixes)))]
}
