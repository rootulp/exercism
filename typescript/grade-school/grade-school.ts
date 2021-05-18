
export default class GradeSchool {

    private roster: Map<number, string[]>

    constructor() {
        this.roster = new Map<number, string[]>()
    }

    public studentRoster(): ReadonlyMap<string, string[]> {
        const result = new Map<string, string[]>();
        this.roster.forEach((students, grade) => {
            result.set(grade.toString(), [...students])
        })
        return result;
    }

    public addStudent(name: string, grade: number): void {
        this.removeStudentIfAlreadyExists(name);
        const students = this.studentsInGrade(grade);
        const newStudents = [...students, name].sort()
        this.roster.set(grade, newStudents)
    }

    public studentsInGrade(grade: number): string[] {
        const students = this.roster.get(grade)
        if (students) {
            return [...students]
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
