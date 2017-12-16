import {keep, discard} from './strain'

describe('strain', () => {
  it('keeps on empty array returns empty array', () => {
    expect(keep<number>([], (e: number) => e < 10)).toEqual([])
  })

  it('keeps everything ', () => {
    expect(keep<number>([1, 2, 3], (e: number) => e < 10)).toEqual([1, 2, 3])
  })

  it('keeps first and last', () => {
    expect(keep<number>([1, 2, 3], (e: number) => e % 2 === 1)).toEqual([1, 3])
  })

  it('keeps neither first nor last', () => {
    expect(keep<number>([1, 2, 3, 4, 5], (e: number) => e % 2 === 0)).toEqual([2, 4])
  })

  it('keeps strings', () => {
    const words = 'apple zebra banana zombies cherimoya zelot'.split(' ')
    const result = keep<string>(words, (word: string) => word.indexOf('z') === 0)
    expect(result).toEqual('zebra zombies zelot'.split(' '))
  })

  it('keeps arrays', () => {
    const rows = [
      [1, 2, 3],
      [5, 5, 5],
      [5, 1, 2],
      [2, 1, 2],
      [1, 5, 2],
      [2, 2, 1],
      [1, 2, 5],
    ]
    const result = keep<number[]>(rows, (row: number[]) => row.indexOf(5) > -1)
    expect(result).toEqual([[5, 5, 5], [5, 1, 2], [1, 5, 2], [1, 2, 5]])
  })

  it('empty discard', () => {
    expect(discard<number>([], (e: number) => e < 10)).toEqual([])
  })

  it('discards nothing', () => {
    expect(discard<number>([1, 2, 3], (e: number) => e > 10)).toEqual([1, 2, 3])
  })

  it('discards first and last', () => {
    expect(discard<number>([1, 2, 3], (e: number) => e % 2 === 1)).toEqual([2])
  })

  it('discards neither first nor last', () => {
    const result = discard<number>([1, 2, 3, 4, 5], (e: number) => e % 2 === 0)
    expect(result).toEqual([1, 3, 5])
  })

  it('discards strings', () => {
    const words = 'apple zebra banana zombies cherimoya zelot'.split(' ')
    const result = discard<string>(words, (word: string) => word.indexOf('z') === 0)
    expect(result).toEqual('apple banana cherimoya'.split(' '))
  })

  it('discards arrays', () => {
    const rows = [
      [1, 2, 3],
      [5, 5, 5],
      [5, 1, 2],
      [2, 1, 2],
      [1, 5, 2],
      [2, 2, 1],
      [1, 2, 5],
    ]
    const result = discard<number[]>(rows, (row: number[]) => row.indexOf(5) > -1)
    expect(result).toEqual([[1, 2, 3], [2, 1, 2], [2, 2, 1]])
  })
})
