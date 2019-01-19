class Node(object):
    def __init__(self, value):
        pass

    def value(self):
        pass

    def next(self):
        pass


class LinkedList(object):
    def __init__(self, values=[]):
        self.size = 0

    def __len__(self):
        return self.size

    def head(self):
        pass

    def push(self, value):
        pass

    def pop(self):
        pass

    def reversed(self):
        pass


class EmptyListException(Exception):
    pass
