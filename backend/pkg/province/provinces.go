package province

const (
	NAt   ShortName = "NAt"
	Nrg             = "Nrg"
	Bar             = "Bar"
	StP             = "StP"
	StPNC           = "StP NC"
	StPSC           = "StP SC"
	Cly             = "Cly"
	Edi             = "Edi"
	Nth             = "Nth"
	Nwy             = "Nwy"
	Swe             = "Swe"
	Bot             = "Bot"
	Fin             = "Fin"
	Iri             = "Iri"
	Lvp             = "Lvp"
	Yor             = "Yor"
	Ska             = "Ska"
	Den             = "Den"
	Bal             = "Bal"
	Lvn             = "Lvn"
	Wal             = "Wal"
	Lon             = "Lon"
	Hol             = "Hol"
	Hel             = "Hel"
	Kie             = "Kie"
	Ber             = "Ber"
	Pru             = "Pru"
	Mos             = "Mos"
	Mid             = "Mid"
	Eng             = "Eng"
	Bre             = "Bre"
	Pic             = "Pic"
	Bel             = "Bel"
	Ruh             = "Ruh"
	Mun             = "Mun"
	Sil             = "Sil"
	War             = "War"
	Ukr             = "Ukr"
	Sev             = "Sev"
	Par             = "Par"
	Bur             = "Bur"
	Boh             = "Boh"
	Gal             = "Gal"
	Gas             = "Gas"
	Mar             = "Mar"
	Pie             = "Pie"
	Ven             = "Ven"
	Adr             = "Adr"
	Tyr             = "Tyr"
	Tri             = "Tri"
	Vie             = "Vie"
	Bud             = "Bud"
	Rum             = "Rum"
	Por             = "Por"
	Spa             = "Spa"
	SpaNC           = "Spa NC"
	SpaSC           = "Spa SC"
	GoL             = "GoL"
	Tus             = "Tus"
	Rom             = "Rom"
	Apu             = "Apu"
	Alb             = "Alb"
	Ser             = "Ser"
	Bul             = "Bul"
	BulSC           = "Bul SC"
	BulEC           = "Bul EC"
	Bla             = "Bla"
	Wes             = "Wes"
	Tyn             = "Tyn"
	Nap             = "Nap"
	Ion             = "Ion"
	Gre             = "Gre"
	Aeg             = "Aeg"
	Con             = "Con"
	Smy             = "Smy"
	Ank             = "Ank"
	Arm             = "Arm"
	NAf             = "NAf"
	Tun             = "Tun"
	Eas             = "Eas"
	Syr             = "Syr"
)

// GetProvinces gets a map of all provinces, with province shortname as key.
func GetProvinces() map[ShortName]Province {
	return map[ShortName]Province{
		NAt:   NorthAtlantic,
		Nrg:   NorwegianSea,
		Bar:   BarnetsSea,
		StP:   StPetersburg,
		StPNC: StPetersburgNorthCoast,
		StPSC: StPetersburgSouthCoast,
		Cly:   Clyde,
		Edi:   Edinburgh,
		Nth:   NorthSea,
		Nwy:   Norway,
		Swe:   Sweden,
		Bot:   GulfOfBotnia,
		Fin:   Finland,
		Iri:   IrishSea,
		Lvp:   Liverpool,
		Yor:   Yorkshire,
		Ska:   Skagerrack,
		Den:   Denmark,
		Bal:   BalticSea,
		Lvn:   Livonia,
		Wal:   Wales,
		Lon:   London,
		Hol:   Holland,
		Hel:   HeligolandBight,
		Kie:   Kiel,
		Ber:   Berlin,
		Pru:   Prussia,
		Mos:   Moscow,
		Mid:   MidAtlanticOcean,
		Eng:   EnglishChannel,
		Bre:   Brest,
		Pic:   Picardy,
		Bel:   Belgium,
		Ruh:   Ruhr,
		Mun:   Munich,
		Sil:   Silesia,
		War:   Warsaw,
		Ukr:   Ukraine,
		Sev:   Sevastopol,
		Par:   Paris,
		Bur:   Burgundy,
		Boh:   Bohemia,
		Gal:   Galicia,
		Gas:   Gascony,
		Mar:   Marielles,
		Pie:   Piedmont,
		Ven:   Venice,
		Adr:   AdriaticSea,
		Tyr:   Tyrolia,
		Tri:   Trieste,
		Vie:   Vienna,
		Bud:   Budapest,
		Rum:   Rumania,
		Por:   Portugal,
		Spa:   Spain,
		GoL:   GulfOfLyon,
		Tus:   Tuscany,
		Rom:   Rome,
		Apu:   Apulia,
		Alb:   Albania,
		Ser:   Serbia,
		Bul:   Bulgaria,
		Bla:   BlackSea,
		Wes:   WesternMeditraneean,
		Tyn:   TyrrhenianSea,
		Nap:   Naples,
		Ion:   IonianSea,
		Gre:   Greace,
		Aeg:   AegeanSea,
		Con:   Constantinople,
		Smy:   Smyrna,
		Ank:   Ankara,
		Arm:   Armenia,
		NAf:   NorthAfrica,
		Tun:   Tunisa,
		Eas:   EasternMeditraneean,
		Syr:   Syria,
	}
}

