import json


class RestAPI(object):

    def __init__(self, database=None):
        self.database = database

    # Public API
    def get(self, url, payload=None):
        if payload is not None:
            payload = json.loads(payload)
            usernames = payload['users']
            return json.dumps({'users': self.get_users(usernames)})
        return json.dumps(self.database)

    def post(self, url, payload=None):
        if payload is None:
            raise ValueError("Payload must not be None.")

        payload = json.loads(payload)

        if url == '/add':
            return self.add(payload)
        elif url == '/iou':
            return self.iou(payload)

    # Private methods

    def add(self, payload):
        username = payload['user']
        self.create_user(username)
        return json.dumps(self.get_user(username))

    def iou(self, payload):
        lender = self.get_user(payload['lender'])
        borrower = self.get_user(payload['borrower'])
        amount = payload['amount']

        lender['owed_by'].setdefault(borrower['name'], 0)
        lender['owed_by'][borrower['name']] += amount
        lender['balance'] += amount

        borrower['owes'].setdefault(lender['name'], 0)
        borrower['owes'][lender['name']] += amount
        borrower['balance'] -= amount

        usernames = [payload['lender'], payload['borrower']]
        return json.dumps({'users': self.get_users(usernames)})

    def create_user(self, username):
        new_user = {
            'name': username,
            'owes': {},
            'owed_by': {},
            'balance': 0,
        }

        self.database['users'].append(new_user)

    def get_users(self, usernames):
        users = [self.get_user(username) for username in usernames]
        return sorted(users, key=lambda user: user['name'])

    def get_user(self, username):
        users = self.database['users']
        for user in users:
            if user['name'] == username:
                return user
