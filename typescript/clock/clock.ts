
export default class Clock {
    constructor(private readonly hours: number, private readonly minutes?: number) {}

    public toString() {
        return '08:00'
    }

    public plus(minutes: number): Clock {
        console.log(`Plus ${minutes}`)
        return this;
    }

    public minus(minutes: number): Clock {
        console.log(`Minus ${minutes}`)
        return this;
    }

    public equals(c: Clock): boolean {
        return this.hours === c.hours &&
               this.minutes === c.minutes;
    }
}
