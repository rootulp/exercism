module.exports = class Luhn
  constructor: (num) ->
    @checkDigit = @getCheckDigit(num)
    @addends = @getAddends(num)
    @checksum = @getCheckSum(@addends)
    @valid = @isValid(@checksum)

  getCheckDigit: (num) ->
    num % 10

  getAddends: (num) ->
    nums = num.toString().split('')
    (@getAddend(i, nums) for i in [0..nums.length - 1]).reverse()

  getAddend: (i, nums) ->
    curr = parseInt(nums[nums.length - 1 - i])
    curr = if (i + 1) % 2 == 0 then curr * 2 else curr
    curr = if curr > 10 then curr - 9 else curr

  getCheckSum: (addends) ->
    addends.reduce (x, y) -> x + y

  isValid: (checksum) ->
    checksum % 10 == 0

  @create: (num) ->
    for i in [0..9]
      curr = Number(num + "" + i)
      return curr if new Luhn(curr).valid
