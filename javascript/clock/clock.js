function at(hours, mins) {
  'use strict';
  var time = (60 * hours) + (mins || 0);

  var plus = function(mins) {
    time = (time + mins) % 1440;
    return clock;
  }

  var minus = function(mins) {
    time -= mins;
    // Find better way to wrap around midnight
    while (time < 0) { time += 1440 }
    return clock;
  }

  var equals = function(other) {
    return getTime() === other.getTime();
  }

  var toString = function() {
    return pad(getHours()) + ':' + pad(getMins());
  }

  var pad = function(n) {
    return (n < 10) ? ('0' + n) : n;
  }

  var getHours = function() {
    return parseInt(time / 60);
  }

  var getMins = function() {
    return parseInt(time % 60);
  }

  var getTime = function() {
    return time;
  }

  var clock = {
    plus: plus,
    minus: minus,
    equals: equals,
    getTime: getTime,
    toString: toString
  };

  return Object.create(clock);
}

exports.at = at;
