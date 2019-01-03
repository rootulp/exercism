from collections import defaultdict

class School:
    def __init__(self):
        self.db = defaultdict(list)

    def add_student(self, name, grade):
        self.db[grade].append(name)
        self.db[grade] = sorted(self.db[grade])

    def roster(self):
        pass

    def grade(self, grade_number):
        return self.db[grade_number]
