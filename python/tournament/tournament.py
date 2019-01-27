class Team:

    def __init__(self, name):
        self.name = name
        self.wins = 0
        self.losses = 0
        self.draws = 0

    def matches_played(self):
        return self.wins + self.losses + self.draws

    def points(self):
        return (self.wins * 3) + self.draws

    def __str__(self):
        return '{:<30} | {:^3}| {:^3}| {:^3}| {:^3}| {:>2}'.format(self.name, self.matches_played(), self.wins, self.draws, self.losses, self.points())


class Tournament:

    WIN = 'win'
    DRAW = 'draw'
    LOSS = 'loss'
    RESULT_SEPERATOR = ';'
    COLUMN_HEADERS = ['Team', 'MP', 'W', 'D', 'L', 'P']

    def __init__(self, results):
        self.teams = {}
        if results:
            self.parse(results)

    def header(self):
        return '{:<30} | {:^3}| {:^3}| {:^3}| {:^3}| {:>2}'.format(*self.COLUMN_HEADERS)

    def results_table(self):
        table = [self.header()]
        for team in self.sorted_teams():
            table.append(str(team))
        return "\n".join(table)

    def sorted_teams(self):
        alphabetic = sorted(self.teams.values(), key=lambda team: team.name)
        alphabetic_descending_points = sorted(alphabetic, key=lambda team: team.points(), reverse=True)
        return alphabetic_descending_points

    def parse(self, results):
        for result in results.split("\n"):
            team_a, team_b, outcome = result.split(self.RESULT_SEPERATOR)
            self.maybe_initialize_teams(team_a, team_b)
            self.tally_outcome(team_a, team_b, outcome)

    def tally_outcome(self, team_a, team_b, outcome):
        if outcome == self.WIN:
            self.teams[team_a].wins += 1
            self.teams[team_b].losses += 1
        if outcome == self.LOSS:
            self.teams[team_a].losses += 1
            self.teams[team_b].wins += 1
        if outcome == self.DRAW:
            self.teams[team_a].draws += 1
            self.teams[team_b].draws += 1

    def maybe_initialize_teams(self, team_a, team_b):
        self.maybe_initialize_team(team_a)
        self.maybe_initialize_team(team_b)

    def maybe_initialize_team(self, name):
        if name not in self.teams:
            self.teams[name] = Team(name)

def tally(tournament_results):
    tournament = Tournament(tournament_results)
    return tournament.results_table()
