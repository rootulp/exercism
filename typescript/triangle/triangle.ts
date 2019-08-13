export default class Triangle {

    sides: number[]

    constructor(...sides: number[]) {
        this.sides = sides
    }

    public kind(): string {
        if (this.is_equilateral()){
            return "equilateral"
        }
        return "error: not a valid triangle"
    }

    private is_equilateral(): boolean {
        return true;
    }
}