// NorthAtlantic Sea, a netural water province.
var NorthAtlantic = Province{
	Name:            "North Atlantic Ocean",
	ShortName:       NAt,
	HasSupplyCenter: false,
	Type:            Water,
	Neighbours:      []ShortName{Nrg, Cly, Lvp, Iri, Mid},
}

// Norwegian Sea, a netural water province.
var NorwegianSea = Province{
	Name:            "Norwegian Sea",
	ShortName:       Nrg,
	HasSupplyCenter: false,
	Type:            Water,
	Neighbours:      []ShortName{Bar, Nwy, Nth, Edi, Cly, NAt},
}

// Barnets Sea, a netural water province.
var BarnetsSea = Province{
	Name:            "Barnets Sea",
	ShortName:       Bar,
	HasSupplyCenter: false,
	Type:            Water,
	Neighbours:      []ShortName{StP, Nwy, Nrg},
}

// StPetersburg, a Russian costal province.
var StPetersburg = Province{
	Name:            "St. Petersburg",
	ShortName:       StP,
	HasSupplyCenter: true,
	Type:            Costal,
	Neighbours:      []ShortName{Mos, Lvn, Bot, Fin, Nwy, Bar},
	ChildProvinces:  []ShortName{StPNC, StPSC},
}

// St. Petersburg North Coast, the north coast of the Russian St. Petersburg province.
var StPetersburgNorthCoast = Province{
	Name:            "St. Petersburg North Coast",
	ShortName:       StPNC,
	HasSupplyCenter: false,
	Type:            Costal,
	Neighbours:      []ShortName{Nwy, Bar},
	Parent:          StP,
}

// St. Petersburg South Coast, the south coast of the Russian St. Petersburg province.
var StPetersburgSouthCoast = Province{
	Name:            "St. Petersburg South Coast",
	ShortName:       StPSC,
	HasSupplyCenter: false,
	Type:            Costal,
	Neighbours:      []ShortName{Fin, Lvn, Bot},
	Parent:          StP,
}

// Clyde, an English costal province.
var Clyde = Province{
	Name:            "Clyde",
	ShortName:       Cly,
	HasSupplyCenter: false,
	Type:            Costal,
	Neighbours:      []ShortName{NAt, Nrg, Edi, Lvp},
}

// Edinburgh, an English costal province.
var Edinburgh = Province{
	Name:            "Edinburgh",
	ShortName:       Edi,
	HasSupplyCenter: true,
	Type:            Costal,
	Neighbours:      []ShortName{Nrg, Nth, Yor, Lvp, Cly},
}

// North Sea, a netural water province.
var NorthSea = Province{
	Name:            "North Sea",
	ShortName:       Nth,
	HasSupplyCenter: false,
	Type:            Water,
	Neighbours:      []ShortName{Nrg, Nwy, Ska, Den, Hel, Hol, Bel, Eng, Lon, Yor, Edi},
}

