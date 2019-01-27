class Team:

    def __init__(self, name):
        self.name = name
        self.wins = 0
        self.losses = 0
        self.draws = 0

    def matches_played(self):
        return self.wins + self.losses + self.draws


class Tournament:

    def __init__(self, results):
        self.teams = {}
        self.results = results
        self.parse_results()

    def header(self):
        return 'Team                           | MP |  W |  D |  L |  P'

    def results_table(self):
        print("Teams {}".format(self.teams))
        return self.header()

    def parse_results(self):
        if (self.results == ''):
            return # Nothing to parse

        for result in self.results.split("\n"):
            team_a, team_b, outcome = result.split(";")
            if team_a not in self.teams:
                self.teams[team_a] = Team(team_a)
            if team_b not in self.teams:
                self.teams[team_b] = Team(team_b)
            self.tally_outcome(team_a, team_b, outcome)

    def tally_outcome(self, team_a, team_b, outcome):
        if outcome == 'win':
            self.teams[team_a].wins += 1
            self.teams[team_b].losses += 1
        if outcome == 'loss':
            self.teams[team_a].losses += 1
            self.teams[team_b].wins += 1
        if outcome == 'draw':
            self.teams[team_a].draws += 1
            self.teams[team_b].draws += 1


def tally(tournament_results):
    tournament = Tournament(tournament_results)
    return tournament.results_table()
