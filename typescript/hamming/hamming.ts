class Hamming {
    public compute(strand1: string, strand2:string) {
        let total = 0
        for (let i = 0; i < strand1.length; i++) {
            total += strand1.charAt(i) === strand2.charAt(i) ? 0 : 1
        }
        return total
    }

}

export default Hamming