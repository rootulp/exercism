function Bob() {
  'use strict';

  this.hey = function(input) {
    if (!input.trim()) {
        return "Fine. Be that way!";
    } else if (input === input.toUpperCase() && input !== input.toLowerCase()){
        return 'Whoa, chill out!';
    } else if (input.slice(-1) === '?'){
        return 'Sure.';
    } else {
        return 'Whatever.';
    }
  }
};

module.exports = Bob;
