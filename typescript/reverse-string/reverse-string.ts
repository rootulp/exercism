class ReverseString {
    static reverse(input: String): String {
        const reversed = input.split("").reverse().join("")
        return reversed
    }
}

export default ReverseString
