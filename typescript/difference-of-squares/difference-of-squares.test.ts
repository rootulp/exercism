import Squares from './difference-of-squares'

describe('Squares', () => {

  describe('up to 5', () => {
    const squares = new Squares(5)

    it('gets the square of sums', () => {
      expect(squares.squareOfSums).toBe(225)
    })

    it('gets the sum of squares', () => {
      expect(squares.sumOfSquares).toBe(55)
    })

    it('gets the difference', () => {
      expect(squares.difference).toBe(170)
    })

  })

  describe('up to 10', () => {
    const squares = new Squares(10)

    it('gets the square of sums', () => {
      expect(squares.squareOfSums).toBe(3025)
    })

    it('gets the sum of squares', () => {
      expect(squares.sumOfSquares).toBe(385)
    })

    it('gets the difference', () => {
      expect(squares.difference).toBe(2640)
    })

  })

  describe('up to 100', () => {
    const squares = new Squares(100)

    it('gets the square of sums', () => {
      expect(squares.squareOfSums).toBe(25502500)
    })

    it('gets the sum of squares', () => {
      expect(squares.sumOfSquares).toBe(338350)
    })

    it('gets the difference', () => {
      expect(squares.difference).toBe(25164150)
    })

  })

})
