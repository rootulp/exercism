class Bob {
    static readonly RESPONSE_TO_YELLING = "Whoa, chill out!"
    static readonly RESPONSE_TO_QUESTION = "Sure."
    static readonly RESPONSE_TO_SILENCE = "Fine. Be that way!"
    public hey(input: string): string {
        if (this.is_yelling(input)) {
            return Bob.RESPONSE_TO_YELLING;
        } else if (this.is_question(input)) {
            return Bob.RESPONSE_TO_QUESTION;
        } else if (this.is_silence(input)) {
            return Bob.RESPONSE_TO_SILENCE;
        }
        return "Whatever.";
    }

    private is_yelling(input: string): boolean {
        return input === input.toUpperCase() && input !== input.toLowerCase();
    }

    private is_question(input: string): boolean {
        return input.endsWith("?");
    }

    private is_silence(input: string): boolean {
        return input.trim().length === 0;
    }
}

export default Bob
