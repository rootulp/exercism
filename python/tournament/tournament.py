class Team:

    POINTS_PER_WIN = 3
    POINTS_PER_DRAW = 1

    def __init__(self, name):
        self.name = name
        self.wins = 0
        self.draws = 0
        self.losses = 0

    @property
    def matches_played(self):
        return self.wins + self.losses + self.draws

    @property
    def points(self):
        return (self.wins * self.POINTS_PER_WIN) + (self.draws * self.POINTS_PER_DRAW)

    def tally_win(self):
        self.wins += 1

    def tally_draw(self):
        self.draws += 1

    def tally_loss(self):
        self.losses += 1

    def __str__(self):
        return '{:<30} | {:^3}| {:^3}| {:^3}| {:^3}| {:>2}'.format(self.name, self.matches_played, self.wins, self.draws, self.losses, self.points)


class Tournament:

    WIN = 'win'
    DRAW = 'draw'
    LOSS = 'loss'
    RESULT_SEPERATOR = ';'
    COLUMN_HEADERS = ['Team', 'MP', 'W', 'D', 'L', 'P']

    def __init__(self, results):
        self._teams = {}
        if results:
            self.parse(results)

    def results_table(self):
        table = [self.table_header()]
        for team in self.sorted_teams():
            table.append(str(team))
        return "\n".join(table)

    def sorted_teams(self):
        alphabetic = sorted(self._teams.values(), key=lambda team: team.name)
        alphabetic_descending_points = sorted(alphabetic, key=lambda team: team.points, reverse=True)
        return alphabetic_descending_points

    def parse(self, results):
        for result in results.split("\n"):
            team_a, team_b, outcome = result.split(self.RESULT_SEPERATOR)
            self.maybe_initialize_teams(team_a, team_b)
            self.tally_outcome(team_a, team_b, outcome)

    def tally_outcome(self, team_a, team_b, outcome):
        if outcome == self.WIN:
            self.tally_win(team_a, team_b)
        if outcome == self.LOSS:
            self.tally_loss(team_a, team_b)
        if outcome == self.DRAW:
            self.tally_draw(team_a, team_b)

    def tally_win(self, winner, loser):
        self._teams[winner].tally_win()
        self._teams[loser].tally_loss()

    def tally_loss(self, loser, winner):
        self._teams[loser].tally_loss()
        self._teams[winner].tally_win()

    def tally_draw(self, team_a, team_b):
        self._teams[team_a].tally_draw()
        self._teams[team_b].tally_draw()

    def maybe_initialize_teams(self, team_a, team_b):
        self.maybe_initialize_team(team_a)
        self.maybe_initialize_team(team_b)

    def maybe_initialize_team(self, name):
        if name not in self._teams:
            self._teams[name] = Team(name)

    def table_header(self):
        return '{:<30} | {:^3}| {:^3}| {:^3}| {:^3}| {:>2}'.format(*self.COLUMN_HEADERS)

def tally(results):
    return Tournament(results).results_table()
