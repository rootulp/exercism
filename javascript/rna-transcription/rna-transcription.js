function toRna(strand) {
  'use strict';

  var nucleotides = strand.split("");
  var pairs = {
    A: 'U',
    T: 'A',
    G: 'C',
    C: 'G'
  }

  function pair_for(nucleotide) {
    return pairs[nucleotide];
  }

  return nucleotides.map(pair_for).join('')
}

module.exports = toRna
