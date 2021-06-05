interface Event {
    binary: number;
    action: string;
}

export default class HandShake {

    private static REVERSE_EVENT: Event = { binary: 10000, action: '' }
    private static EVENTS: Event[] = [
        { binary: 1000, action: 'jump' },
        { binary: 100, action: 'close your eyes' },
        { binary: 10, action: 'double blink' },
        { binary: 1, action: 'wink' }
    ]

    constructor(private readonly number: number) {}

    public commands(): string[] {
        let result: string[] = []
        let binary = this.binary();
        let shouldReverse: boolean = false;

        if (binary >= HandShake.REVERSE_EVENT.binary) {
            shouldReverse = true;
            binary -= HandShake.REVERSE_EVENT.binary
        }

        HandShake.EVENTS.forEach(event => {
            if (binary >= event.binary) {
                binary -= event.binary
                result.unshift(event.action);
            }
        })

        if (shouldReverse) {
            return result.reverse();
        }
        return result;
    }

    private binary(): number {
        return parseInt(this.number.toString(2));
    }
}
