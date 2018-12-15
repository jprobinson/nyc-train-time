import csv
import re
import json
from collections import defaultdict


trains = {
        "1": {"northbound": "Bronx","southbound": "South&nbsp;Ferry"},
        "2": {"northbound": "Bronx","southbound": "Brooklyn"},
        "3": {"northbound": "Harlem","southbound": "Brooklyn"},
        "4": {"northbound": "Bronx","southbound": "Brooklyn"},
        "5": {"northbound": "Bronx","southbound": "Brooklyn"},
        "5X": {"northbound": "Bronx","southbound": "Brooklyn"},
        "6": {"northbound": "Bronx","southbound": "Brooklyn Brdg"},
        "7": {"northbound": "Queens","southbound": "Manhattan"},
        "6X": {"northbound": "Bronx","southbound": "Brooklyn Brdg"},
        "S": {"northbound": "","southbound": ""},
        "L": {"northbound": "Manhattan","southbound": "Brooklyn"},
        "B": {"northbound": "Bronx","southbound": "Brooklyn"},
        "D": {"northbound": "Bronx","southbound": "Brooklyn"},
        "A": {"northbound": "Manhattan","southbound": "Queens"},
        "G": {"northbound": "Queens","southbound": "Brooklyn"},
        "C": {"northbound": "Manhattan","southbound": "Brooklyn"},
        "E": {"northbound": "Queens","southbound": "Manhattan"},
        "N": {"northbound": "Manhattan","southbound": "Brooklyn"},
        "Q": {"northbound": "Manhattan","southbound": "Brooklyn"},
        "R": {"northbound": "Queens","southbound": "Brooklyn"},
        "W": {"northbound": "Queens","southbound": "Manhattan"},
        "J": {"northbound": "Queens","southbound": "Manhattan"},
        "F": {"northbound": "Queens","southbound": "Brooklyn"},
        "M": {"northbound": "Queens","southbound": "Brooklyn"},
        "B": {"northbound": "Bronx","southbound": "Brooklyn"},
        "D": {"northbound": "Bronx","southbound": "Brooklyn"},
        "Z": {"northbound": "Queens","southbound": "Manhattan"},
}

trips = dict()
for t in trains:
    trips[t] = dict()

with open('trips.txt','r') as csvin:
    reader=csv.DictReader(csvin)
    for line in reader:
        if line['route_id'] in trips:
            trips[line['trip_id']] = line['route_id']


sstop_seqs = dict()
nstop_seqs = dict()

with open('stop_times.txt','r') as csvin:
    reader=csv.DictReader(csvin)
    for line in reader:
        if line['trip_id'] in trips:
            train = trips[line['trip_id']]

            if train not in nstop_seqs:
                nstop_seqs[train] = defaultdict(list)

            if train not in sstop_seqs:
                sstop_seqs[train] = defaultdict(list)

            if line['stop_id'].endswith('S'):
                sstop_seqs[train][line['trip_id']].append([line['stop_id']])

train_stops = defaultdict(list)

for train in sstop_seqs:
    trips = sstop_seqs[train]
    for trip_id in trips:
        stops = trips[trip_id]

        if len(train_stops[train]) < len(stops):
            train_stops[train] = {'stops': stops}


all_stops = dict()
avenue = re.compile("Av($| -)")
avenues = re.compile("Avs($| -)")
street = re.compile("St($| -)")
streets = re.compile("Sts($| -)")

with open('stops.txt','r') as csvin:
    reader=csv.DictReader(csvin)
    out = dict()
    for line in reader:
        # all_stops[line['stop_id']] = {"name":line['stop_name'],"lat":line['stop_lat'],"long":line['stop_lon']}
        all_stops[line['stop_id']] = line['stop_name']
        '''
        name = line['stop_name'].replace("(",",").replace(")","")
        name = name.replace("Hts", "Heights")
        name = name.replace("Sq", "Square")
        name = name.replace("Pkwy", "Parkway")
        name = name.replace("Blvd", "Boulevard")
        name = name.replace("Hwy", "Highway")
        name = name.replace("Ctr", "Center")
        name = name.replace("1 ", "1st ")
        name = name.replace("2 ", "2nd ")
        name = name.replace("2-", "2nd ")
        name = name.replace("3 ", "3rd ")
        name = name.replace("4 ", "4th ")
        name = name.replace("4-", "4th ")
        name = name.replace("5 ", "5th ")
        name = name.replace("6 ", "6th ")
        name = name.replace("7 ", "7th ")
        name = name.replace("7-", "7th ")
        name = name.replace("8 ", "8th ")
        name = name.replace("9 ", "9th ")
        name = name.replace("0 ", "0th ")
        name = name.replace("E ", "East ")
        name = name.replace("W ", "West ")
        name = name.replace("N ", "North ")
        name = name.replace("S ", "South ")
        name = avenue.sub("Avenue", name)
        name = avenues.sub("Avenues", name)
        name = street.sub("Street", name)
        name = streets.sub("Streets", name)
        name = name.replace(" - ", " ")
        name = name.replace("/", " ")
        names = set(name.split(","))
        names.add(name)
        all_stops[line['stop_id']] = {"value":name,"synonyms":list(names)}
        out[name] = {"synonyms":list(names),"value":name}

    outs = []
    for k in out:
        outs.append(out[k])
    print json.dumps(outs)
    exit(0)

        '''


out = dict()
for train, stops in train_stops.items():
    train_stops[train]["northbound"] = trains[train]["northbound"]
    train_stops[train]["southbound"] = trains[train]["southbound"]
    for i, stop in enumerate(stops['stops']):

        train_stops[train]['stops'][i].append(all_stops[stop[0]])
        train_stops[train]['stops'][i][0] = train_stops[train]['stops'][i][0][:-1]
        '''
        if all_stops[stop[0]]["value"] in out:
            out[all_stops[stop[0]]["value"]][train] = stop[0]
            continue

        out[all_stops[stop[0]]["value"]] = {train:stop[0]}
      '''

print json.dumps(train_stops)
# print json.dumps(out)

