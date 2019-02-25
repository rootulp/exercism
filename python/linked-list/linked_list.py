class Node(object):
    def __init__(self, value, succeeding=None, previous=None):
        self.value = value
        self.succeeding = succeeding
        self.previous = previous


class LinkedList(object):

    def __init__(self):
        self.head = None
        self.tail = None

    def push(self, value):
        new_node = Node(value)
        if self.tail is not None:
            self.tail.succeeding = new_node
            new_node.previous = self.tail
        self.tail = new_node
        if self.head is None:
            self.head = new_node

    def pop(self):
        if self.tail is not None:
            popped_node = self.tail
            self.tail = popped_node.previous
            return popped_node.value

    def unshift(self, value):
        new_node = Node(value)
        if self.head is not None:
            self.head.previous = new_node
            new_node.succeeding = self.head
        self.head = new_node
        if self.tail is None:
            self.tail = new_node

    def shift(self):
        if self.head is not None:
            shifted_node = self.head
            self.head = shifted_node.succeeding
            return shifted_node.value
