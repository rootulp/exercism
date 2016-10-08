from collections import deque


class CircularBuffer:

    def __init__(self, capacity):
        self.capacity = capacity
        self.buffer = deque([], capacity)

    def read(self):
        if self.empty():
            raise BufferEmptyException
        return self.buffer.popleft()

    def write(self, data):
        if self.full():
            raise BufferFullException
        self.buffer.append(data)

    def overwrite(self, data):
        if self.full():
            self.buffer.popleft()
        self.buffer.append(data)

    def clear(self):
        self.buffer.clear()

    def empty(self):
        return len(self.buffer) == 0

    def full(self):
        return len(self.buffer) == self.capacity


class BufferFullException(Exception):
    pass


class BufferEmptyException(Exception):
    pass
