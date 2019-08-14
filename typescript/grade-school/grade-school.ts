
export default class GradeSchool {

    roster: Map<number, string[]>

    constructor() {
        this.roster = new Map<number, string[]>()
    }

    public studentRoster(): ReadonlyMap<string, string[]> {
        const roster = Array.from(this.roster).map( ([key, value]) => [key.toString(), value] )
        return new Map(roster)
    }

    public addStudent(name: string, grade: number): void {
        const currentStudentsInGrade = this.roster.get(grade) || []
        const newStudentsInGrade = [...currentStudentsInGrade, name].sort()
        this.roster.set(grade, newStudentsInGrade )

    }
    public studentsInGrade(grade: number): string[] {
        const currentStudentsInGrade = this.roster.get(grade)
        if (currentStudentsInGrade) {
            return [...currentStudentsInGrade]
        }
        return []
    }

}
