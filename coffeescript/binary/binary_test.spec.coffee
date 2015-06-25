Binary = require './binary'

describe 'binary', ->

  it '1 is decimal 1', ->
    expect(new Binary('1').toDecimal()).toEqual 1

  it '10 is decimal 2', ->
    expect(new Binary('10').toDecimal()).toEqual 2

  it '11 is decimal 3', ->
    expect(new Binary('11').toDecimal()).toEqual 3

  it '100 is decimal 4', ->
    expect(new Binary('100').toDecimal()).toEqual 4

  it '1001 is decimal 9', ->
    expect(new Binary('1001').toDecimal()).toEqual 9

  it '11010 is decimal 26', ->
    expect(new Binary('11010').toDecimal()).toEqual 26

  it '10001101000 is decimal 1128', ->
    expect(new Binary('10001101000').toDecimal()).toEqual 1128

  it 'carrot is decimal 0', ->
    expect(new Binary('carrot').toDecimal()).toEqual 0

  it 'numeric but non-binary string is decimal 0', ->
    expect(new Binary('42').toDecimal()).toEqual 0
