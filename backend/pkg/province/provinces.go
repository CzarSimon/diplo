package province

// GetProvinces gets a map of all provinces, with province shortname as key.
func GetProvinces() map[string]Province {
	return map[string]Province{
		NorthAtlantic.ShortName:    NorthAtlantic,
		NorwegianSea.ShortName:     NorwegianSea,
		BarnetsSea.ShortName:       BarnetsSea,
		StPetersburg.ShortName:     StPetersburg,
		Clyde.ShortName:            Clyde,
		Edinburgh.ShortName:        Edinburgh,
		NorthSea.ShortName:         NorthSea,
		Norway.ShortName:           Norway,
		Sweden.ShortName:           Sweden,
		GulfOfBotnia.ShortName:     GulfOfBotnia,
		IrishSea.ShortName:         IrishSea,
		Liverpool.ShortName:        Liverpool,
		Yorkshire.ShortName:        Yorkshire,
		Skagerrack.ShortName:       Skagerrack,
		Denmark.ShortName:          Denmark,
		BalticSea.ShortName:        BalticSea,
		Wales.ShortName:            Wales,
		London.ShortName:           London,
		Holland.ShortName:          Holland,
		HeligolandBight.ShortName:  HeligolandBight,
		Kiel.ShortName:             Kiel,
		Berlin.ShortName:           Berlin,
		Prussia.ShortName:          Prussia,
		Moscow.ShortName:           Moscow,
		MidAtlanticOcean.ShortName: MidAtlanticOcean,
		Brest.ShortName:            Brest,
		Picardy.ShortName:          Picardy,
		Belgium.ShortName:          Belgium,
		Ruhr.ShortName:             Ruhr,
		Munich.ShortName:           Munich,
		Warsaw.ShortName:           Warsaw,
		Ukraine.ShortName:          Ukraine,
		Sevastopol.ShortName:       Sevastopol,
		Paris.ShortName:            Paris,
		Gascony.ShortName:          Gascony,
		Munich.ShortName:           Munich,
		Bohemia.ShortName:          Bohemia,
	}
}

// NorthAtlantic Sea, a netural water province.
var NorthAtlantic = Province{
	Name:            "North Atlantic Ocean",
	ShortName:       "NAt",
	HasSupplyCenter: false,
	Type:            Water,
	Neighbours:      []string{"Nrg", "Cly", "Lvp", "Iri", "Mid"},
}

// Norwegian Sea, a netural water province.
var NorwegianSea = Province{
	Name:            "Norwegian Sea",
	ShortName:       "Nrg",
	HasSupplyCenter: false,
	Type:            Water,
	Neighbours:      []string{"Bar", "Nwy", "Nth", "Edi", "Cly", "NAt"},
}

// Barnets Sea, a netural water province.
var BarnetsSea = Province{
	Name:            "Barnets Sea",
	ShortName:       "Bar",
	HasSupplyCenter: false,
	Type:            Water,
	Neighbours:      []string{"StP", "Nwy", "Nrg"},
}

// StPetersburg, a Russian costal province.
var StPetersburg = Province{
	Name:            "St. Petersburg",
	ShortName:       "StP",
	HasSupplyCenter: true,
	Type:            Costal,
	Neighbours:      []string{"Mos", "Lvn", "Bot", "Fin", "Nwy", "Bar"},
}

// Clyde, an English costal province.
var Clyde = Province{
	Name:            "Clyde",
	ShortName:       "Cly",
	HasSupplyCenter: false,
	Type:            Costal,
	Neighbours:      []string{"NAt", "Nrg", "Edi", "Lvp"},
}

// Edinburgh, an English costal province.
var Edinburgh = Province{
	Name:            "Edinburgh",
	ShortName:       "Edi",
	HasSupplyCenter: true,
	Type:            Costal,
	Neighbours:      []string{"Nrg", "Nth", "Yor", "Lvp", "Cly"},
}

// North Sea, a netural water province.
var NorthSea = Province{
	Name:            "North Sea",
	ShortName:       "Nth",
	HasSupplyCenter: false,
	Type:            Water,
	Neighbours:      []string{"Nrg", "Nwy", "Ska", "Den", "Hel", "Hol", "Bel", "Eng", "Lon", "Yor", "Edi"},
}

// Norway, a netural costal province.
var Norway = Province{
	Name:            "Norway",
	ShortName:       "Nwy",
	HasSupplyCenter: true,
	Type:            Costal,
	Neighbours:      []string{"Nrg", "Bar", "StP", "Fin", "Swe", "Ska", "Nth"},
}

// Sweden, a neutral costal province.
var Sweden = Province{
	Name:            "Sweden",
	ShortName:       "Swe",
	HasSupplyCenter: true,
	Type:            Costal,
	Neighbours:      []string{"Nwy", "Fin", "Bot", "Bal", "Den", "Ska"},
}

// Gulf of Botnia, a netural water province.
var GulfOfBotnia = Province{
	Name:            "Gulf of Botnia",
	ShortName:       "Bot",
	HasSupplyCenter: false,
	Type:            Water,
	Neighbours:      []string{"Swe", "Fin", "StP", "Lvn", "Bal"},
}

