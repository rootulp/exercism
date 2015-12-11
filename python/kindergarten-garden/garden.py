class Garden:

    DEFAULT_STUDENTS = ("Alice Bob Charlie David Eve Fred Ginny "
                        "Harriet Ileana Joseph Kincaid Larry").split()

    PLANTS = {'G': 'Grass',
              'C': 'Clover',
              'R': 'Radishes',
              'V': 'Violets'}

    def __init__(self, diagram, students = DEFAULT_STUDENTS):
        self.diagram = diagram
        self.rows = [list(row) for row in diagram.split()]
        self.plant_rows = [[self.PLANTS[c] for c in row] for row in self.rows]
        self.students = sorted(students)

    def plants(self, name):
        return self.plants_for_index(self.students.index(name))

    # Dislike how these are hardcoded indices
    def plants_for_index(self, i):
        return [self.plant_rows[0][i * 2],
                self.plant_rows[0][i * 2 + 1],
                self.plant_rows[1][i * 2],
                self.plant_rows[1][i * 2 + 1]]
