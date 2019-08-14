export default class Triangle {

    sides: number[]

    constructor(...sides: number[]) {
        this.sides = sides
    }

    public kind(): string {
        if (this.is_illegal()) {
            throw('Illegal triangle sides');
        } else if (this.is_equilateral()){
            return 'equilateral'
        } else if (this.is_isosceles()) {
            return 'isosceles'
        }
        return 'scalene'
    }

    private is_illegal(): boolean {
        return this.is_any_side_illegal() || this.is_triangle_inequality_illegal()
    }

    private is_any_side_illegal(): boolean {
        return this.sides.some(side => side <= 0);
    }

    private is_triangle_inequality_illegal(): boolean {
        return this.sides[0] + this.sides[1] < this.sides[2] ||
               this.sides[0] + this.sides[2] < this.sides[1] ||
               this.sides[1] + this.sides[2] < this.sides[0]
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