// Irish sea.
var IrishSea = Province{
	Name:            "Irish Sea",
	ShortName:       "Iri",
	HasSupplyCenter: false,
	Type:            Water,
	Neighbours:      []string{"Nat", "Cly", "Lvp", "Wal", "Eng", "Mid"},
}

// Liverpool, an English costal province.
var Liverpool = Province{
	Name:            "Liverpool",
	ShortName:       "Lvp",
	HasSupplyCenter: true,
	Type:            Costal,
	Neighbours:      []string{"Cly", "Edi", "Yor", "Wal", "Iri", "NAt"},
}

// Yorkshire, an English costal province.
var Yorkshire = Province{
	Name:            "Yorkshire",
	ShortName:       "Yor",
	HasSupplyCenter: false,
	Type:            Costal,
	Neighbours:      []string{"Edi", "Nth", "Lon", "Wal", "Lvp"},
}

// Skagerrack, a netural water province.
var Skagerrack = Province{
	Name:            "Skagerrack",
	ShortName:       "Ska",
	HasSupplyCenter: false,
	Type:            Water,
	Neighbours:      []string{"Nwy", "Swe", "Den", "Nth"},
}

// Denmark, a netural costal province.
var Denmark = Province{
	Name:            "Denmark",
	ShortName:       "Den",
	HasSupplyCenter: true,
	Type:            Costal,
	Neighbours:      []string{"Ska", "Swe", "Bal", "Kie", "Hel", "Nth"},
}

// Baltic Sea, a netural water province.
var BalticSea = Province{
	Name:            "Baltic Sea",
	ShortName:       "Bal",
	HasSupplyCenter: false,
	Type:            Water,
	Neighbours:      []string{"Swe", "Bot", "Liv", "Pru", "Ber", "Kie", "Den"},
}

// Wales, an English costal province.
var Wales = Province{
	Name:            "Wales",
	ShortName:       "Wal",
	HasSupplyCenter: false,
	Type:            Costal,
	Neighbours:      []string{"Lvp", "Yor", "Lon", "Eng", "Iri"},
}

// London, an English costal province.
var London = Province{
	Name:            "London",
	ShortName:       "Lon",
	HasSupplyCenter: true,
	Type:            Costal,
	Neighbours:      []string{"Yor", "Nth", "Lon", "Eng", "Wal"},
}

// Holland, a netural costal province.
var Holland = Province{
	Name:            "Holland",
	ShortName:       "Hol",
	HasSupplyCenter: true,
	Type:            Costal,
	Neighbours:      []string{"Nth", "Hel", "Kie", "Ruh", "Bel"},
}

// Heligoland Bight, a neutral water province.
var HeligolandBight = Province{
	Name:            "HeligolandBight",
	ShortName:       "Hel",
	HasSupplyCenter: false,
	Type:            Water,
	Neighbours:      []string{"Nth", "Den", "Kie", "Hol"},
}

// Kiel, German costal province.
var Kiel = Province{
	Name:            "Kiel",
	ShortName:       "Kie",
	HasSupplyCenter: true,
	Type:            Costal,
	Neighbours:      []string{"Den", "Ber", "Mun", "Ruh", "Hol", "Hel"},
}

// Berlin, a German costal province.
var Berlin = Province{
	Name:            "Berlin",
	ShortName:       "Ber",
	HasSupplyCenter: true,
	Type:            Costal,
	Neighbours:      []string{"Bal", "Pru", "Sil", "Mun", "Kie"},
}

// Prussia, a German costal province.
var Prussia = Province{
	Name:            "Prussia",
	ShortName:       "Pru",
	HasSupplyCenter: false,
	Type:            Costal,
	Neighbours:      []string{"Bal", "Lvn", "War", "Sil", "Ber"},
}

// Moscow, a Russian land province.
var Moscow = Province{
	Name:            "Moscow",
	ShortName:       "Mos",
	HasSupplyCenter: true,
	Type:            Land,
	Neighbours:      []string{"StP", "Sev", "Ukr", "War", "Lvn"},
}

// Mid-Atlantic Ocean, a netural water province.
var MidAtlanticOcean = Province{
	Name:            "Mid-Atlantic Ocean",
	ShortName:       "Mid",
	HasSupplyCenter: false,
	Type:            Water,
	Neighbours:      []string{"NAt", "Iri", "Eng", "Bre", "Gas", "Spa", "Por"},
}

// Brest, a French costal province.
var Brest = Province{
	Name:            "Brest",
	ShortName:       "Bre",
	HasSupplyCenter: true,
	Type:            Costal,
	Neighbours:      []string{"Eng", "Pic", "Par", "Gas", "Mid"},
}

// Picardy, a French costal province.
var Picardy = Province{
	Name:            "Picardy",
	ShortName:       "Pic",
	HasSupplyCenter: false,
	Type:            Costal,
	Neighbours:      []string{"Eng", "Bel", "Bur", "Par", "Bre"},
}

// Belgium, a neutral costal province.
var Belgium = Province{
	Name:            "Belgium",
	ShortName:       "Bel",
	HasSupplyCenter: true,
	Type:            Costal,
	Neighbours:      []string{"Nth", "Hol", "Ruh", "Bur", "Pic", "Eng"},
}

