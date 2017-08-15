class Transcriptor {

    static readonly NUCLEOTIDE_COMPLEMENTS = new Map<string, string>([
        ['G', 'C'],
        ['C', 'G'],
        ['T', 'A'],
        ['A', 'U']
    ]);

    public toRna(strand: string) {
        if (this.invalidInputData(strand)) {
            throw 'Invalid input DNA.'
        }
        return this.transcribeStrand(strand)
    }

    private transcribeStrand(strand: string) {
        return strand.split('').map(nucleotide => this.transcribeNucleotide(nucleotide)).join('')
    }

    private transcribeNucleotide(nucleotide: string) {
        return Transcriptor.NUCLEOTIDE_COMPLEMENTS.get(nucleotide)
    }

    private invalidInputData(strand: string): boolean {
        for (let nucleotide of strand) {
            if (!Transcriptor.NUCLEOTIDE_COMPLEMENTS.has(nucleotide)) {
                return true
            }
        }
        return false
    }
}

export default Transcriptor