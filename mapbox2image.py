import requests
from keys import parameters

URL_ENDPOINT = "https://a.tiles.mapbox.com/v4/{mapid}/12/{x}/{y}@2x.png?access_token={access_token}"
TILE_FILENAME_FORMAT = "tile_{x}_{y}.png"
TILE_FILENAME_DIR = "output"


def save_tile(x,y):
    parameters["x"] = x
    parameters["y"] = y
    r = requests.get(URL_ENDPOINT.format(**parameters))
    if r.status_code != 200:
        print 'Error accessing {} ...'.format(target_url)
        print '(status: {}): {}'.format(r.status_code, r.text)
        return
    with open(TILE_FILENAME_DIR + "/" + TILE_FILENAME_FORMAT.format(**parameters),"wb") as tilefile:
        tilefile.write(r.content)
        print '--success with {x}x{y}'.format(**parameters)


for x in xrange(655,675):
    for y in xrange(1581,1591):
        save_tile(x,y)
