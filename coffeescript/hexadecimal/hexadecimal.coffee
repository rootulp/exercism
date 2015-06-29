module.exports = class Hexadecimal
  constructor: (str) ->
    @reverse_str = str.split('').reverse().join('')
    @CHAR_VALS =
      a: 10
      b: 11
      c: 12
      d: 13
      e: 14
      f: 15

  toDecimal: ->
    return 0 if (/[^a-f0-9]/).test(@reverse_str)

    result = 0
    for i in [0..@reverse_str.length - 1]
      result += (16 ** i) * @valFor(@reverse_str[i])
    result

  valFor: (symbol) ->
    if @CHAR_VALS[symbol] then @CHAR_VALS[symbol] else symbol
