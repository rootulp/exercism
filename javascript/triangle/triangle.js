function Triangle(sideA, sideB, sideC) {
  "use strict";

  this.kind = function() {
    if (this.error()) {
      throw Error;
    } else if (this.equilateral()) {
      return "equilateral";
    } else if (this.isosceles()) {
      return "isosceles";
    } else {
      return "scalene";
    }
  };

  this.equilateral = function() {
    return sideA === sideB && sideB === sideC;
  };

  this.isosceles = function() {
    return sideA === sideB || sideB === sideC || sideA === sideC;
  };

  this.error = function() {
    return (
      this.inequality(sideA, sideB, sideC) ||
      this.inequality(sideB, sideC, sideA) ||
      this.inequality(sideA, sideC, sideB)
    );
  };

  this.inequality = function(x, y, z) {
    return x + y <= z;
  };
}

module.exports = Triangle;
