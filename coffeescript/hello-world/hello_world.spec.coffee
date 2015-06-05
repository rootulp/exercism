HelloWorld = require "./hello_world"
describe "HelloWorld", ->
  hello_world = new HelloWorld()
  it "given no name", ->
    result = hello_world.hello()
    expect(result).toEqual "Hello, World!"
  it "given the name 'Alice'", ->
    result = hello_world.hello 'Alice'
    expect(result).toEqual "Hello, Alice!"
  it "given the name 'Bob'", ->
    result = hello_world.hello 'Bob'
    expect(result).toEqual "Hello, Bob!"
