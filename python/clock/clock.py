class Clock:
    def __init__(self, hours, mins):
        self.hours = hours
        self.mins = mins
        self.fixup()

    def __eq__(self, other):
        return self.hours == other.hours and self.mins == other.mins

    def __str__(self):
        return (self.format_hours() + ':' +
                self.format_mins())

    def add(self, additional_mins):
        self.mins += additional_mins
        self.fixup()
        return self

    def fixup(self):
        self.fixup_mins()
        self.fixup_hours()

    def fixup_hours(self):
        self.hours = int(self.hours % 24)

    def fixup_mins(self):
        self.hours += self.mins // 60
        self.mins = int(self.mins % 60)

    def format_hours(self):
        return str(self.hours).zfill(2)

    def format_mins(self):
        return str(self.mins).zfill(2)
