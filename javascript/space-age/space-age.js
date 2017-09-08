function SpaceAge(seconds) {
  "use script";

  var earthTo = {
    Mercury: 0.2408467,
    Venus: 0.61519726,
    Mars: 1.8808158,
    Jupiter: 11.862615,
    Saturn: 29.447498,
    Uranus: 84.016846,
    Neptune: 164.79132
  };

  function chompFloat(val) {
    return parseFloat(val.toFixed(2));
  }

  function onEarth() {
    return seconds / 31557600;
  }

  this.seconds = seconds;

  this.onEarth = function() {
    return chompFloat(onEarth());
  };

  this.onMercury = function() {
    return chompFloat(onEarth() / earthTo["Mercury"]);
  };

  this.onVenus = function() {
    return chompFloat(onEarth() / earthTo["Venus"]);
  };

  this.onMars = function() {
    return chompFloat(onEarth() / earthTo["Mars"]);
  };

  this.onJupiter = function() {
    return chompFloat(onEarth() / earthTo["Jupiter"]);
  };

  this.onSaturn = function() {
    return chompFloat(onEarth() / earthTo["Saturn"]);
  };

  this.onUranus = function() {
    return chompFloat(onEarth() / earthTo["Uranus"]);
  };

  this.onNeptune = function() {
    return chompFloat(onEarth() / earthTo["Neptune"]);
  };
}

module.exports = SpaceAge;
