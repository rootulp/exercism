interface Event {
    decimal: number;
    action: string;
}
export default class HandShake {
    private static EVENTS: Event[] = [
        { decimal: 1000, action: 'jump' },
        { decimal: 100, action: 'close your eyes' },
        { decimal: 10, action: 'double blink' },
        { decimal: 1, action: 'wink' }
    ]
    constructor(private readonly number: number) {}
    private binary(): number {
        return parseInt(this.number.toString(2));
    }
    public commands(): string[] {
        let total = this.binary();
        console.log(`total: ${total}, number: ${this.number}`);
        const result: string[] = []
        HandShake.EVENTS.forEach(event => {
            if (total >= event.decimal) {
                total -= event.decimal
                result.unshift(event.action);
            }
        })
        return result;
    }
}
