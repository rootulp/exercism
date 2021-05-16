export class ResistorColor {

  constructor (private readonly colors: string[]) {
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
