import Transcriptor from './rna-transcription';

describe('Transcriptor', () => {
  const transcriptor = new Transcriptor();

  test('transcribes cytosine to guanine', () => {
    expect(transcriptor.toRna('C')).toEqual('G');
  });

  test('transcribes guanine to cytosine', () => {
    expect(transcriptor.toRna('G')).toEqual('C');
  });

  test('transcribes adenine to uracil', () => {
    expect(transcriptor.toRna('A')).toEqual('U');
  });

  test('transcribes thymine to adenine', () => {
    expect(transcriptor.toRna('T')).toEqual('A');
  });

  test('transcribes all dna nucleotides to their rna complements', () => {
    expect(transcriptor.toRna('ACGTGGTCTTAA'))
        .toEqual('UGCACCAGAAUU');
  });

  test('correctly handles invalid input', () => {
    expect(() => transcriptor.toRna('U')).toThrow(
      new Error('Invalid input DNA.'),
    );
  });

  test('correctly handles completely invalid input', () => {
    expect(() => transcriptor.toRna('XXX')).toThrow(
      new Error('Invalid input DNA.'),
    );
  });

  test('correctly handles partially invalid input', () => {
    expect(() => transcriptor.toRna('ACGTXXXCTTAA')).toThrow(
      new Error('Invalid input DNA.'),
    );
  });
});
