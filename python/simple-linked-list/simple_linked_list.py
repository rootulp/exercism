class Node(object):
    def __init__(self, value):
        pass

    def value(self):
        pass

    def next(self):
        pass


class LinkedList(object):
    def __init__(self, values=[]):
        self._size = 0
        self._head = None

        for value in values:
            self.push(value)

    def __len__(self):
        return self._size

    def head(self):
        if self._head is None:
            raise EmptyListException("Head is empty")

    def push(self, value):
        node = Node(value)
        if self._head is None:
            self._head = node

        self._size += 1

    def pop(self):
        pass

    def reversed(self):
        pass


class EmptyListException(Exception):
    def __init__(self, message):
        self.message = message
