module.exports = class Triangle
  constructor: (@a, @b, @c) ->
    @sides = [@a, @b, @c]
    @isNegative()
    @isInequality()

  kind: ->
    return 'equilateral' if @isEquilateral()
    return 'isosceles'   if @isIsosceles()
    'scalene'

  isIsosceles: ->
    @a == @b ||
    @a == @c ||
    @b == @c

  isEquilateral: ->
    @a == @b &&
    @a == @c &&
    @b == @c

  isNegative: ->
    for side in @sides
      throw 'negative sides are illegal' if side < 0

  isInequality: ->
    if @a + @b <= @c ||
       @a + @c <= @b ||
       @b + @c <= @a
      throw 'violation of triangle inequality'
