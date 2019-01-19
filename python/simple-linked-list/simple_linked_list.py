class Node:
    def __init__(self, value, next_node=None):
        self._value = value
        self._next_node = next_node

    def value(self):
        return self._value

    def next(self):
        return self._next_node

    def __str__(self):
        return "({} -> {})".format(self._value, self._next_node)


class LinkedList:
    def __init__(self, values=[]):
        self._size = 0
        self._head = None
        self._current = None

        for value in values:
            self.push(value)

    def __len__(self):
        return self._size

    def head(self):
        if self._head is None:
            raise EmptyListException("Head is empty")
        return self._head

    def push(self, value):
        node = Node(value, self._head)
        self._head = node
        self._size += 1

    def pop(self):
        current_head = self.head()
        self._head = current_head.next()
        self._size -= 1
        return current_head.value()

    def reversed(self):
        pass

    def __iter__(self):
        return LinkedListIterator(self)


class LinkedListIterator:
    def __init__(self, linked_list):
        self.current = linked_list._head

    def __iter__(self):
        return self

    def __next__(self):
        if self.current is None:
            raise StopIteration
        value = self.current.value()
        self.current = self.current.next()
        return value


class EmptyListException(Exception):
    def __init__(self, message):
        self.message = message
