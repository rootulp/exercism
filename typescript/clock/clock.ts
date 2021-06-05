
export default class Clock {
    constructor(private readonly hours: number, private readonly minutes?: number) {}

    public toString() {
        const formattedHours = this.hours.toString().padStart(2, '0')
        const formattedMinutes = this.minutes ? this.minutes.toString().padStart(2, '0') : '00'
        return [formattedHours, formattedMinutes].join(':')
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