// Ruhr, a German land province.
var Ruhr = Province{
	Name:            "Ruhr",
	ShortName:       "Ruh",
	HasSupplyCenter: false,
	Type:            Land,
	Neighbours:      []string{"Kie", "Mun", "Bur", "Bel", "Hol"},
}

// Munich, a German land province.
var Munich = Province{
	Name:            "Munich",
	ShortName:       "Mun",
	HasSupplyCenter: true,
	Type:            Land,
	Neighbours:      []string{"Ber", "Sil", "Boh", "Tyr", "Ruh", "Kie"},
}

// Silesia, a German land province.
var Silesia = Province{
	Name:            "Silesia",
	ShortName:       "Sil",
	HasSupplyCenter: false,
	Type:            Land,
	Neighbours:      []string{"Pru", "War", "Gal", "Boh", "Mun", "Ber"},
}

// Warsaw, a Russian land province.
var Warsaw = Province{
	Name:            "Warsaw",
	ShortName:       "War",
	HasSupplyCenter: true,
	Type:            Land,
	Neighbours:      []string{"Pru", "Lvn", "Mos", "Ukr", "Gal", "Sil"},
}

// Ukraine, a Russian land province.
var Ukraine = Province{
	Name:            "Ukraine",
	ShortName:       "Ukr",
	HasSupplyCenter: false,
	Type:            Land,
	Neighbours:      []string{"Mos", "Sev", "Rum", "Gal", "War"},
}

// Sevastopol, a Russian costal province.
var Sevastopol = Province{
	Name:            "Sevastopol",
	ShortName:       "Sev",
	HasSupplyCenter: true,
	Type:            Costal,
	Neighbours:      []string{"Mos", "Arm", "Bla", "Rum", "Ukr"},
}

// Paris, a French land province.
var Paris = Province{
	Name:            "Paris",
	ShortName:       "Par",
	HasSupplyCenter: true,
	Type:            Land,
	Neighbours:      []string{"Pic", "Bur", "Gas", "Bre"},
}

// Gascony, a French land province.
var Gascony = Province{
	Name:            "Gascony",
	ShortName:       "Gas",
	HasSupplyCenter: false,
	Type:            Land,
	Neighbours:      []string{"Bel", "Ruh", "Mun", "Mar", "Par", "Pic"},
}

// Bohemia, an Austrian land province.
var Bohemia = Province{
	Name:            "Bohemia",
	ShortName:       "Boh",
	HasSupplyCenter: false,
	Type:            Land,
	Neighbours:      []string{"Sil", "Gal", "Vie", "Boh", "Tyr", "Mun"},
}

/*
var  = Province{
	Name:            "",
	ShortName:       "",
	HasSupplyCenter: ,
	Type:            ,
	Neighbours:        []string{"", "", "", "", ""},
}

/*
var  = Province{
	Name:            "",
	ShortName:       "",
	HasSupplyCenter: ,
	Type:            ,
	Neighbours:        []string{"", "", "", "", ""},
}

/*
var  = Province{
	Name:            "",
	ShortName:       "",
	HasSupplyCenter: ,
	Type:            ,
	Neighbours:        []string{"", "", "", "", ""},
}

/*
var  = Province{
	Name:            "",
	ShortName:       "",
	HasSupplyCenter: ,
	Type:            ,
	Neighbours:        []string{"", "", "", "", ""},
}

/*
var  = Province{
	Name:            "",
	ShortName:       "",
	HasSupplyCenter: ,
	Type:            ,
	Neighbours:        []string{"", "", "", "", ""},
}

/*
var  = Province{
	Name:            "",
	ShortName:       "",
	HasSupplyCenter: ,
	Type:            ,
	Neighbours:        []string{"", "", "", "", ""},
}

/*
var  = Province{
	Name:            "",
	ShortName:       "",
	HasSupplyCenter: ,
	Type:            ,
	Neighbours:        []string{"", "", "", "", ""},
}

/*
var  = Province{
	Name:            "",
	ShortName:       "",
	HasSupplyCenter: ,
	Type:            ,
	Neighbours:        []string{"", "", "", "", ""},
}

/*
var  = Province{
	Name:            "",
	ShortName:       "",
	HasSupplyCenter: ,
	Type:            ,
	Neighbours:        []string{"", "", "", "", ""},
}

/*
var  = Province{
	Name:            "",
	ShortName:       "",
	HasSupplyCenter: ,
	Type:            ,
	Neighbours:        []string{"", "", "", "", ""},
}
*/

// Syria province.
var Syria = Province{
	Name:            "Syria",
	ShortName:       "Syr",
	HasSupplyCenter: false,
	Type:            Costal,
	Neighbours:      []string{"Arm", "Eas", "Smy"},
}

// Armenia province.
var Armenia = Province{
	Name:            "Armenia",
	ShortName:       "Arm",
	HasSupplyCenter: false,
	Type:            Costal,
	Neighbours:      []string{"Sev", "Syr", "Smy", "Ank", "Bla"},
}
