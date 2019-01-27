class Tournament:

    def __init__(self, results):
        self.results = results

    def header(self):
        return 'Team                           | MP |  W |  D |  L |  P'

    def results_table(self):
        if (self.results == ''):
            return self.header()

def tally(tournament_results):
    tournament = Tournament(tournament_results)
    return tournament.results_table()
