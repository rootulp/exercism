class Gigasecond {
  date_of_birth: Date

  constructor(date_of_birth: Date) {
    this.date_of_birth = date_of_birth
  }

  public date() {
    return this.date_of_birth_plus_gigasecond()
  }

  private date_of_birth_plus_gigasecond() {
    return new Date(this.date_of_birth.getTime() + 1e12)
  }

}

export default Gigasecond