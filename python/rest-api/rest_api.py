import json


class RestAPI(object):

    def __init__(self, database=None):
        self.database = database

    def get(self, url, payload=None):
        if payload is not None:
            payload = json.loads(payload)
            # print(payload['users'])
            usernames = payload['users']
            print(usernames)
            return self.get_usernames(usernames)
        return json.dumps(self.database)

    def post(self, url, payload=None):
        payload = json.loads(payload)
        username = payload['user']
        self.create_user(username)
        return self.get_user(username)

    def create_user(self, username):
        new_user = {
            'name': username,
            'owes': {},
            'owed_by': {},
            'balance': 0,
        }

        self.database['users'].append(new_user)

    def get_user(self, username):
        users = self.database['users']
        for user in users:
            if user['name'] == username:
                return json.dumps(user)

    def get_usernames(self, usernames):
        return json.dumps({'users': [user for user in self.database['users']
                           if user['name'] in usernames]})

