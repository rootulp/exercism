class Transcriptor {

    static readonly NUCLEOTIDE_COMPLEMENTS = new Map<string, string>([
        ['G', 'C'],
        ['C', 'G'],
        ['T', 'A'],
        ['A', 'U']
    ]);

    public toRna(strand: string) {
        return strand.split('').map(nucleotide => this.transcribeNucleotide(nucleotide)).join('')
    }

    private transcribeNucleotide(nucleotide: string) {
        return Transcriptor.NUCLEOTIDE_COMPLEMENTS.get(nucleotide)
    }
}

export default Transcriptor