// Norway, a netural costal province.
var Norway = Province{
	Name:            "Norway",
	ShortName:       Nwy,
	HasSupplyCenter: true,
	Type:            Costal,
	Neighbours:      []ShortName{Nrg, Bar, StP, Fin, Swe, Ska, Nth},
}

// Sweden, a neutral costal province.
var Sweden = Province{
	Name:            "Sweden",
	ShortName:       Swe,
	HasSupplyCenter: true,
	Type:            Costal,
	Neighbours:      []ShortName{Nwy, Fin, Bot, Bal, Den, Ska},
}

// Gulf of Botnia, a netural water province.
var GulfOfBotnia = Province{
	Name:            "Gulf of Botnia",
	ShortName:       Bot,
	HasSupplyCenter: false,
	Type:            Water,
	Neighbours:      []ShortName{Swe, Fin, StP, Lvn, Bal},
}

// Finland, a Russian costal province
var Finland = Province{
	Name:            "Finland",
	ShortName:       Fin,
	HasSupplyCenter: false,
	Type:            Costal,
	Neighbours:      []ShortName{Nwy, StP, Bot, Swe},
}

// Irish sea, a neutral water province.
var IrishSea = Province{
	Name:            "Irish Sea",
	ShortName:       Iri,
	HasSupplyCenter: false,
	Type:            Water,
	Neighbours:      []ShortName{NAt, Lvp, Wal, Eng, Mid},
}

// Liverpool, an English costal province.
var Liverpool = Province{
	Name:            "Liverpool",
	ShortName:       Lvp,
	HasSupplyCenter: true,
	Type:            Costal,
	Neighbours:      []ShortName{Cly, Edi, Yor, Wal, Iri, NAt},
}

// Yorkshire, an English costal province.
var Yorkshire = Province{
	Name:            "Yorkshire",
	ShortName:       Yor,
	HasSupplyCenter: false,
	Type:            Costal,
	Neighbours:      []ShortName{Edi, Nth, Lon, Wal, Lvp},
}

// Skagerrack, a netural water province.
var Skagerrack = Province{
	Name:            "Skagerrack",
	ShortName:       "Ska",
	HasSupplyCenter: false,
	Type:            Water,
	Neighbours:      []ShortName{Nwy, Swe, Den, Nth},
}

// Denmark, a netural costal province.
var Denmark = Province{
	Name:            "Denmark",
	ShortName:       Den,
	HasSupplyCenter: true,
	Type:            Costal,
	Neighbours:      []ShortName{Ska, Swe, Bal, Kie, Hel, Nth},
}

// Baltic Sea, a netural water province.
var BalticSea = Province{
	Name:            "Baltic Sea",
	ShortName:       Bal,
	HasSupplyCenter: false,
	Type:            Water,
	Neighbours:      []ShortName{Swe, Bot, Lvn, Pru, Ber, Kie, Den},
}

// Livonia, a Russian costal province.
var Livonia = Province{
	Name:            "Livonia",
	ShortName:       Lvn,
	HasSupplyCenter: false,
	Type:            Costal,
	Neighbours:      []ShortName{Bot, StP, Mos, War, Pru, Bal},
}

// Wales, an English costal province.
var Wales = Province{
	Name:            "Wales",
	ShortName:       Wal,
	HasSupplyCenter: false,
	Type:            Costal,
	Neighbours:      []ShortName{Lvp, Yor, Lon, Eng, Iri},
}

// London, an English costal province.
var London = Province{
	Name:            "London",
	ShortName:       Lon,
	HasSupplyCenter: true,
	Type:            Costal,
	Neighbours:      []ShortName{Yor, Nth, Lon, Eng, Wal},
}

// Holland, a netural costal province.
var Holland = Province{
	Name:            "Holland",
	ShortName:       Hol,
	HasSupplyCenter: true,
	Type:            Costal,
	Neighbours:      []ShortName{Nth, Hel, Kie, Ruh, Bel},
}

// Heligoland Bight, a neutral water province.
var HeligolandBight = Province{
	Name:            "HeligolandBight",
	ShortName:       Hel,
	HasSupplyCenter: false,
	Type:            Water,
	Neighbours:      []ShortName{Nth, Den, Kie, Hol},
}

