
export default class Clock {
    private hours: number;
    private minutes?: number;

    constructor(hours: number, minutes?: number) {
        this.minutes = minutes ? minutes % 60 : 0;
        const additionalHours = minutes ? Math.floor(minutes / 60) : 0;
        this.hours = Clock.mod((hours + additionalHours), 24);
    }

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

    // Can't use % because JavaScript module bug
    // https://web.archive.org/web/20090717035140if_/javascript.about.com/od/problemsolving/a/modulobug.htm
    private static mod(x: number, y: number): number {
        return ((x % y) + y) % y;
    }
}
