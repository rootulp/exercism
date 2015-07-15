Triangle = require './triangle'

describe "Triangle", ->
  it 'is equilateral with equal sides', ->
    triangle = new Triangle(2,2,2)
    expect(triangle.kind()).toBe 'equilateral'

  it 'is equilateral for larger triangle', ->
    triangle = new Triangle(10,10,10)
    expect(triangle.kind()).toBe 'equilateral'

  it 'is isosceles when last two sides equal', ->
    triangle = new Triangle(2,4,4)
    expect(triangle.kind()).toBe 'isosceles'

  it 'is isosceles when first and last equal', ->
    triangle = new Triangle(4,2,4)
    expect(triangle.kind()).toBe 'isosceles'

  it 'is isosceles when first two equal', ->
    triangle = new Triangle(4,4,2)
    expect(triangle.kind()).toBe 'isosceles'

  it 'is isosceles with arbitrary sides', ->
    triangle = new Triangle(10,10,2)
    expect(triangle.kind()).toBe 'isosceles'

  it 'is scalene when no sides equal', ->
    triangle = new Triangle(3,4,5)
    expect(triangle.kind()).toBe 'scalene'

  it 'is scalene when larger as well', ->
    triangle = new Triangle(10,11,12)
    expect(triangle.kind()).toBe 'scalene'

  it 'is scalene in decesending order', ->
    triangle = new Triangle(5,4,3)
    expect(triangle.kind()).toBe 'scalene'

  it 'is legal when very small', ->
    triangle = new Triangle(.03,.06,.04)
    expect(triangle.kind()).toBe 'scalene'

  it 'is illegal when a side is negative', ->
    expect(-> new Triangle(2,3,-5)).toThrow("negative sides are illegal")

  it 'is illegal when violating triangle inequality', ->
    expect(-> new Triangle(1,1,3)).toThrow("violation of triangle inequality")

  it 'is illegal when violating triangle inequality 2', ->
    expect(-> new Triangle(2,2,4)).toThrow("violation of triangle inequality")

  it 'is illegal when violating triangle inequality 3', ->
    expect(-> new Triangle(7,3,2)).toThrow("violation of triangle inequality")
