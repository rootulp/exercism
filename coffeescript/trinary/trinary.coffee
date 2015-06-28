module.exports = class Trinary
  constructor: (@str) ->
    @reverse_str = @str.split("").reverse().join("")

  toDecimal: ->
    return 0 if (/[^012]/).test(@str)

    result = 0
    for i in [0..@reverse_str.length - 1]
      result += @reverse_str[i] * (3 ** i)

    result
