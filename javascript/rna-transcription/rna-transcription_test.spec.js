var toRna = require('./rna-transcription');

describe("toRna()", function() {
  it("transcribes cytidine to guanosine", function() {
    expect(toRna('C')).toEqual('G');
  });

  it("transcribes guanosine to cytidine", function() {
    expect(toRna('G')).toEqual('C');
  });

  it("transcribes adenosine to uracil", function() {
    expect(toRna('A')).toEqual('U');
  });

  it("transcribes thymidine to adenosine", function() {
    expect(toRna('T')).toEqual('A');
  });

  it("transcribes all dna nucleotides to their rna complements", function() {
    expect(toRna('ACGTGGTCTTAA'))
        .toEqual('UGCACCAGAAUU');
  });
});
