import requests
import json
from keys import parameters
import os.path



URL_ENDPOINT = "https://api.mapbox.com/styles/v1/{username}/{mapid}/static/-79.295664,43.716727,10.16,0.00,0.00/1280x1280?access_token={access_token}"

def get_tile(tile_filename):
    r = requests.get(URL_ENDPOINT.format(**parameters))
    if r.status_code != 200:
        print '  error(status: {}): {}'.format(r.status_code, r.text)
        return
    with open(tile_filename,"wb") as tilefile:
        tilefile.write(r.content)
        print '--success'.format(**parameters)

get_tile("static.png")
