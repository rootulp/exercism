class Transcriptor {

    static readonly NUCLEOTIDE_COMPLEMENTS = new Map<string, string>([
        ['G', 'C'],
        ['C', 'G'],
        ['T', 'A'],
        ['A', 'U']
    ]);

    toRna(strand: string) {
        return Transcriptor.NUCLEOTIDE_COMPLEMENTS.get(strand)
    }
}

export default Transcriptor