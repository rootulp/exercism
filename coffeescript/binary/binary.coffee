module.exports = class Binary
  constructor: (@str) ->
    @reverse_str = @str.split("").reverse().join("")

  toDecimal: ->
    return 0 if (/[^01]/).test(@str)

    result = 0
    for i in [0..@reverse_str.length - 1]
      result += @reverse_str[i] * (2 ** i)

    result
