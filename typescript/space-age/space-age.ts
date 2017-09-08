class SpaceAge {

  public seconds: number

    private ORBITAL_PERIOD_IN_EARTH_YEARS = Object.freeze({
        Mercury: 0.2408467,
        Venus: 0.61519726,
        Mars: 1.8808158,
        Jupiter: 11.862615,
        Saturn: 29.447498,
        Uranus: 84.016846,
        Neptune: 164.79132
    })

    private SECONDS_IN_EARTH_YEAR = 31557600

  constructor(seconds: number) {
    this.seconds = seconds
  }

  private earthYears(): number {
    return this.seconds / this.SECONDS_IN_EARTH_YEAR
  }

  public onEarth(): number {
    return SpaceAge.roundToTwoDecimalPlaces(this.earthYears())
  }

  public onMercury(): number {
    return SpaceAge.roundToTwoDecimalPlaces(this.earthYears() / this.ORBITAL_PERIOD_IN_EARTH_YEARS.Mercury)
  }

  public onVenus(): number {
    return SpaceAge.roundToTwoDecimalPlaces(this.earthYears() / this.ORBITAL_PERIOD_IN_EARTH_YEARS.Venus)
  }

  public onMars(): number {
    return SpaceAge.roundToTwoDecimalPlaces(this.earthYears() / this.ORBITAL_PERIOD_IN_EARTH_YEARS.Mars)
  }

  public onJupiter(): number {
    return SpaceAge.roundToTwoDecimalPlaces(this.earthYears() / this.ORBITAL_PERIOD_IN_EARTH_YEARS.Jupiter)
  }

  public onSaturn(): number {
    return SpaceAge.roundToTwoDecimalPlaces(this.earthYears() / this.ORBITAL_PERIOD_IN_EARTH_YEARS.Saturn)
  }

  public onUranus(): number {
    return SpaceAge.roundToTwoDecimalPlaces(this.earthYears() / this.ORBITAL_PERIOD_IN_EARTH_YEARS.Uranus)
  }

  public onNeptune(): number {
    return SpaceAge.roundToTwoDecimalPlaces(this.earthYears() / this.ORBITAL_PERIOD_IN_EARTH_YEARS.Neptune)
  }

  private static roundToTwoDecimalPlaces(num: number): number {
    return Math.round(num * 100) / 100
  }

}

export default SpaceAge