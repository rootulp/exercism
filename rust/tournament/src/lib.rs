use std::fmt;

const POINTS_PER_WIN: u32 = 3;
const POINTS_PER_DRAW: u32 = 1;

#[derive(Debug, Clone)]
enum Outcome {
    Win,
    Loss,
    Draw,
}

#[derive(Debug, Clone)]
struct MatchResult {
    team_a: String,
    team_b: String,
    result: Outcome,
}

#[derive(Debug, Eq, Ord, PartialEq)]
struct Team {
    name: String,
    matches_won: u32,
    matches_lost: u32,
    matches_drawn: u32,
}

impl Team {
    fn new(name: String) -> Team {
        Team {
            name,
            matches_won: 0,
            matches_lost: 0,
            matches_drawn: 0,
        }
    }

    fn matches_played(&self) -> u32 {
        self.matches_won + self.matches_lost + self.matches_drawn
    }

    fn points(&self) -> u32 {
        (self.matches_won * POINTS_PER_WIN) + (self.matches_drawn * POINTS_PER_DRAW)
    }
}

impl fmt::Display for Team {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(
            f,
            "{: <30} |  {} |  {} |  {} |  {} |  {}",
            self.name,
            self.matches_played(),
            self.matches_won,
            self.matches_drawn,
            self.matches_lost,
            self.points()
        )
    }
}

impl PartialOrd for Team {
    fn partial_cmp(&self, other: &Team) -> Option<std::cmp::Ordering> {
        // Sort by points descending then alphabetically
        if other.points() != self.points() {
            Some(other.points().cmp(&self.points()))
        } else {
            Some(self.name.cmp(&other.name))
        }
    }
}

struct Tournament {
    teams: Vec<Team>,
}
impl Tournament {
    fn new(match_results: &str) -> Tournament {
        let mut tournament = Tournament { teams: Vec::new() };
        for line in match_results.lines() {
            let match_result = tournament.parse(line);
            tournament.tally(match_result);
        }

        tournament
    }

    fn tally(&mut self, match_result: MatchResult) -> () {
        self.init_team(match_result.team_a.as_str());
        self.init_team(match_result.team_b.as_str());

        match match_result.result {
            Outcome::Win => self.win(match_result.team_a, match_result.team_b),
            Outcome::Loss => self.win(match_result.team_b, match_result.team_a),
            Outcome::Draw => self.draw(match_result.team_a, match_result.team_b),
        }
    }

    fn init_team(&mut self, team_name: &str) -> () {
        if self.teams.iter().find(|t| t.name == team_name).is_none() {
            self.teams.push(Team::new(team_name.to_string()));
        }
    }

    fn find_team(&mut self, team_name: String) -> &mut Team {
        self.teams.iter_mut().find(|t| t.name == team_name).unwrap()
    }

    fn win(&mut self, team_a: String, team_b: String) -> () {
        let mut team_a = self.find_team(team_a);
        team_a.matches_won += 1;

        let mut team_b = self.find_team(team_b);
        team_b.matches_lost += 1;
    }

    fn draw(&mut self, team_a: String, team_b: String) -> () {
        let mut team_a = self.find_team(team_a);
        team_a.matches_drawn += 1;

        let mut team_b = self.find_team(team_b);
        team_b.matches_drawn += 1;
    }

    fn parse(&self, line: &str) -> MatchResult {
        let parts: Vec<&str> = line.split(';').collect();

        if parts.len() != 3 {
            panic!("Invalid match result: {}", line);
        }
        let result = match parts[2] {
            "win" => Outcome::Win,
            "loss" => Outcome::Loss,
            "draw" => Outcome::Draw,
            _ => panic!("Invalid match result outcome: {}", line),
        };

        MatchResult {
            team_a: parts[0].to_string(),
            team_b: parts[1].to_string(),
            result,
        }
    }

    fn results_table(&mut self) -> String {
        let mut results: Vec<String> = vec![self.header()];

        self.teams.sort();
        for team in self.teams.iter() {
            results.push(team.to_string());
        }

        results.join("\n")
    }

    fn header(&self) -> String {
        "Team                           | MP |  W |  D |  L |  P".to_string()
    }
}

pub fn tally(match_results: &str) -> String {
    let mut tournament = Tournament::new(match_results);
    tournament.results_table()
}
