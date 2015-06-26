module.exports = class Binary
  constructor: (@str) ->
    @invalid = (/^01/).test(@str)
    @reverse_str = @str.split("").reverse().join("")

  toDecimal: ->
    return 0 if @invalid

    result = 0
    for i in [0..@reverse_str.length - 1]
      result += 2 ** i if @reverse_str[i] == "1"

    result
