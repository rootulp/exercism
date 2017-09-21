class Year {

  constructor(year) {
    this.year = year;
  }

  isLeap() {
    if (this.year % 400 === 0) {
      return true;
    } else if (this.year % 100 === 0) {
      return false;
    } else if (this.year % 4 === 0) {
      return true;
    }
    return false;
  }
}

export default Year;