// Kiel, German costal province.
var Kiel = Province{
	Name:            "Kiel",
	ShortName:       Kie,
	HasSupplyCenter: true,
	Type:            Costal,
	Neighbours:      []ShortName{Den, Bal, Ber, Mun, Ruh, Hol, Hel},
}

// Berlin, a German costal province.
var Berlin = Province{
	Name:            "Berlin",
	ShortName:       Ber,
	HasSupplyCenter: true,
	Type:            Costal,
	Neighbours:      []ShortName{Bal, Pru, Sil, Mun, Kie},
}

// Prussia, a German costal province.
var Prussia = Province{
	Name:            "Prussia",
	ShortName:       Pru,
	HasSupplyCenter: false,
	Type:            Costal,
	Neighbours:      []ShortName{Bal, Lvn, War, Sil, Ber},
}

// Moscow, a Russian land province.
var Moscow = Province{
	Name:            "Moscow",
	ShortName:       Mos,
	HasSupplyCenter: true,
	Type:            Land,
	Neighbours:      []ShortName{StP, Sev, Ukr, War, Lvn},
}

// Mid-Atlantic Ocean, a netural water province.
var MidAtlanticOcean = Province{
	Name:            "Mid-Atlantic Ocean",
	ShortName:       Mid,
	HasSupplyCenter: false,
	Type:            Water,
	Neighbours:      []ShortName{NAt, Iri, Eng, Bre, Gas, Spa, Por, NAf},
}

// English Channel, a netural water province.
var EnglishChannel = Province{
	Name:            "English Channel",
	ShortName:       Eng,
	HasSupplyCenter: false,
	Type:            Water,
	Neighbours:      []ShortName{Wal, Lon, Nth, Bel, Pic, Bre, Mid, Iri},
}

// Brest, a French costal province.
var Brest = Province{
	Name:            "Brest",
	ShortName:       Bre,
	HasSupplyCenter: true,
	Type:            Costal,
	Neighbours:      []ShortName{Eng, Pic, Par, Gas, Mid},
}

// Picardy, a French costal province.
var Picardy = Province{
	Name:            "Picardy",
	ShortName:       Pic,
	HasSupplyCenter: false,
	Type:            Costal,
	Neighbours:      []ShortName{Eng, Bel, Bur, Par, Bre},
}

// Belgium, a neutral costal province.
var Belgium = Province{
	Name:            "Belgium",
	ShortName:       Bel,
	HasSupplyCenter: true,
	Type:            Costal,
	Neighbours:      []ShortName{Nth, Hol, Ruh, Bur, Pic, Eng},
}

// Ruhr, a German land province.
var Ruhr = Province{
	Name:            "Ruhr",
	ShortName:       Ruh,
	HasSupplyCenter: false,
	Type:            Land,
	Neighbours:      []ShortName{Kie, Mun, Bur, Bel, Hol},
}

// Munich, a German land province.
var Munich = Province{
	Name:            "Munich",
	ShortName:       Mun,
	HasSupplyCenter: true,
	Type:            Land,
	Neighbours:      []ShortName{Ber, Sil, Boh, Tyr, Bur, Ruh, Kie},
}

// Silesia, a German land province.
var Silesia = Province{
	Name:            "Silesia",
	ShortName:       Sil,
	HasSupplyCenter: false,
	Type:            Land,
	Neighbours:      []ShortName{Pru, War, Gal, Boh, Mun, Ber},
}

// Warsaw, a Russian land province.
var Warsaw = Province{
	Name:            "Warsaw",
	ShortName:       War,
	HasSupplyCenter: true,
	Type:            Land,
	Neighbours:      []ShortName{Pru, Lvn, Mos, Ukr, Gal, Sil},
}

// Ukraine, a Russian land province.
var Ukraine = Province{
	Name:            "Ukraine",
	ShortName:       Ukr,
	HasSupplyCenter: false,
	Type:            Land,
	Neighbours:      []ShortName{Mos, Sev, Rum, Gal, War},
}

// Sevastopol, a Russian costal province.
var Sevastopol = Province{
	Name:            "Sevastopol",
	ShortName:       Sev,
	HasSupplyCenter: true,
	Type:            Costal,
	Neighbours:      []ShortName{Mos, Arm, Bla, Rum, Ukr},
}

