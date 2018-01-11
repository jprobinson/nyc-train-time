package main

var trainDirs = map[string]map[string]string{
	"1":  {"northbound": "Bronx", "southbound": "South Ferry"},
	"2":  {"northbound": "Bronx", "southbound": "Brooklyn"},
	"3":  {"northbound": "Harlem", "southbound": "Brooklyn"},
	"4":  {"northbound": "Bronx", "southbound": "Brooklyn"},
	"5":  {"northbound": "Bronx", "southbound": "Brooklyn"},
	"5X": {"northbound": "Bronx", "southbound": "Brooklyn"},
	"6":  {"northbound": "Bronx", "southbound": "Brooklyn Bridge"},
	"6X": {"northbound": "Bronx", "southbound": "Brooklyn Bridge"},
	"S":  {"northbound": "", "southbound": ""},
	"L":  {"northbound": "Manhattan", "southbound": "Brooklyn"},
	"B":  {"northbound": "Bronx", "southbound": "Brooklyn"},
	"D":  {"northbound": "Bronx", "southbound": "Brooklyn"},
	"A":  {"northbound": "Manhattan", "southbound": "Queens"},
	"G":  {"northbound": "Queens", "southbound": "Brooklyn"},
	"C":  {"northbound": "Manhattan", "southbound": "Brooklyn"},
	"E":  {"northbound": "Queens", "southbound": "Manhattan"},
	"N":  {"northbound": "Manhattan", "southbound": "Brooklyn"},
	"Q":  {"northbound": "Manhattan", "southbound": "Brooklyn"},
	"R":  {"northbound": "Queens", "southbound": "Brooklyn"},
	"W":  {"northbound": "Queens", "southbound": "Manhattan"},
	"J":  {"northbound": "Queens", "southbound": "Manhattan"},
	"F":  {"northbound": "Queens", "southbound": "Brooklyn"},
	"M":  {"northbound": "Queens", "southbound": "Brooklyn"},
	"Z":  {"northbound": "Queens", "southbound": "Manhattan"},
}

