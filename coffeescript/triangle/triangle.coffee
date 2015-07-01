module.exports = class Triangle
  constructor: (@a, @b, @c) ->
    @sides = [@a, @b, @c]

  kind: ->
    # @isInvalid()
    return 'equilateral' if @isEquilateral()
    return 'isosceles'   if @isIsosceles()
    'scalene'

  isInvalid: ->
    # if @sides.min < 0  then throw type: 'a', msg: 'negative sides are illegal'

  isIsosceles: ->
    @a == @b ||
    @a == @c ||
    @b == @c

  isEquilateral: ->
    @a == @b &&
    @a == @c &&
    @b == @c

