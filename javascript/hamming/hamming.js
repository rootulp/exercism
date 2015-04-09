function compute_it(strand1, strand2) {
  'use strict';

  if (strand1.length !== strand2.length) {
    throw new Error('DNA strands must be of equal length.');
  }

  var differences = 0;

  for (var i = 0; i < strand1.length; i++) {
    if (strand1[i] !== strand2[i]) {
      differences += 1;
    }
  }

  return differences;
};

module.exports = {
  compute: function(strand1, strand2) {
   return  compute_it(strand1, strand2)
  }
};
