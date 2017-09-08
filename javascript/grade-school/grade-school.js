var School = function() {
  "use strict";

  var roster = {};

  function getRoster() {
    return roster;
  }

  function setRoster(name, grade) {
    if (roster[grade]) {
      roster[grade].push(name);
      roster[grade].sort();
    } else {
      roster[grade] = [name];
    }
  }

  function grade(grade) {
    if (roster[grade]) {
      return roster[grade];
    } else {
      return [];
    }
  }

  return {
    roster: getRoster,
    add: setRoster,
    grade: grade
  };
};

module.exports = School;