// Paris, a French land province.
var Paris = Province{
	Name:            "Paris",
	ShortName:       "Par",
	HasSupplyCenter: true,
	Type:            Land,
	Neighbours:      []ShortName{Pic, Bur, Gas, Bre},
}

// Burgundy, a French land province.
var Burgundy = Province{
	Name:            "Burgundy",
	ShortName:       Bur,
	HasSupplyCenter: false,
	Type:            Land,
	Neighbours:      []ShortName{Bel, Ruh, Mun, Mar, Gas, Par, Pic},
}

// Bohemia, an Austrian land province.
var Bohemia = Province{
	Name:            "Bohemia",
	ShortName:       Boh,
	HasSupplyCenter: false,
	Type:            Land,
	Neighbours:      []ShortName{Sil, Gal, Vie, Boh, Tyr, Mun},
}

// Galicia, an Austrian land province.
var Galicia = Province{
	Name:            "Galicia",
	ShortName:       Gal,
	HasSupplyCenter: false,
	Type:            Land,
	Neighbours:      []ShortName{Ukr, Rum, Bud, Vie, Boh, Sil, War},
}

// Gascony, a French costal province.
var Gascony = Province{
	Name:            "Gascony",
	ShortName:       Gas,
	HasSupplyCenter: false,
	Type:            Costal,
	Neighbours:      []ShortName{Par, Bur, Mar, Spa, Mid, Bre},
}

// Marielles, a French costal province.
var Marielles = Province{
	Name:            "Marielles",
	ShortName:       Mar,
	HasSupplyCenter: true,
	Type:            Costal,
	Neighbours:      []ShortName{Bur, Pie, GoL, Spa, Gas},
}

// Piedmont, an Italian costal province.
var Piedmont = Province{
	Name:            "Piedmont",
	ShortName:       Pie,
	HasSupplyCenter: false,
	Type:            Costal,
	Neighbours:      []ShortName{Tyr, Ven, Tus, GoL, Mar},
}

// Venice, an Italan costal province.
var Venice = Province{
	Name:            "Venice",
	ShortName:       Ven,
	HasSupplyCenter: true,
	Type:            Costal,
	Neighbours:      []ShortName{Tyr, Tri, Adr, Apu, Rom, Tus, Pie},
}

// Adriatic Sea, a neutral water province.
var AdriaticSea = Province{
	Name:            "Adriatic Sea",
	ShortName:       Adr,
	HasSupplyCenter: false,
	Type:            Water,
	Neighbours:      []ShortName{Tri, Alb, Ion, Apu, Ven},
}

// Tyrolia, an Austrian land province.
var Tyrolia = Province{
	Name:            "Tyrolia",
	ShortName:       Tyr,
	HasSupplyCenter: false,
	Type:            Land,
	Neighbours:      []ShortName{Mun, Boh, Vie, Tri, Ven, Pie},
}

// Trieste, an Austrian costal province.
var Trieste = Province{
	Name:            "Trieste",
	ShortName:       Tri,
	HasSupplyCenter: true,
	Type:            Costal,
	Neighbours:      []ShortName{Vie, Bud, Ser, Alb, Adr, Ven, Tyr},
}

// Vienna, an Austrian land province.
var Vienna = Province{
	Name:            "Vienna",
	ShortName:       Vie,
	HasSupplyCenter: true,
	Type:            Land,
	Neighbours:      []ShortName{Boh, Gal, Bud, Tri, Tyr},
}

// Budapest, an Austrian land province.
var Budapest = Province{
	Name:            "Budapest",
	ShortName:       Bud,
	HasSupplyCenter: true,
	Type:            Land,
	Neighbours:      []ShortName{Gal, Rum, Ser, Tri, Vie},
}

// Rumania, a netural costal province.
var Rumania = Province{
	Name:            "Rumania",
	ShortName:       Rum,
	HasSupplyCenter: true,
	Type:            Costal,
	Neighbours:      []ShortName{Ukr, Sev, Bla, Bul, Ser, Bud, Gal},
}

