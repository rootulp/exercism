var dna = require('./nucleotide-count');

describe('DNA', function() {

  it('Empty DNA strand has no adenosine', function() {
    expect(dna().count('A')).toEqual(0);
  });

  it('Repetitive cytidine gets counted', function() {
    expect(dna('CCCCC').count('C')).toEqual(5);
  });

  it('Counts only thymidine', function() {
    expect(dna('GGGGGTAACCCGG').count('T')).toEqual(1);
  });

  it('Counts a nucleotide only once', function() {
    var acid = dna('CGATTGGG');
    acid.count('T');
    acid.count('T');
    expect(acid.count('T')).toEqual(2);
  });

  it('Empty DNS strand has no nucleotides', function() {
    var expected = {A: 0, T: 0, C: 0, G: 0};
    expect(dna().histogram()).toEqual(expected);
  });

  it('Repetitive sequence has only guanosine', function() {
    var expected = {A: 0, T: 0, C: 0, G: 8};
    expect(dna('GGGGGGGG').histogram()).toEqual(expected);
  });

  it('Counts all nucleotides', function() {
    var strand = 'AGCTTTTCATTCTGACTGCAACGGGCAATATGTCTCTGTGTGGATTAAAAAAAGAGTGTCTGATAGCAGC';
    var expected = {A: 20, T: 21, C: 12, G: 17};
    expect(dna(strand).histogram()).toEqual(expected);
  });

  it('Validates DNA', function() {
    expect(dna.bind(null, 'JOHNNYAPPLESEED')).toThrow();
  });

});
