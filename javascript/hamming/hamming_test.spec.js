var compute = require('./hamming').compute;

describe('Hamming', function () {

  it('no difference between identical strands', function () {
    expect(compute('A', 'A')).toEqual(0);
  });

  it('complete hamming distance for single nucleotide strand', function () {
    expect(compute('A','G')).toEqual(1);
  });

  it('complete hamming distance for small strand', function () {
    expect(compute('AG','CT')).toEqual(2);
  });

  it('small hamming distance', function () {
    expect(compute('AT','CT')).toEqual(1);
  });

  it('small hamming distance in longer strand', function () {
    expect(compute('GGACG', 'GGTCG')).toEqual(1);
  });

  it('large hamming distance', function () {
    expect(compute('GATACA', 'GCATAA')).toEqual(4);
  });

  it('hamming distance in very long strand', function () {
    expect(compute('GGACGGATTCTG', 'AGGACGGATTCT')).toEqual(9);
  });

  it('throws error when strands are not equal length', function() {
    expect(function() { compute('GGACGGATTCTG', 'AGGAC'); }).toThrow(
      new Error('DNA strands must be of equal length.')
    );
  });

});