// Portugal, a netural costal province.
var Portugal = Province{
	Name:            "Portugal",
	ShortName:       Por,
	HasSupplyCenter: true,
	Type:            Costal,
	Neighbours:      []ShortName{Spa, Mid},
}

// Spain, a netural costal province.
var Spain = Province{
	Name:            "Spain",
	ShortName:       Spa,
	HasSupplyCenter: true,
	Type:            Costal,
	Neighbours:      []ShortName{Mid, Gas, Mar, GoL, Wes, Por},
	ChildProvinces:  []ShortName{SpaNC, SpaSC},
}

// Spain North Coast, the north coast of the neutral Spain province.
var SpainNorthCoast = Province{
	Name:            "Spain North Coast",
	ShortName:       SpaNC,
	HasSupplyCenter: false,
	Type:            Costal,
	Neighbours:      []ShortName{Mid, Gas, Por},
	Parent:          Spa,
}

// Spain South Coast, the south coast of the neutral Spain province.
var SpainSouthCoast = Province{
	Name:            "Spain South Coast",
	ShortName:       SpaSC,
	HasSupplyCenter: false,
	Type:            Costal,
	Neighbours:      []ShortName{Mar, Wes, Por, Mid},
	Parent:          Spa,
}

// Gulf of Lyon, a neutral water province.
var GulfOfLyon = Province{
	Name:            "Gulf of Lyon",
	ShortName:       GoL,
	HasSupplyCenter: false,
	Type:            Water,
	Neighbours:      []ShortName{Mar, Pie, Tus, Tyn, Wes, Spa},
}

// Tuscany, an Italian costal province.
var Tuscany = Province{
	Name:            "Tuscany",
	ShortName:       Tus,
	HasSupplyCenter: false,
	Type:            Costal,
	Neighbours:      []ShortName{Ven, Rom, Tyn, GoL, Pie},
}

// Rome, an Italian costal province.
var Rome = Province{
	Name:            "Rome",
	ShortName:       Rom,
	HasSupplyCenter: true,
	Type:            Costal,
	Neighbours:      []ShortName{Ven, Apu, Nap, Tyn, Tus},
}

// Apulia, an Italian costal province.
var Apulia = Province{
	Name:            "Apulia",
	ShortName:       Apu,
	HasSupplyCenter: false,
	Type:            Costal,
	Neighbours:      []ShortName{Adr, Ion, Nap, Rom, Ven},
}

// Albania, a neutral costal province.
var Albania = Province{
	Name:            "Albania",
	ShortName:       Alb,
	HasSupplyCenter: false,
	Type:            Costal,
	Neighbours:      []ShortName{Tri, Ser, Gre, Ion, Adr},
}

// Serbia, a neutral land province.
var Serbia = Province{
	Name:            "Serbia",
	ShortName:       Ser,
	HasSupplyCenter: true,
	Type:            Land,
	Neighbours:      []ShortName{Bud, Rum, Bul, Gre, Alb, Tri},
}

// Bulgaria, a netural costal province.
var Bulgaria = Province{
	Name:            "Bulgaria",
	ShortName:       Bul,
	HasSupplyCenter: true,
	Type:            Costal,
	Neighbours:      []ShortName{Rum, Bla, Con, Aeg, Gre, Ser},
	ChildProvinces:  []ShortName{BulEC, BulSC},
}

// Bulgaria East Coast, the east coast of the neutral province Bulgaria.
var BulgariaEastCoast = Province{
	Name:            "Bulgaria East Coast",
	ShortName:       BulEC,
	HasSupplyCenter: false,
	Type:            Costal,
	Neighbours:      []ShortName{Rum, Bla, Con},
	Parent:          Bul,
}

// Bulgaria South Coast, the south coast of the neutral province Bulgaria.
var BulgariaSouthCoast = Province{
	Name:            "Bulgaria South Coast",
	ShortName:       BulSC,
	HasSupplyCenter: false,
	Type:            Costal,
	Neighbours:      []ShortName{Con, Aeg, Gre},
	Parent:          Bul,
}

// Black Sea, a neutral water province.
var BlackSea = Province{
	Name:            "Black Sea",
	ShortName:       Bla,
	HasSupplyCenter: false,
	Type:            Water,
	Neighbours:      []ShortName{Sev, Arm, Ank, Con, Bul, Rum},
}

