class School:

    def __init__(self, name):
        self.name = name
        self.db = {}

    def add(self, student, grade):
        if grade not in self.db:
            self.initialize_grade(grade)
        self.db[grade].add(student)

    def grade(self, grade):
        return self.db[grade] if grade in self.db else set()

    def sort(self):
        students = []
        for grade in self.db.keys():
            students.append((grade, tuple(self.db[grade])))
        return students

    def initialize_grade(self, grade):
        self.db[grade] = set()
