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
        for value in values:
            self.push(value)

    def __len__(self):
        return self.size

    def head(self):
        pass

    def push(self, value):
        node = Node(value)
        if self.head is None:
            self.head = node

        self.size += 1

    def pop(self):
        pass

    def reversed(self):
        pass


class EmptyListException(Exception):
    pass
