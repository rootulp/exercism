export class ResistorColor {

  private colors: string[];

  constructor (colors: string[]) {
    if (colors.length < 2) {
      throw Error('At least two colors need to be present');
    }

    // Ignore additional colors
    this.colors = colors.splice(0, 2);
  }
  public value(): number {
    let result = ""
    for (let color of this.colors) {
      result += colorCode(color);
    }
    return Number.parseInt(result)
  }
}

export const colorCode = (color: string) => {
  return COLORS.indexOf(color)
}

export const COLORS = [
  'black',
  'brown',
  'red',
  'orange',
  'yellow',
  'green',
  'blue',
  'violet',
  'grey',
  'white',
]