// Western Meditraneean, a neutral water province.
var WesternMeditraneean = Province{
	Name:            "Western Meditraneean",
	ShortName:       Wes,
	HasSupplyCenter: false,
	Type:            Water,
	Neighbours:      []ShortName{GoL, Tyn, Tun, NAf, Spa},
}

// Tyrrhenian Sea, a neutral water province.
var TyrrhenianSea = Province{
	Name:            "Tyrrhenian Sea",
	ShortName:       Tyn,
	HasSupplyCenter: false,
	Type:            Water,
	Neighbours:      []ShortName{Tus, Rom, Nap, Ion, Tun, Wes, GoL},
}

// Naples, an Italian costal province.
var Naples = Province{
	Name:            "Naples",
	ShortName:       Nap,
	HasSupplyCenter: true,
	Type:            Costal,
	Neighbours:      []ShortName{Apu, Ion, Tyn, Rom},
}

// Ionian Sea, a neutral water province.
var IonianSea = Province{
	Name:            "Ionian Sea",
	ShortName:       Ion,
	HasSupplyCenter: false,
	Type:            Water,
	Neighbours:      []ShortName{Apu, Adr, Alb, Gre, Aeg, Eas, Tun, Tyn, Nap},
}

// Greace, a neutral costal province.
var Greace = Province{
	Name:            "Greace",
	ShortName:       Gre,
	HasSupplyCenter: true,
	Type:            Costal,
	Neighbours:      []ShortName{Ser, Bul, Aeg, Ion, Alb},
}

// Aegean Sea, a netural water province.
var AegeanSea = Province{
	Name:            "Aegean Sea",
	ShortName:       Aeg,
	HasSupplyCenter: false,
	Type:            Water,
	Neighbours:      []ShortName{Bul, Con, Smy, Eas, Ion, Gre},
}

// Constantinople, a Turkish costal province.
var Constantinople = Province{
	Name:            "Constantinople",
	ShortName:       Con,
	HasSupplyCenter: true,
	Type:            Costal,
	Neighbours:      []ShortName{Bla, Ank, Smy, Aeg, Bul},
}

// Smyrna, a Turkish costal province.
var Smyrna = Province{
	Name:            "Smyrna",
	ShortName:       Smy,
	HasSupplyCenter: true,
	Type:            Costal,
	Neighbours:      []ShortName{Ank, Arm, Syr, Eas, Aeg, Con},
}

// Ankara, a Turkish costal province.
var Ankara = Province{
	Name:            "Ankara",
	ShortName:       Ank,
	HasSupplyCenter: true,
	Type:            Costal,
	Neighbours:      []ShortName{Bla, Arm, Smy, Con},
}

// Armenia, a Turkish costal province.
var Armenia = Province{
	Name:            "Armenia",
	ShortName:       Arm,
	HasSupplyCenter: false,
	Type:            Costal,
	Neighbours:      []ShortName{Sev, Syr, Smy, Ank, Bla},
}

// North Africa, a neutral costal province.
var NorthAfrica = Province{
	Name:            "North Africa",
	ShortName:       NAf,
	HasSupplyCenter: false,
	Type:            Costal,
	Neighbours:      []ShortName{Mid, Wes, Tun},
}

// Tunisa, a neutral costal province.
var Tunisa = Province{
	Name:            "Tunisa",
	ShortName:       Tun,
	HasSupplyCenter: true,
	Type:            Costal,
	Neighbours:      []ShortName{Tyn, Ion, NAf, Wes},
}

// Eastern Meditraneean, a netural water province.
var EasternMeditraneean = Province{
	Name:            "Eastern Meditraneean",
	ShortName:       Eas,
	HasSupplyCenter: false,
	Type:            Water,
	Neighbours:      []ShortName{Smy, Syr, Ion, Aeg},
}

// Syria, a Turkish costal province.
var Syria = Province{
	Name:            "Syria",
	ShortName:       Syr,
	HasSupplyCenter: false,
	Type:            Costal,
	Neighbours:      []ShortName{Arm, Eas, Smy},
}
