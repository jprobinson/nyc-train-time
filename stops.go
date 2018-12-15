package nyctraintime

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
	"Winthrop Street": {
		"2": "243",
	},
	"New Lots Avenue": {
		"3": "257",
		"L": "L27",
	},
	"Buhre Avenue": {
		"6X": "602",
	},
	"Burnside Avenue": {
		"4": "409",
	},
	"168th Street": {
		"A": "A09",
		"C": "A09",
	},
	"Spring Street": {
		"A":  "A33",
		"6X": "638",
		"C":  "A33",
		"E":  "A33",
		"6":  "638",
	},
	"High Street": {
		"A": "A40",
		"C": "A40",
	},
	"Ralph Avenue": {
		"A": "A49",
		"C": "A49",
	},
	"Grand Army Plaza": {
		"3": "237",
		"2": "237",
	},
	"21st Street": {
		"G": "G24",
	},
	"Middletown Rd": {
		"6X": "603",
	},
	"Grand Central 42nd Street": {
		"6X": "631",
		"4":  "631",
		"7":  "723",
		"6":  "631",
	},
	"15th Street Prospect Park": {
		"G": "F25",
		"F": "F25",
	},
	"Bushwick Avenue Aberdeen Street": {
		"L": "L21",
	},
	"Hunters Point Avenue": {
		"7": "720",
	},
	"Woodhaven Boulevard": {
		"R": "G11",
		"J": "J15",
		"Z": "J15",
	},
	"Lorimer Street": {
		"J": "M13",
		"L": "L10",
	},
	"116th Street": {
		"A":  "A16",
		"C":  "A16",
		"B":  "A16",
		"3":  "226",
		"2":  "226",
		"6":  "622",
		"6X": "622",
	},
	"Franklin Avenue": {
		"A": "A45",
		"C": "A45",
		"3": "239",
		"2": "239",
		"4": "239",
	},
	"Rector Street": {
		"1": "139",
		"R": "R26",
		"W": "R26",
	},
	"5th Avenue": {
		"7": "724",
	},
	"42nd Street Port Authority Bus Terminal": {
		"A": "A27",
		"C": "A27",
		"E": "A27",
	},
	"53rd Street": {
		"R": "R40",
		"N": "R40",
	},
	"46th Street": {
		"R": "G18",
		"7": "714",
	},
	"Bronx Park East": {
		"2": "212",
	},
	"1st Avenue": {
		"L": "L06",
	},
	"Ocean Parkway": {
		"Q": "D41",
	},
	"Delancey Street": {
		"F": "F15",
	},
	"231st Street": {
		"1": "104",
	},
	"Beach 25th Street": {
		"A": "H10",
	},
	"138th Street Grand Concourse": {
		"4": "416",
	},
	"Myrtle Willoughby Avenues": {
		"G": "G32",
	},
	"South Ferry": {
		"1": "142",
	},
	"Astoria Ditmars Boulevard": {
		"W": "R01",
		"N": "R01",
	},
	"Rockaway Boulevard": {
		"A": "A61",
	},
	"14th Street": {
		"A": "A31",
		"C": "A31",
		"E": "A31",
		"F": "D19",
		"1": "132",
		"3": "132",
		"2": "132",
	},
	"Ditmas Avenue": {
		"F": "F29",
	},
	"Burke Avenue": {
		"2": "209",
	},
	"Cypress Hills": {
		"J": "J19",
	},
	"5th Av 53rd Street": {
		"E": "F12",
	},
	"176th Street": {
		"4": "410",
	},
	"Eastchester Dyre Avenue": {
		"5": "501",
	},
	"169th Street": {
		"F": "F02",
	},
	"190th Street": {
		"A": "A05",
	},
	"Canarsie Rockaway Parkway": {
		"L": "L29",
	},
	"2nd Avenue": {
		"F": "F14",
	},
	"Utica Avenue": {
		"A": "A48",
		"C": "A48",
	},
	"Newkirk Plaza": {
		"Q": "D31",
		"B": "D31",
	},
	"63rd Dr Rego Park": {
		"R": "G10",
	},
	"207th Street": {
		"1": "108",
	},
	"Mosholu Parkway": {
		"4": "402",
	},
	"36th Street": {
		"R": "R36",
		"D": "R36",
		"N": "R36",
	},
	"Court Street": {
		"R": "R28",
	},
	"Gates Avenue": {
		"Z": "J30",
		"J": "J30",
	},
	"Bowling Green": {
		"4": "420",
	},
	"34th Street 11st Avenue": {
		"7": "726",
	},
	"71st Street": {
		"D": "B17",
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
	"Lexington Av 59th Street": {
		"R": "R11",
		"W": "R11",
		"N": "R11",
	},
	"183rd Street": {
		"4": "408",
	},
	"111st Street": {
		"J": "J13",
		"7": "705",
	},
	"Sutphin Boulevard Archer Avenue JFK Airport": {
		"Z": "G06",
		"J": "G06",
	},
	"Bay Parkway": {
		"N": "N07",
		"D": "B21",
		"F": "F32",
	},
	"Pennsylvania Avenue": {
		"3": "255",
	},
	"St Lawrence Avenue": {
		"6": "609",
	},
	"WTC Cortlandt": {
		"1": "138",
	},
	"Castle Hill Avenue": {
		"6X": "607",
	},
	"Kings Highway": {
		"Q": "D35",
		"N": "N08",
		"B": "D35",
		"F": "F35",
	},
	"Montrose Avenue": {
		"L": "L13",
	},
	"Seneca Avenue": {
		"M": "M06",
	},
	"Beach 60th Street": {
		"A": "H07",
	},
	"Steinway Street": {
		"R": "G19",
	},
	"137th Street City College": {
		"1": "115",
	},
	"79th Street": {
		"1": "122",
		"D": "B18",
	},
	"Euclid Avenue": {
		"A": "A55",
		"C": "A55",
	},
	"9th Avenue": {
		"D": "B12",
	},
	"Elder Avenue": {
		"6": "611",
	},
	"Brooklyn Bridge City Hall": {
		"6X": "640",
		"4":  "640",
		"6":  "640",
	},
	"23rd Street": {
		"A":  "A30",
		"C":  "A30",
		"E":  "A30",
		"F":  "D18",
		"N":  "R19",
		"1":  "130",
		"R":  "R19",
		"W":  "R19",
		"6":  "634",
		"6X": "634",
	},
	"Myrtle Avenue": {
		"Z": "M11",
		"J": "M11",
		"M": "M11",
	},
	"Far Rockaway Mott Avenue": {
		"A": "H11",
	},
	"Bedford Park Boulevard Lehman College": {
		"4": "405",
	},
	"42nd Street Bryant Pk": {
		"B": "D16",
		"D": "D16",
		"F": "D16",
	},
	"Bay Ridge 95th Street": {
		"R": "R45",
	},
	"25th Avenue": {
		"D": "B22",
	},
	"Alabama Avenue": {
		"Z": "J24",
		"J": "J24",
	},
	"Junction Boulevard": {
		"7": "707",
	},
	"Beverly Rd": {
		"2": "245",
	},
	"50th Street": {
		"1": "126",
		"A": "A25",
		"C": "A25",
		"E": "A25",
		"D": "B14",
	},
	"Beach 36th Street": {
		"A": "H09",
	},
	"Hewes Street": {
		"J": "M14",
	},
	"69th Street": {
		"7": "711",
	},
	"Lafayette Avenue": {
		"A": "A43",
		"C": "A43",
	},
	"47th 50th Streets Rockefeller Center": {
		"B": "D15",
		"D": "D15",
		"F": "D15",
	},
	"Union Square 14th Street": {
		"L": "L03",
	},
	"Carroll Street": {
		"G": "F21",
		"F": "F21",
	},
	"Coney Island Stillwell Avenue": {
		"Q": "D43",
		"N": "D43",
		"D": "D43",
		"F": "D43",
	},
	"75th Street": {
		"Z": "J17",
		"J": "J17",
	},
	"President Street": {
		"2": "241",
	},
	"Norwood Avenue": {
		"Z": "J21",
		"J": "J21",
	},
	"Livonia Avenue": {
		"L": "L26",
	},
	"62nd Street": {
		"D": "B16",
	},
	"Cortelyou Rd": {
		"Q": "D30",
	},
	"3rd Avenue 138th Street": {
		"6X": "619",
		"6":  "619",
	},
	"57th Street": {
		"F": "B10",
	},
	"Briarwood Van Wyck Boulevard": {
		"F": "F05",
	},
	"West 8th Street NY Aquarium": {
		"Q": "D42",
		"F": "D42",
	},
	"Fordham Rd": {
		"4": "407",
		"D": "D05",
	},
	"Grand Avenue Newtown": {
		"R": "G12",
	},
	"Wall Street": {
		"3": "230",
		"2": "230",
		"4": "419",
	},
	"36th Avenue": {
		"W": "R06",
		"N": "R06",
	},
	"Sutter Avenue Rutland Rd": {
		"3": "251",
	},
	"34th Street Penn Station": {
		"A": "A28",
		"C": "A28",
		"E": "A28",
		"1": "128",
		"3": "128",
		"2": "128",
	},
	"Lexington Av 63rd Street": {
		"Q": "B08",
		"F": "B08",
	},
	"81st Street Museum of Natural History": {
		"A": "A21",
		"C": "A21",
		"B": "A21",
	},
	"East Broadway": {
		"F": "F16",
	},
	"Essex Street": {
		"Z": "M18",
		"J": "M18",
	},
	"Flushing Avenue": {
		"J": "M12",
		"G": "G31",
	},
	"DeKalb Avenue": {
		"Q": "R30",
		"R": "R30",
		"B": "R30",
		"L": "L16",
	},
	"Parsons Boulevard": {
		"E": "F03",
		"F": "F03",
	},
	"Flushing Main Street": {
		"7": "701",
	},
	"86th Street": {
		"A":  "A20",
		"Q":  "Q04",
		"C":  "A20",
		"B":  "A20",
		"N":  "N10",
		"1":  "121",
		"R":  "R44",
		"4":  "626",
		"6":  "626",
		"6X": "626",
	},
	"Gun Hill Rd": {
		"2": "208",
		"5": "503",
	},
	"Canal Street": {
		"A":  "A34",
		"Q":  "Q01",
		"C":  "A34",
		"E":  "A34",
		"J":  "M20",
		"N":  "Q01",
		"1":  "135",
		"R":  "R23",
		"W":  "R23",
		"6":  "639",
		"6X": "639",
		"Z":  "M20",
	},
	"Central Park North ,110th Street": {
		"3": "227",
		"2": "227",
	},
	"103rd Street": {
		"A":  "A18",
		"C":  "A18",
		"B":  "A18",
		"1":  "119",
		"6":  "624",
		"6X": "624",
	},
	"Broadway-Lafayette Street": {
		"B": "D21",
		"D": "D21",
		"F": "D21",
	},
	"Harlem 148th Street": {
		"3": "301",
	},
	"Forest Hills 71st Avenue": {
		"R": "G08",
		"E": "G08",
		"F": "G08",
	},
	"Pelham Bay Park": {
		"6X": "601",
	},
	"Bleecker Street": {
		"6X": "637",
		"6":  "637",
	},
	"Kingston Avenue": {
		"3": "249",
	},
	"Brighton Beach": {
		"Q": "D40",
		"B": "D40",
	},
	"3rd Avenue 149th Street": {
		"2": "221",
	},
	"Clark Street": {
		"3": "231",
		"2": "231",
	},
	"Baychester Avenue": {
		"5": "502",
	},
	"Prospect Park": {
		"Q": "D26",
		"B": "D26",
	},
	"Sutter Avenue": {
		"L": "L25",
	},
	"Crescent Street": {
		"Z": "J20",
		"J": "J20",
	},
	"104th Street": {
		"Z": "J14",
		"J": "J14",
	},
	"Chambers Street": {
		"A": "A36",
		"C": "A36",
		"J": "M21",
		"1": "137",
		"3": "137",
		"2": "137",
		"Z": "M21",
	},
	"Allerton Avenue": {
		"2": "210",
	},
	"Shepherd Avenue": {
		"A": "A54",
		"C": "A54",
	},
	"Flatbush Avenue Brooklyn College": {
		"2": "247",
	},
	"Cathedral Parkway ,110th Street": {
		"A": "A17",
		"C": "A17",
		"B": "A17",
	},
	"Bedford Avenue": {
		"L": "L08",
	},
	"Atlantic Avenue Barclays Center": {
		"B": "D24",
		"D": "R31",
		"N": "R31",
		"Q": "D24",
		"3": "235",
		"2": "235",
		"4": "235",
		"R": "R31",
	},
	"121st Street": {
		"Z": "J12",
		"J": "J12",
	},
	"Marcy Avenue": {
		"Z": "M16",
		"J": "M16",
	},
	"145th Street": {
		"A": "A12",
		"C": "A12",
		"B": "D13",
		"D": "D13",
		"1": "114",
		"3": "302",
	},
	"67th Avenue": {
		"R": "G09",
	},
	"Liberty Avenue": {
		"A": "A52",
		"C": "A52",
	},
	"Bedford Nostrand Avenues": {
		"G": "G33",
	},
	"Simpson Street": {
		"2": "217",
	},
	"Roosevelt Island": {
		"F": "B06",
	},
	"168th Street Washington Heights": {
		"1": "112",
	},
	"Jamaica 179th Street": {
		"E": "F01",
		"F": "F01",
	},
	"65th Street": {
		"R": "G15",
	},
	"Norwood 205th Street": {
		"D": "D01",
	},
	"Jefferson Street": {
		"L": "L15",
	},
	"Houston Street": {
		"1": "134",
	},
	"Pelham Parkway": {
		"2": "211",
		"5": "504",
	},
	"170th Street": {
		"4": "412",
	},
	"Nassau Avenue": {
		"G": "G28",
	},
	"Times Square 42nd Street": {
		"Q": "R16",
		"W": "R16",
		"N": "R16",
		"1": "127",
		"3": "127",
		"2": "127",
		"7": "725",
		"R": "R16",
	},
	"191st Street": {
		"1": "110",
	},
	"Greenpoint Avenue": {
		"G": "G26",
	},
	"Bowery": {
		"Z": "M19",
		"J": "M19",
	},
	"Beach 44th Street": {
		"A": "H08",
	},
	"Broadway Jct": {
		"A": "A51",
		"C": "A51",
		"J": "J27",
		"Z": "J27",
		"L": "L22",
	},
	"Morris Park": {
		"5": "505",
	},
	"West 4th Street": {
		"A": "A32",
		"C": "A32",
		"B": "D20",
		"E": "A32",
		"D": "D20",
		"F": "D20",
	},
	"8th Avenue": {
		"L": "L01",
		"N": "N02",
	},
	"York Street": {
		"F": "F18",
	},
	"Jamaica Center Parsons Archer": {
		"Z": "G05",
		"J": "G05",
	},
	"80th Street": {
		"A": "A59",
	},
	"110th Street": {
		"6X": "623",
		"6":  "623",
	},
	"103rd Street Corona Plaza": {
		"7": "706",
	},
	"Sutphin Boulevard": {
		"F": "F04",
	},
	"215th Street": {
		"1": "107",
	},
	"Woodside 61st Street": {
		"7": "712",
	},
	"Bay Ridge Avenue": {
		"R": "R42",
	},
	"Wilson Avenue": {
		"L": "L20",
	},
	"Hunts Point Avenue": {
		"6X": "613",
		"6":  "613",
	},
	"4th Avenue": {
		"G": "F23",
		"F": "F23",
	},
	"Mt Eden Avenue": {
		"4": "411",
	},
	"Aqueduct North Conduit Avenue": {
		"A": "H02",
	},
	"20th Avenue": {
		"D": "B20",
		"N": "N06",
	},
	"Van Cortlandt Park 242nd Street": {
		"1": "101",
	},
	"8th Street NYU": {
		"R": "R21",
		"W": "R21",
		"N": "R21",
	},
	"Halsey Street": {
		"J": "J29",
		"L": "L19",
	},
	"Freeman Street": {
		"2": "216",
	},
	"East 180th Street": {
		"2": "213",
		"5": "213",
	},
	"Jackson Avenue": {
		"2": "220",
	},
	"World Trade Center": {
		"E": "E01",
	},
	"Newkirk Avenue": {
		"2": "246",
	},
	"5th Av 59th Street": {
		"R": "R13",
		"W": "R13",
		"N": "R13",
	},
	"Dyckman Street": {
		"1": "109",
		"A": "A03",
	},
	"Broadway": {
		"W": "R05",
		"G": "G30",
		"N": "R05",
	},
	"88th Street": {
		"A": "A60",
	},
	"96th Street": {
		"A":  "A19",
		"Q":  "Q05",
		"C":  "A19",
		"B":  "A19",
		"1":  "120",
		"3":  "120",
		"2":  "120",
		"6":  "625",
		"6X": "625",
	},
	"Middle Village Metropolitan Avenue": {
		"M": "M01",
	},
	"Tremont Avenue": {
		"D": "D07",
	},
	"Astor Pl": {
		"6X": "636",
		"6":  "636",
	},
	"49th Street": {
		"R": "R15",
		"W": "R15",
		"N": "R15",
	},
	"135th Street": {
		"A": "A14",
		"3": "224",
		"2": "224",
		"C": "A14",
		"B": "A14",
	},
	"Hoyt Schermerhorn Streets": {
		"A": "A42",
		"C": "A42",
		"G": "A42",
	},
	"116th Street Columbia University": {
		"1": "117",
	},
	"39th Avenue": {
		"W": "R08",
		"N": "R08",
	},
	"New Utrecht Avenue": {
		"N": "N04",
	},
	"Beach 67th Street": {
		"A": "H06",
	},
	"Whitehall Street": {
		"R": "R27",
		"W": "R27",
	},
	"Astoria Boulevard": {
		"W": "R03",
		"N": "R03",
	},
	"Cleveland Street": {
		"J": "J22",
	},
	"59th Street Columbus Circle": {
		"1": "125",
		"A": "A24",
		"C": "A24",
		"B": "A24",
		"D": "A24",
	},
	"Fort Hamilton Parkway": {
		"N": "N03",
		"D": "B13",
		"G": "F26",
		"F": "F26",
	},
	"18th Street": {
		"1": "131",
	},
	"Church Avenue": {
		"Q": "D28",
		"B": "D28",
		"2": "244",
		"G": "F27",
		"F": "F27",
	},
	"90th Street Elmhurst Avenue": {
		"7": "708",
	},
	"Parkchester": {
		"6X": "608",
		"6":  "608",
	},
	"Nereid Avenue": {
		"2": "204",
	},
	"Elmhurst Avenue": {
		"R": "G13",
	},
	"167th Street": {
		"4": "413",
	},
	"72nd Street": {
		"A": "A22",
		"Q": "Q03",
		"C": "A22",
		"B": "A22",
		"1": "123",
		"3": "123",
		"2": "123",
	},
	"Northern Boulevard": {
		"R": "G16",
	},
	"Mets Willets Point": {
		"7": "702",
	},
	"238th Street": {
		"1": "103",
	},
	"174th Street": {
		"2": "215",
	},
	"Howard Beach JFK Airport": {
		"A": "H03",
	},
	"163rd Street Amsterdam Avenue": {
		"A": "A10",
		"C": "A10",
	},
	"Nevins Street": {
		"3": "234",
		"2": "234",
		"4": "234",
	},
	"33rd Street": {
		"6X": "632",
		"7":  "716",
		"6":  "632",
	},
	"Prospect Avenue": {
		"R": "R34",
		"2": "219",
	},
	"Zerega Avenue": {
		"6X": "606",
	},
	"28th Street": {
		"N":  "R18",
		"1":  "129",
		"R":  "R18",
		"W":  "R18",
		"6":  "633",
		"6X": "633",
	},
	"74th Street Broadway": {
		"7": "710",
	},
	"9th Street": {
		"R": "R33",
	},
	"157th Street": {
		"1": "113",
	},
	"Graham Avenue": {
		"L": "L11",
	},
	"East 105th Street": {
		"L": "L28",
	},
	"Bay 50th Street": {
		"D": "B23",
	},
	"Inwood 207th Street": {
		"A": "A02",
	},
	"Brook Avenue": {
		"6": "618",
	},
	"21st Street Queensbridge": {
		"F": "B04",
	},
	"Parkside Avenue": {
		"Q": "D27",
	},
	"Westchester Square East Tremont Avenue": {
		"6X": "604",
	},
	"57th Street 7th Avenue": {
		"Q": "R14",
		"R": "R14",
		"W": "R14",
		"N": "R14",
	},
	"181st Street": {
		"1": "111",
		"A": "A06",
	},
	"6th Avenue": {
		"L": "L02",
	},
	"82nd Street Jackson Heights": {
		"7": "709",
	},
	"55th Street": {
		"D": "B15",
	},
	"225th Street": {
		"2": "206",
	},
	"Metropolitan Avenue": {
		"G": "G29",
	},
	"30th Avenue": {
		"W": "R04",
		"N": "R04",
	},
	"Fresh Pond Rd": {
		"M": "M04",
	},
	"Longwood Avenue": {
		"6": "614",
	},
	"Wakefield 241st Street": {
		"2": "201",
	},
	"Morgan Avenue": {
		"L": "L14",
	},
	"Atlantic Avenue": {
		"L": "L24",
	},
	"Grant Avenue": {
		"A": "A57",
	},
	"Clinton Washington Avenues": {
		"A": "A44",
		"C": "A44",
		"G": "G35",
	},
	"Whitlock Avenue": {
		"6": "612",
	},
	"125th Street": {
		"A":  "A15",
		"C":  "A15",
		"B":  "A15",
		"D":  "A15",
		"1":  "116",
		"3":  "225",
		"2":  "225",
		"4":  "621",
		"6":  "621",
		"6X": "621",
	},
	"Lexington Av 53rd Street": {
		"E": "F11",
	},
	"18th Avenue": {
		"N": "N05",
		"D": "B19",
		"F": "F30",
	},
	"Classon Avenue": {
		"G": "G34",
	},
	"Rockaway Avenue": {
		"A": "A50",
		"3": "253",
		"C": "A50",
	},
	"Borough Hall": {
		"3": "232",
		"2": "232",
		"4": "423",
	},
	"52nd Street": {
		"7": "713",
	},
	"Marble Hill 225th Street": {
		"1": "106",
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
	"149th Street Grand Concourse": {
		"2": "222",
		"4": "415",
	},
	"East 143rd Street St Mary's Street": {
		"6": "616",
	},
	"Sterling Street": {
		"2": "242",
	},
	"Eastern Parkway Brooklyn Museum": {
		"3": "238",
		"2": "238",
	},
	"Broad Street": {
		"Z": "M23",
		"J": "M23",
	},
	"Beverley Rd": {
		"Q": "D29",
	},
	"Myrtle Wyckoff Avenues": {
		"M": "M08",
		"L": "L17",
	},
	"Vernon Boulevard Jackson Avenue": {
		"7": "721",
	},
	"Morrison Av- Sound View": {
		"6": "610",
	},
	"Kingsbridge Rd": {
		"4": "406",
		"D": "D04",
	},
	"233rd Street": {
		"2": "205",
	},
	"Woodlawn": {
		"4": "401",
	},
	"Franklin Street": {
		"1": "136",
	},
	"Kosciuszko Street": {
		"J": "J31",
	},
	"51st Street": {
		"6X": "630",
		"6":  "630",
	},
	"Knickerbocker Avenue": {
		"M": "M09",
	},
	"Court Square": {
		"E": "F09",
		"G": "G22",
		"7": "719",
	},
	"66th Street Lincoln Center": {
		"1": "124",
	},
	"3rd Avenue": {
		"L": "L05",
	},
	"Bergen Street": {
		"3": "236",
		"2": "236",
		"G": "F20",
		"F": "F20",
	},
	"75th Avenue": {
		"F": "F07",
	},
	"East 149th Street": {
		"6": "615",
	},
	"59th Street": {
		"6X": "629",
		"R":  "R41",
		"N":  "R41",
		"4":  "629",
		"6":  "629",
	},
	"Cypress Avenue": {
		"6": "617",
	},
	"Sheepshead Bay": {
		"Q": "D39",
		"B": "D39",
	},
	"45th Street": {
		"R": "R39",
		"N": "R39",
	},
	"Cortlandt Street": {
		"R": "R25",
		"W": "R25",
	},
	"Central Avenue": {
		"M": "M10",
	},
	"7th Avenue": {
		"B": "D25",
		"E": "D14",
		"D": "D14",
		"G": "F24",
		"F": "F24",
		"Q": "D25",
	},
	"Kew Gardens Union Tpke": {
		"E": "F06",
		"F": "F06",
	},
	"Saratoga Avenue": {
		"3": "252",
	},
	"Neptune Avenue": {
		"F": "F39",
	},
	"Van Siclen Avenue": {
		"A": "A53",
		"3": "256",
		"J": "J23",
		"Z": "J23",
		"C": "A53",
	},
	"Queens Plaza": {
		"R": "G21",
		"E": "G21",
	},
	"175th Street": {
		"A": "A07",
	},
	"Forest Avenue": {
		"M": "M05",
	},
	"Nostrand Avenue": {
		"A": "A46",
		"3": "248",
		"C": "A46",
	},
	"Hoyt Street": {
		"3": "233",
		"2": "233",
	},
	"Intervale Avenue": {
		"2": "218",
	},
	"Neck Rd": {
		"Q": "D38",
	},
	"40th Street": {
		"7": "715",
	},
	"161st Street Yankee Stadium": {
		"4": "414",
	},
	"Cathedral Parkway": {
		"1": "118",
	},
	"68th Street Hunter College": {
		"6X": "628",
		"6":  "628",
	},
	"Grand Street": {
		"B": "D22",
		"D": "D22",
		"L": "L12",
	},
	"Broad Channel": {
		"A": "H04",
	},
	"Bedford Park Boulevard": {
		"D": "D03",
	},
	"Queensboro Plaza": {
		"W": "R09",
		"7": "718",
		"N": "R09",
	},
	"Junius Street": {
		"3": "254",
	},
	"25th Street": {
		"R": "R35",
	},
	"Chauncey Street": {
		"Z": "J28",
		"J": "J28",
	},
	"City Hall": {
		"R": "R24",
		"W": "R24",
	},
	"Christopher Street Sheridan Square": {
		"1": "133",
	},
	"West Farms Square East Tremont Avenue": {
		"2": "214",
	},
	"Avenue J": {
		"Q": "D33",
	},
	"Avenue H": {
		"Q": "D32",
	},
	"Avenue I": {
		"F": "F31",
	},
	"Avenue N": {
		"F": "F33",
	},
	"Avenue M": {
		"Q": "D34",
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
	"Avenue P": {
		"F": "F34",
	},
	"Avenue U": {
		"Q": "D37",
		"N": "N09",
		"F": "F36",
	},
	"Avenue X": {
		"F": "F38",
	},
	"Jackson Heights Roosevelt Avenue": {
		"R": "G14",
		"E": "G14",
		"F": "G14",
	},
	"155th Street": {
		"A": "A11",
		"C": "A11",
	},
	"Crown Heights Utica Avenue": {
		"3": "250",
		"4": "250",
	},
	"85th Street Forest Parkway": {
		"J": "J16",
	},
	"Prince Street": {
		"R": "R22",
		"W": "R22",
		"N": "R22",
	},
	"Park Pl": {
		"3": "228",
		"2": "228",
	},
	"77th Street": {
		"6X": "627",
		"R":  "R43",
		"6":  "627",
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
	"219th Street": {
		"2": "207",
	},
	"Smith 9th Streets": {
		"G": "F22",
		"F": "F22",
	},
	"Kingston Throop Avenues": {
		"A": "A47",
		"C": "A47",
	},
}
