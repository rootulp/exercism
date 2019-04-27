import json


class RestAPI(object):

    def __init__(self, database=None):
        self.database = database

    def get(self, url, payload=None):
        return json.dumps(self.database)

    def post(self, url, payload=None):
        pass
