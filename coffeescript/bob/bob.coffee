class Bob
  hey: (input) ->
    if !input.trim()
      "Fine. Be that way!"
    else if input is input.toUpperCase() and input isnt input.toLowerCase()
      'Whoa, chill out!'
    else if input.slice(-1) is '?'
      'Sure.'
    else
      'Whatever.'

module.exports = Bob
