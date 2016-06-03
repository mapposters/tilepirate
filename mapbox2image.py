import requests
import json
from keys import parameters
import os.path

URL_ENDPOINT = "https://api.mapbox.com/styles/v1/{username}/{mapid}/tiles/{z}/{x}/{y}@2x?access_token={access_token}"
TILE_FILENAME_FORMAT = "tile{mapid}-{z}_{x}x{y}.png"
TILE_FILENAME_DIR = "tiles"


def save_tile(style, x,y, z):
    parameters["mapid"] = style["mapid"]
    parameters["x"] = x
    parameters["y"] = y
    parameters["z"] = z
    tile_filename = TILE_FILENAME_DIR + "/" + TILE_FILENAME_FORMAT.format(**parameters)
    if os.path.isfile(tile_filename):
      print "--- tile for %s seems to exist " % tile_filename
      return
    r = requests.get(URL_ENDPOINT.format(**parameters))
    if r.status_code != 200:
        print '--error accesing {z}:{x}x{y}'.format(parameters)
        return
    with open(tile_filename,"wb") as tilefile:
        tilefile.write(r.content)
        print '--success with {z}:{x}x{y}'.format(**parameters)


areas = json.loads(open("areas.json","rb").read())
styles = json.loads(open("styles.json","rb").read())

for name, area in areas.iteritems():
  for style_name, style in styles.iteritems():
    print 'Getting tiles for %s in %s ...' % (name, style_name)
    for x in xrange(area["tileRange"]["Min"]["X"],area["tileRange"]["Max"]["X"] + 1):
        for y in xrange(area["tileRange"]["Min"]["Y"],area["tileRange"]["Max"]["Y"] + 1):
            save_tile(style, x,y, area["z"])
