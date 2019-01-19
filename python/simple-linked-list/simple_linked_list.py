class Node:
    def __init__(self, value, next_node=None):
        self._value = value
        self._next_node = next_node

    def value(self):
        return self._value

    def next(self):
        return self._next_node


class LinkedList:
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
        new_node = Node(value, self._head)
        self._head = new_node
        self._size += 1

    def pop(self):
        current_head = self.head()
        self._head = current_head.next()
        self._size -= 1
        return current_head.value()

    def reversed(self):
        return LinkedList(self)

    def __iter__(self):
        return LinkedListIterator(self)


class LinkedListIterator:
    def __init__(self, linked_list):
        self._current = linked_list._head

    def __iter__(self):
        return self

    def __next__(self):
        if self._current is None:
            raise StopIteration
        current_value = self._current.value()
        self._current = self._current.next()
        return current_value


class EmptyListException(Exception):
    def __init__(self, message):
        self.message = message
