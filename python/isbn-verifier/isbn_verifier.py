class IsbnVerifier(object):

    def __init__(self, string):
        self.string = string

    def is_valid(self):
        return True



def verify(isbn):
    return IsbnVerifier(isbn).is_valid()
