pub fn plants(diagram: &str, name: &str) -> Vec<&'static str> {
    let student = parse_student(name);
    let mut result = vec![];
    let lines = diagram.split("\n");
    for line in lines {
        for index in student.indices() {
            let encoding = line.chars().nth(index).unwrap();
            let plant = get_plant(encoding);
            result.push(plant.string());
        }
    };
    result
}

enum Plant {
    Grass,
    Clover,
    Radish,
    Violet
}

impl Plant {
    fn string(&self) -> &'static str {
        match self {
            Plant::Grass => "grass",
            Plant::Clover => "clover",
            Plant::Radish => "radishes",
            Plant::Violet => "violets",
        }
    }
}

fn get_plant(encoding: char) -> Plant {
    return match encoding {
        'G' => Plant::Grass,
        'C' => Plant::Clover,
        'R' => Plant::Radish,
        'V' => Plant::Violet,
        _ => Plant::Grass // TODO: replace this with an error
    }
}


enum Student {
    Alice,
    Bob,
    Charlie,
    David,
    Eve,
    Fred,
    Ginny,
    Harriet,
    Ileana,
    Joseph,
    Kincaid,
    Larry
}

fn parse_student(name: &str) -> Student {
    return match name {
        "Alice" => Student::Alice,
        "Bob" => Student::Bob,
        "Charlie" => Student::Charlie,
        "David" => Student::David,
        "Eve" => Student::Eve,
        "Fred" => Student::Fred,
        "Ginny" => Student::Ginny,
        "Harriet" => Student::Harriet,
        "Ileana" => Student::Ileana,
        "Joseph" => Student::Joseph,
        "Kincaid" => Student::Kincaid,
        "Larry" => Student::Larry,
        _ => panic!("Invalid student name")
    }
}

impl Student {
    fn indices(&self) -> Vec<usize> {
        return match self {
            Student::Alice => vec![0, 1],
            Student::Bob => vec![2, 3],
            Student::Charlie => vec![4, 5],
            Student::David => vec![6, 7],
            Student::Eve => vec![8, 9],
            Student::Fred => vec![10, 11],
            Student::Ginny => vec![12, 13],
            Student::Harriet => vec![14, 15],
            Student::Ileana => vec![16, 17],
            Student::Joseph => vec![18, 19],
            Student::Kincaid => vec![20, 21],
            Student::Larry => vec![22, 23],
        }
    }
}
