class Rails:

    def __init__(self, num_rails):
        self.num_rails = num_rails
        self.rails = [[] for _ in range(self.num_rails)]

    def populate_rails_linear(self, message, rail_lengths):
        message_list = list(message)
        for index in xrange(len(self.rails)):
            for rail_length in xrange(rail_lengths[index]):
                self.rails[index].append(message_list.pop(0))

    def populate_rails_zig_zag(self, message):
        self.reset()
        for data in message:
            self.rails[self.rails_index].append(data)
            self.increment_rails_index()

    def to_string_linear(self):
        return ''.join([data for rail in self.rails for data in rail])

    def to_string_zig_zag(self, message):
        self.reset()
        output = []
        for _ in range(len(message)):
            output.append(self.rails[self.rails_index].pop(0))
            self.increment_rails_index()
        return ''.join(output)

    def increment_rails_index(self):
        self.rails_index_reverse_direction()
        if self.rails_index_increasing:
            self.rails_index += 1
        else:
            self.rails_index -= 1

    def rails_index_reverse_direction(self):
        if self.rails_index == 0:
            self.rails_index_increasing = True
        elif self.rails_index == self.num_rails - 1:
            self.rails_index_increasing = False

    def reset(self):
        self.rails_index = 0
        self.rails_index_increasing = True

class RailsFenceCipher:

    def __init__(self, message, num_rails):
        self.message = message
        self.num_rails = num_rails
        self.rails = Rails(num_rails)

    def encode(self):
        self.rails.populate_rails_zig_zag(self.message)
        return self.rails.to_string_linear()

    def decode(self):
        faulty_rails = Rails(self.num_rails)
        faulty_rails.populate_rails_zig_zag(self.message)
        rail_lengths = [len(rail) for rail in faulty_rails.rails]

        self.rails.populate_rails_linear(self.message, rail_lengths)
        return self.rails.to_string_zig_zag(self.message)


def encode(message, num_rails):
    return RailsFenceCipher(message, num_rails).encode()


def decode(message, num_rails):
    return RailsFenceCipher(message, num_rails).decode()
