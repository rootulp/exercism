export default class Triangle {

    sides: number[]

    constructor(...sides: number[]) {
        this.sides = sides
    }

    public kind(): string {
        if (this.is_equilateral()){
            return "equilateral"
        } else if (this.is_isosceles()) {
            return "isosceles"
        } else if (this.is_scalene()) {
            return "scalene"
        }
        return "error: not a valid triangle"
    }

    private is_equilateral(): boolean {
        return this.sides[0] === this.sides[1]  && this.sides[1] === this.sides[2]
    }

    private is_isosceles(): boolean {
        return this.sides[0] === this.sides[1] ||
               this.sides[1] === this.sides[2] ||
               this.sides[0] === this.sides[2]
    }

    private is_scalene(): boolean {
        return this.sides[0] !== this.sides[1] &&
               this.sides[1] !== this.sides[2] &&
               this.sides[0] !== this.sides[2]
    }
}