// stop name> train line> stop ID
var stopNameToID = map[string]map[string]string{
	"Fresh Pond Road": {
		"M": "M04",
	},
	"Central Park North ,110th Street": {
		"3": "227",
		"2": "227",
		"5": "227",
	},
	"231st Street": {
		"1": "104",
	},
	"111st Street": {
		"A": "A64",
		"J": "J13",
	},
	"182nd-183rd Streets": {
		"D": "D06",
	},
	"Avenue N": {
		"F": "F33",
	},
	"Hewes Street": {
		"Z": "M14",
		"J": "M14",
	},
	"Hunts Point Avenue": {
		"6X": "613",
		"6":  "613",
	},
	"Clinton Washington Avenues": {
		"C": "A44",
		"G": "G35",
	},
	"Steinway Street": {
		"R": "G19",
		"E": "G19",
	},
	"169th Street": {
		"F": "F02",
	},
	"34th Street Penn Station": {
		"A": "A28",
		"C": "A28",
		"E": "A28",
		"1": "128",
		"3": "128",
		"2": "128",
		"5": "128",
	},
	"Bedford Park Blvd Lehman College": {
		"4": "405",
	},
	"Saint Lawrence Avenue": {
		"6X": "609",
		"6":  "609",
	},
	"Burnside Avenue": {
		"4": "409",
	},
	"Ralph Avenue": {
		"C": "A49",
	},
	"Union Street": {
		"R": "R32",
	},
	"Jay Street MetroTech": {
		"A": "A41",
		"C": "A41",
		"R": "R29",
		"F": "A41",
	},
	"Queensboro Plaza": {
		"W": "R09",
		"N": "R09",
	},
	"Avenue M": {
		"Q": "D34",
	},
	"Prince Street": {
		"Q": "R22",
		"R": "R22",
		"W": "R22",
	},
	"168th Street Washington Heights": {
		"1": "112",
	},
	"Halsey Street": {
		"Z": "J29",
		"J": "J29",
		"L": "L19",
	},
	"Briarwood Van Wyck Blvd": {
		"E": "F05",
		"F": "F05",
	},
	"86th Street": {
		"Q":  "Q04",
		"C":  "A20",
		"B":  "A20",
		"N":  "N10",
		"1":  "121",
		"2":  "121",
		"5":  "121",
		"4":  "626",
		"6":  "626",
		"6X": "626",
		"R":  "R44",
	},
	"207th Street": {
		"1": "108",
	},
	"Woodhaven Blvd": {
		"Z": "J15",
		"R": "G11",
		"J": "J15",
		"E": "G11",
	},
	"55th Street": {
		"D": "B15",
	},
	"42nd Street Port Authority Bus Terminal": {
		"A": "A27",
		"C": "A27",
		"E": "A27",
	},
	"Atlantic Av Barclays Ctr": {
		"B": "D24",
		"D": "R31",
		"N": "R31",
		"Q": "D24",
		"3": "235",
		"2": "235",
		"4": "235",
		"R": "R31",
	},
	"Borough Hall": {
		"3": "232",
		"2": "232",
		"4": "423",
	},
	"Northern Blvd": {
		"R": "G16",
		"E": "G16",
	},
	"39th Avenue": {
		"W": "R08",
		"N": "R08",
	},
	"68th Street Hunter College": {
		"6X": "628",
		"6":  "628",
	},
	"Grand Army Plaza": {
		"3": "237",
		"2": "237",
		"4": "237",
	},
	"103rd Street": {
		"1":  "119",
		"6X": "624",
		"C":  "A18",
		"B":  "A18",
		"6":  "624",
	},
	"63 Dr Rego Park": {
		"R": "G10",
		"E": "G10",
	},
	"Jamaica 179th Street": {
		"F": "F01",
	},
	"Nassau Avenue": {
		"G": "G28",
	},
	"Graham Avenue": {
		"L": "L11",
	},
	"Wilson Avenue": {
		"L": "L20",
	},
	"77th Street": {
		"6X": "627",
		"R":  "R43",
		"6":  "627",
	},
	"Broadway Jct": {
		"A": "A51",
		"C": "A51",
		"J": "J27",
		"Z": "J27",
		"L": "L22",
	},
	"Parsons Blvd": {
		"F": "F03",
	},
	"59th Street Columbus Circle": {
		"A": "A24",
		"C": "A24",
		"B": "A24",
		"D": "A24",
		"1": "125",
		"2": "125",
		"5": "125",
	},
	"Prospect Park": {
		"Q": "D26",
		"B": "D26",
	},
	"23rd Street": {
		"Q":  "R19",
		"C":  "A30",
		"E":  "A30",
		"F":  "D18",
		"1":  "130",
		"2":  "130",
		"5":  "130",
		"W":  "R19",
		"6":  "634",
		"6X": "634",
		"R":  "R19",
	},
	"36th Street": {
		"R": "R36",
		"E": "G20",
		"D": "R36",
		"N": "R36",
	},
	"Wall Street": {
		"3": "230",
		"2": "230",
		"4": "419",
	},
	"Tremont Avenue": {
		"D": "D07",
	},
	"Freeman Street": {
		"2": "216",
		"5": "216",
	},
	"Jackson Avenue": {
		"2": "220",
		"5": "220",
	},
	"World Trade Center": {
		"E": "E01",
	},
	"18th Avenue": {
		"N": "N05",
		"D": "B19",
		"F": "F30",
	},
	"135th Street": {
		"3": "224",
		"2": "224",
		"5": "224",
		"C": "A14",
		"B": "A14",
	},
	"59th Street": {
		"6X": "629",
		"R":  "R41",
		"N":  "R41",
		"4":  "629",
		"6":  "629",
	},
	"Cortelyou Road": {
		"Q": "D30",
	},
	"Metropolitan Avenue": {
		"G": "G29",
	},
	"57th Street": {
		"F": "B10",
	},
	"104th Street": {
		"A": "A63",
		"Z": "J14",
		"J": "J14",
	},
	"Union Square 14th Street": {
		"L": "L03",
	},
	"Bedford Park Blvd": {
		"D": "D03",
	},
	"Court Street": {
		"R": "R28",
	},
	"New Lots Avenue": {
		"3": "257",
		"4": "257",
		"L": "L27",
	},
	"East 149th Street": {
		"6": "615",
	},
	"145th Street": {
		"A": "A12",
		"C": "A12",
		"B": "D13",
		"D": "D13",
		"1": "114",
		"3": "302",
	},
	"Westchester Square East Tremont Avenue": {
		"6X": "604",
		"6":  "604",
	},
	"Rockaway Blvd": {
		"A": "A61",
	},
	"25th Avenue": {
		"D": "B22",
	},
	"57th Street 7th Avenue": {
		"Q": "R14",
		"R": "R14",
		"W": "R14",
		"N": "R14",
	},
	"President Street": {
		"2": "241",
	},
	"174th Street": {
		"2": "215",
		"5": "215",
	},
	"Fordham Road": {
		"4": "407",
		"D": "D05",
	},
	"34th Street Herald Square": {
		"B": "D17",
		"D": "D17",
		"F": "D17",
		"N": "R17",
		"Q": "R17",
		"R": "R17",
		"W": "R17",
	},
	"138th Street Grand Concourse": {
		"4": "416",
	},
	"Bedford Nostrand Avenues": {
		"G": "G33",
	},
	"Whitehall Street": {
		"R": "R27",
		"W": "R27",
	},
	"Intervale Avenue": {
		"2": "218",
		"5": "218",
	},
	"Lexington Av/59th Street": {
		"R": "R11",
		"W": "R11",
		"N": "R11",
	},
	"Astoria Ditmars Blvd": {
		"W": "R01",
		"N": "R01",
	},
	"Morrison Av- Sound View": {
		"6X": "610",
		"6":  "610",
	},
	"Seneca Avenue": {
		"M": "M06",
	},
	"121st Street": {
		"Z": "J12",
		"J": "J12",
	},
	"Simpson Street": {
		"2": "217",
		"5": "217",
	},
	"Wakefield 241st Street": {
		"2": "201",
	},
	"Hoyt Street": {
		"3": "233",
		"2": "233",
	},
	"Grand Street": {
		"B": "D22",
		"D": "D22",
		"L": "L12",
	},
	"Woodlawn": {
		"4": "401",
	},
	"Avenue H": {
		"Q": "D32",
	},
	"110th Street": {
		"6X": "623",
		"6":  "623",
	},
	"Parkside Avenue": {
		"Q": "D27",
	},
	"Houston Street": {
		"1": "134",
		"2": "134",
		"5": "134",
	},
	"96th Street": {
		"Q":  "Q05",
		"C":  "A19",
		"B":  "A19",
		"1":  "120",
		"3":  "120",
		"2":  "120",
		"5":  "120",
		"6":  "625",
		"6X": "625",
	},
	"East 143rd Street Saint Mary's Street": {
		"6": "616",
	},
	"20th Avenue": {
		"D": "B20",
		"N": "N06",
	},
	"Livonia Avenue": {
		"L": "L26",
	},
	"62nd Street": {
		"D": "B16",
	},
	"Myrtle Willoughby Avenues": {
		"G": "G32",
	},
	"Sutter Avenue": {
		"L": "L25",
	},
	"Junius Street": {
		"3": "254",
		"4": "254",
	},
	"Franklin Street": {
		"1": "136",
		"2": "136",
		"5": "136",
	},
	"Saratoga Avenue": {
		"3": "252",
		"4": "252",
	},
	"72nd Street": {
		"Q": "Q03",
		"C": "A22",
		"B": "A22",
		"1": "123",
		"3": "123",
		"2": "123",
		"5": "123",
	},
	"Harlem 148th Street": {
		"3": "301",
	},
	"Van Siclen Avenue": {
		"C": "A53",
		"3": "256",
		"J": "J23",
		"Z": "J23",
		"4": "256",
	},
	"Clark Street": {
		"3": "231",
		"2": "231",
	},
	"Kosciuszko Street": {
		"Z": "J31",
		"J": "J31",
	},
	"67th Avenue": {
		"R": "G09",
		"E": "G09",
	},
	"Inwood 207th Street": {
		"A": "A02",
	},
	"Beverley Road": {
		"Q": "D29",
	},
	"Liberty Avenue": {
		"C": "A52",
	},
	"Bronx Park East": {
		"2": "212",
	},
	"Sterling Street": {
		"2": "242",
	},
	"219th Street": {
		"2": "207",
	},
	"174nd-175nd Streets": {
		"D": "D08",
	},
	"5 Av/53rd Street": {
		"E": "F12",
	},
	"West 4th Street": {
		"A": "A32",
		"C": "A32",
		"B": "D20",
		"E": "A32",
		"D": "D20",
		"F": "D20",
	},
	"Nevins Street": {
		"3": "234",
		"2": "234",
		"4": "234",
	},
	"149th Street Grand Concourse": {
		"2": "222",
		"5": "222",
		"4": "415",
	},
	"Bay Ridge Avenue": {
		"R": "R42",
	},
	"Sutphin Blvd": {
		"F": "F04",
	},
	"Gun Hill Road": {
		"2": "208",
		"5": "503",
	},
	"Ocean Parkway": {
		"Q": "D41",
	},
	"York Street": {
		"F": "F18",
	},
	"Myrtle Wyckoff Avenues": {
		"M": "M08",
		"L": "L17",
	},
	"Court Square": {
		"E": "F09",
		"G": "G22",
	},
	"49th Street": {
		"Q": "R15",
		"R": "R15",
		"W": "R15",
		"N": "R15",
	},
	"Nereid Avenue": {
		"2": "204",
	},
	"High Street": {
		"A": "A40",
		"C": "A40",
	},
	"Broad Street": {
		"Z": "M23",
		"J": "M23",
	},
	"Burke Avenue": {
		"2": "209",
	},
	"3rd Avenue": {
		"L": "L05",
	},
	"Elmhurst Avenue": {
		"R": "G13",
		"E": "G13",
	},
	"Carroll Street": {
		"G": "F21",
		"F": "F21",
	},
	"Broadway-Lafayette Street": {
		"B": "D21",
		"D": "D21",
		"F": "D21",
	},
	"50th Street": {
		"C": "A25",
		"E": "A25",
		"D": "B14",
		"1": "126",
		"2": "126",
		"5": "126",
	},
	"Winthrop Street": {
		"2": "243",
	},
	"Bushwick Av Aberdeen Street": {
		"L": "L21",
	},
	"9th Avenue": {
		"D": "B12",
	},
	"Smith 9th Streets": {
		"G": "F22",
		"F": "F22",
	},
	"Times Square 42nd Street": {
		"Q": "R16",
		"N": "R16",
		"1": "127",
		"3": "127",
		"2": "127",
		"5": "127",
		"W": "R16",
		"R": "R16",
	},
	"Broadway": {
		"W": "R05",
		"G": "G30",
		"N": "R05",
	},
	"Flatbush Av Brooklyn College": {
		"2": "247",
	},
	"Nostrand Avenue": {
		"A": "A46",
		"C": "A46",
		"3": "248",
		"4": "248",
	},
	"21st Street": {
		"G": "G24",
	},
	"1st Avenue": {
		"L": "L06",
	},
	"Pelham Bay Park": {
		"6X": "601",
		"6":  "601",
	},
	"South Ferry": {
		"1": "142",
		"5": "142",
	},
	"36th Avenue": {
		"W": "R06",
		"N": "R06",
	},
	"Grand Av Newtown": {
		"R": "G12",
		"E": "G12",
	},
	"Brighton Beach": {
		"Q": "D40",
		"B": "D40",
	},
	"Myrtle Avenue": {
		"Z": "M11",
		"J": "M11",
	},
	"Rockefeller Center": {
		"B": "D15",
		"D": "D15",
		"F": "D15",
	},
	"47th-50th Streets Rockefeller Center": {
		"B": "D15",
		"D": "D15",
		"F": "D15",
	},
	"Van Cortlandt Park 242nd Street": {
		"1": "101",
	},
	"42nd Street Bryant Park": {
		"B": "D16",
		"D": "D16",
		"F": "D16",
	},
	"Dyckman Street": {
		"1": "109",
		"A": "A03",
	},
	"Classon Avenue": {
		"G": "G34",
	},
	"4th Avenue": {
		"G": "F23",
		"F": "F23",
	},
	"3 Av 149th Street": {
		"2": "221",
		"5": "221",
	},
	"80th Street": {
		"A": "A59",
	},
	"Parkchester": {
		"6X": "608",
		"6":  "608",
	},
	"Kingston Avenue": {
		"3": "249",
		"4": "249",
	},
	"Sheepshead Bay": {
		"Q": "D39",
		"B": "D39",
	},
	"161 Street Yankee Stadium": {
		"4": "414",
		"D": "D11",
	},
	"6th Avenue": {
		"L": "L02",
	},
	"Bedford Avenue": {
		"L": "L08",
	},
	"Sutter Av Rutland Road": {
		"3": "251",
		"4": "251",
	},
	"Bowery": {
		"Z": "M19",
		"J": "M19",
	},
	"190th Street": {
		"A": "A05",
	},
	"Beverly Road": {
		"2": "245",
	},
	"Greenpoint Avenue": {
		"G": "G26",
	},
	"225th Street": {
		"2": "206",
	},
	"238th Street": {
		"1": "103",
	},
	"15th Street Prospect Park": {
		"G": "F25",
		"F": "F25",
	},
	"21 Street Queensbridge": {
		"F": "B04",
	},
	"Brooklyn Bridge City Hall": {
		"6X": "640",
		"4":  "640",
		"6":  "640",
	},
	"3 Av 138th Street": {
		"6X": "619",
		"6":  "619",
	},
	"Ditmas Avenue": {
		"F": "F29",
	},
	"Montrose Avenue": {
		"L": "L13",
	},
	"137th Street City College": {
		"1": "115",
	},
	"Mt Eden Avenue": {
		"4": "411",
	},
	"Cypress Hills": {
		"J": "J19",
	},
	"Brook Avenue": {
		"6": "618",
	},
	"155th Street": {
		"C": "A11",
		"D": "D12",
	},
	"88th Street": {
		"A": "A60",
	},
	"Chauncey Street": {
		"Z": "J28",
		"J": "J28",
	},
	"116th Street": {
		"C":  "A16",
		"B":  "A16",
		"3":  "226",
		"2":  "226",
		"5":  "226",
		"6":  "622",
		"6X": "622",
	},
	"Bergen Street": {
		"3": "236",
		"2": "236",
		"4": "236",
		"G": "F20",
		"F": "F20",
	},
	"233rd Street": {
		"2": "205",
	},
	"75th Avenue": {
		"E": "F07",
		"F": "F07",
	},
	"Elder Avenue": {
		"6X": "611",
		"6":  "611",
	},
	"81 Street Museum of Natural History": {
		"C": "A21",
		"B": "A21",
	},
	"Lorimer Street": {
		"Z": "M13",
		"J": "M13",
		"L": "L10",
	},
	"Baychester Avenue": {
		"5": "502",
	},
	"Queens Plaza": {
		"R": "G21",
		"E": "G21",
	},
	"Bay Ridge 95th Street": {
		"R": "R45",
	},
	"215th Street": {
		"1": "107",
	},
	"Castle Hill Avenue": {
		"6X": "607",
		"6":  "607",
	},
	"East Broadway": {
		"F": "F16",
	},
	"191st Street": {
		"1": "110",
	},
	"Flushing Avenue": {
		"Z": "M12",
		"J": "M12",
		"G": "G31",
	},
	"Euclid Avenue": {
		"A": "A55",
		"C": "A55",
	},
	"167th Street": {
		"4": "413",
		"D": "D10",
	},
	"28th Street": {
		"Q":  "R18",
		"1":  "129",
		"R":  "R18",
		"5":  "129",
		"W":  "R18",
		"6":  "633",
		"6X": "633",
		"2":  "129",
	},
	"Allerton Avenue": {
		"2": "210",
	},
	"Newkirk Avenue": {
		"2": "246",
	},
	"Delancey Street": {
		"F": "F15",
	},
	"Cathedral Parkway": {
		"1": "118",
	},
	"Cypress Avenue": {
		"6": "617",
	},
	"30th Avenue": {
		"W": "R04",
		"N": "R04",
	},
	"175th Street": {
		"A": "A07",
	},
	"5 Av/59th Street": {
		"R": "R13",
		"W": "R13",
		"N": "R13",
	},
	"168th Street": {
		"A": "A09",
		"C": "A09",
	},
	"183rd Street": {
		"4": "408",
	},
	"45th Street": {
		"R": "R39",
	},
	"Newkirk Plaza": {
		"Q": "D31",
		"B": "D31",
	},
	"Marcy Avenue": {
		"Z": "M16",
		"J": "M16",
	},
	"Astor Pl": {
		"6X": "636",
		"6":  "636",
	},
	"Middle Village Metropolitan Avenue": {
		"M": "M01",
	},
	"Neck Road": {
		"Q": "D38",
	},
	"8th Street NYU": {
		"Q": "R21",
		"R": "R21",
		"W": "R21",
	},
	"71st Street": {
		"D": "B17",
	},
	"Roosevelt Island": {
		"F": "B06",
	},
	"Bleecker Street": {
		"6X": "637",
		"6":  "637",
	},
	"Crown Heights Utica Avenue": {
		"3": "250",
		"4": "250",
	},
	"Cathedral Parkway ,110th Street": {
		"C": "A17",
		"B": "A17",
	},
	"Forest Avenue": {
		"M": "M05",
	},
	"Rector Street": {
		"1": "139",
		"R": "R26",
		"5": "139",
		"W": "R26",
	},
	"Jamaica Center Parsons/Archer": {
		"Z": "G05",
		"J": "G05",
		"E": "G05",
	},
	"Lexington Av/53rd Street": {
		"E": "F11",
	},
	"157th Street": {
		"1": "113",
	},
	"East 180th Street": {
		"2": "213",
		"5": "213",
	},
	"East 105th Street": {
		"L": "L28",
	},
	"Whitlock Avenue": {
		"6X": "612",
		"6":  "612",
	},
	"Bowling Green": {
		"4": "420",
	},
	"Jamaica Van Wyck": {
		"E": "G07",
	},
	"DeKalb Avenue": {
		"Q": "R30",
		"R": "R30",
		"B": "R30",
		"L": "L16",
	},
	"Avenue U": {
		"Q": "D37",
		"N": "N09",
		"F": "F36",
	},
	"Fort Hamilton Parkway": {
		"N": "N03",
		"D": "B13",
		"G": "F26",
		"F": "F26",
	},
	"Grand Central 42nd Street": {
		"6X": "631",
		"4":  "631",
		"6":  "631",
	},
	"Pelham Parkway": {
		"2": "211",
		"5": "504",
	},
	"Bay 50th Street": {
		"D": "B23",
	},
	"Canarsie Rockaway Parkway": {
		"L": "L29",
	},
	"Norwood 205th Street": {
		"D": "D01",
	},
	"Hoyt Schermerhorn Streets": {
		"A": "A42",
		"C": "A42",
		"G": "A42",
	},
	"Lexington Av/63rd Street": {
		"Q": "B08",
		"F": "B08",
	},
	"Church Avenue": {
		"Q": "D28",
		"B": "D28",
		"2": "244",
		"G": "F27",
		"F": "F27",
	},
	"Spring Street": {
		"6X": "638",
		"C":  "A33",
		"E":  "A33",
		"6":  "638",
	},
	"Kingsbridge Road": {
		"4": "406",
		"D": "D04",
	},
	"City Hall": {
		"R": "R24",
		"W": "R24",
	},
	"New Utrecht Avenue": {
		"N": "N04",
	},
	"Shepherd Avenue": {
		"C": "A54",
	},
	"Mosholu Parkway": {
		"4": "402",
	},
	"7th Avenue": {
		"B": "D25",
		"E": "D14",
		"D": "D14",
		"G": "F24",
		"F": "F24",
		"Q": "D25",
	},
	"Astoria Blvd": {
		"W": "R03",
		"N": "R03",
	},
	"181st Street": {
		"1": "111",
		"A": "A06",
	},
	"Eastchester Dyre Avenue": {
		"5": "501",
	},
	"Gates Avenue": {
		"Z": "J30",
		"J": "J30",
	},
	"Bay Parkway": {
		"N": "N07",
		"D": "B21",
		"F": "F32",
	},
	"2nd Avenue": {
		"F": "F14",
	},
	"170th Street": {
		"4": "412",
		"D": "D09",
	},
	"18th Street": {
		"1": "131",
		"2": "131",
		"5": "131",
	},
	"Jefferson Street": {
		"L": "L15",
	},
	"Coney Island Stillwell Avenue": {
		"Q": "D43",
		"N": "D43",
		"D": "D43",
		"F": "D43",
	},
	"25th Street": {
		"R": "R35",
	},
	"Prospect Avenue": {
		"R": "R34",
		"2": "219",
		"5": "219",
	},
	"Grant Avenue": {
		"A": "A57",
	},
	"Avenue J": {
		"Q": "D33",
	},
	"Cortlandt Street": {
		"1": "138",
		"R": "R25",
		"5": "138",
		"W": "R25",
	},
	"Avenue I": {
		"F": "F31",
	},
	"163rd Street Amsterdam Avenue": {
		"C": "A10",
	},
	"125th Street": {
		"A":  "A15",
		"C":  "A15",
		"B":  "A15",
		"D":  "A15",
		"1":  "116",
		"3":  "225",
		"2":  "225",
		"5":  "225",
		"4":  "621",
		"6":  "621",
		"6X": "621",
	},
	"Christopher Street Sheridan Square": {
		"1": "133",
		"2": "133",
		"5": "133",
	},
	"14th Street Union Square": {
		"N":  "R20",
		"Q":  "R20",
		"R":  "R20",
		"4":  "635",
		"W":  "R20",
		"6":  "635",
		"6X": "635",
	},
	"Avenue P": {
		"F": "F34",
	},
	"9th Street": {
		"R": "R33",
	},
	"14th Street": {
		"A": "A31",
		"C": "A31",
		"E": "A31",
		"F": "D19",
		"1": "132",
		"3": "132",
		"2": "132",
		"5": "132",
	},
	"Pennsylvania Avenue": {
		"3": "255",
		"4": "255",
	},
	"Eastern Parkway Brooklyn Museum": {
		"3": "238",
		"2": "238",
		"4": "238",
	},
	"8th Avenue": {
		"L": "L01",
		"N": "N02",
	},
	"Avenue X": {
		"F": "F38",
	},
	"West Farms Square East Tremont Avenue": {
		"2": "214",
		"5": "214",
	},
	"Fulton Street": {
		"A": "A38",
		"C": "A38",
		"G": "G36",
		"J": "M22",
		"3": "229",
		"2": "229",
		"4": "418",
		"Z": "M22",
	},
	"Sutphin Blvd Archer Av JFK Airport": {
		"Z": "G06",
		"J": "G06",
		"E": "G06",
	},
	"Morris Park": {
		"5": "505",
	},
	"West 8th Street NY Aquarium": {
		"Q": "D42",
		"F": "D42",
	},
	"Morgan Avenue": {
		"L": "L14",
	},
	"Ozone Park Lefferts Blvd": {
		"A": "A65",
	},
	"Rockaway Avenue": {
		"C": "A50",
		"3": "253",
		"4": "253",
	},
	"65th Street": {
		"R": "G15",
		"E": "G15",
	},
	"53rd Street": {
		"R": "R40",
	},
	"Atlantic Avenue": {
		"L": "L24",
	},
	"Jackson Heights Roosevelt Avenue": {
		"R": "G14",
		"E": "G14",
		"F": "G14",
	},
	"66th Street Lincoln Center": {
		"1": "124",
		"2": "124",
		"5": "124",
	},
	"Utica Avenue": {
		"A": "A48",
		"C": "A48",
	},
	"Alabama Avenue": {
		"Z": "J24",
		"J": "J24",
	},
	"Longwood Avenue": {
		"6": "614",
	},
	"85th Street Forest Parkway": {
		"J": "J16",
	},
	"Zerega Avenue": {
		"6X": "606",
		"6":  "606",
	},
	"Middletown Road": {
		"6X": "603",
		"6":  "603",
	},
	"Lafayette Avenue": {
		"C": "A43",
	},
	"176th Street": {
		"4": "410",
	},
	"75th Street": {
		"Z": "J17",
		"J": "J17",
	},
	"79th Street": {
		"1": "122",
		"2": "122",
		"5": "122",
		"D": "B18",
	},
	"Cleveland Street": {
		"J": "J22",
	},
	"Kingston Throop Avenues": {
		"C": "A47",
	},
	"Chambers Street": {
		"A": "A36",
		"C": "A36",
		"J": "M21",
		"1": "137",
		"3": "137",
		"2": "137",
		"5": "137",
		"Z": "M21",
	},
	"Park Pl": {
		"3": "228",
		"2": "228",
	},
	"Norwood Avenue": {
		"Z": "J21",
		"J": "J21",
	},
	"Canal Street": {
		"A":  "A34",
		"Q":  "Q01",
		"C":  "A34",
		"E":  "A34",
		"J":  "M20",
		"N":  "Q01",
		"1":  "135",
		"2":  "135",
		"5":  "135",
		"W":  "R23",
		"6":  "639",
		"6X": "639",
		"R":  "R23",
		"Z":  "M20",
	},
	"Franklin Avenue": {
		"C": "A45",
		"3": "239",
		"2": "239",
		"4": "239",
	},
	"116th Street Columbia University": {
		"1": "117",
	},
	"46th Street": {
		"R": "G18",
		"E": "G18",
	},
	"Essex Street": {
		"Z": "M18",
		"J": "M18",
	},
	"Crescent Street": {
		"Z": "J20",
		"J": "J20",
	},
	"Forest Hills 71st Avenue": {
		"R": "G08",
		"E": "G08",
		"F": "G08",
	},
	"Kings Highway": {
		"Q": "D35",
		"N": "N08",
		"B": "D35",
		"F": "F35",
	},
	"51st Street": {
		"6X": "630",
		"6":  "630",
	},
	"Marble Hill 225th Street": {
		"1": "106",
	},
	"Neptune Avenue": {
		"F": "F39",
	},
	"Buhre Avenue": {
		"6X": "602",
		"6":  "602",
	},
	"33rd Street": {
		"6X": "632",
		"6":  "632",
	},
	"Kew Gardens Union Turnpike": {
		"E": "F06",
		"F": "F06",
	},
}
