export default class Triangle {

    sides: number[]

    constructor(...sides: number[]) {
        this.sides = sides
    }

    public kind(): string {
        if (this.is_equilateral()){
            return 'equilateral'
        } else if (this.is_isosceles()) {
            return 'isosceles'
        }
        return 'scalene'
    }

    private is_equilateral(): boolean {
        return this.sides[0] === this.sides[1]  && this.sides[1] === this.sides[2]
    }

    private is_isosceles(): boolean {
        return this.sides[0] === this.sides[1] ||
               this.sides[1] === this.sides[2] ||
               this.sides[0] === this.sides[2]
    }

}
