
const bandToNumber: {[key: string]: number} = {
  'black': 0,
  'brown': 1,
  'red': 2,
  'orange': 3,
  'yellow': 4,
  'green': 5,
  'blue': 6,
  'violet': 7,
  'grey': 8,
  'white': 9,
}

export function decodedResistorValue(bands: string[]): string {
  const number = decodeResistorDuo(bands.slice(0, 2))
  console.log(number)
  const result = number * 10 ** decode(bands[2])
  return `${result} ohms`
}

function decodeResistorDuo(bands: string[]): number {
  const band1 = decode(bands[0])
  const band2 = decode(bands[1])
  return parseInt(`${band1}${band2}`)
}

function decode(band: string): number {
  return bandToNumber[band]
}
