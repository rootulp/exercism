class Node(object):
    def __init__(self, value, next_node=None):
        self._value = value
        self._next_node = next_node

    def value(self):
        return self._value

    def next(self):
        return self._next_node


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
        return self._head

    def push(self, value):
        current = self._head
        while current is not None:
            current = current.next()

        node = Node(value, current)
        self._head = node
        self._size += 1

    def pop(self):
        pass

    def reversed(self):
        pass


class EmptyListException(Exception):
    def __init__(self, message):
        self.message = message
