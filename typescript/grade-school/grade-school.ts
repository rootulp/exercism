
export default class GradeSchool {

    private roster: Map<number, string[]>

    constructor() {
        this.roster = new Map<number, string[]>()
    }

    public studentRoster(): ReadonlyMap<string, string[]> {
        const result = new Map<string, string[]>();
        this.roster.forEach((students, grade) => {
            // Use array.slice() to prevent modifications of roster outside this
            // module
            result.set(grade.toString(), students.slice())
        })
        return result;
    }

    public addStudent(name: string, grade: number): void {
        this.removeStudentIfAlreadyExists(name);
        const currentStudentsInGrade = this.roster.get(grade) || []
        const newStudentsInGrade = [...currentStudentsInGrade, name].sort()
        this.roster.set(grade, newStudentsInGrade)
    }

    public studentsInGrade(grade: number): string[] {
        const currentStudentsInGrade = this.roster.get(grade)
        if (currentStudentsInGrade) {
            return [...currentStudentsInGrade]
        }
        return []
    }

    private removeStudentIfAlreadyExists(name: string) {
        let newRoster = new Map<number, string[]>();

        this.roster.forEach((students, grade) => {
            const newStudents = students.filter(student => student !== name)
            newRoster.set(grade, newStudents);
        })

        this.roster = newRoster;
    }

}
