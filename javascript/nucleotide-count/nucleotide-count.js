function dna(strand) {
  "use strict";

  strand = strand || "";
  var invalid_strand = /[^ATGC]/;
  var nucleotides = strand.split("");
  var counts = {
    A: 0,
    T: 0,
    G: 0,
    C: 0
  };

  if (invalid_strand.test(strand)) {
    throw "Invalid strand";
  }

  nucleotides.forEach(function(nucleotide) {
    counts[nucleotide] += 1;
  });

  function count(nucleotide) {
    return counts[nucleotide] || 0;
  }

  function histogram() {
    return counts;
  }

  return {
    count: count,
    histogram: histogram
  };
}

module.exports = dna;
