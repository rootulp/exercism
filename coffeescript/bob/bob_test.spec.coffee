Bob = require "./bob"
describe "Bob", ->
  bob = new Bob()
  it "stating something", ->
    result = bob.hey "Tom-ay-to, tom-aaaah-to."
    expect(result).toEqual "Whatever."

  it "shouting", ->
    result = bob.hey "WATCH OUT!"
    expect(result).toEqual "Whoa, chill out!"

  it "asking a question", ->
    result = bob.hey "Does this cryogenic chamber make me look fat?"
    expect(result).toEqual "Sure."

  it "talking forcefully", ->
    result = bob.hey "Let's go make out behind the gym!"
    expect(result).toEqual "Whatever."

  it "using acronyms in regular speech", ->
    result = bob.hey "It's OK if you don't want to go to the DMV."
    expect(result).toEqual "Whatever."

  it "forceful questions", ->
    result = bob.hey "WHAT THE HELL WERE YOU THINKING?"
    expect(result).toEqual "Whoa, chill out!"

  it "shouting numbers", ->
    result = bob.hey "1, 2, 3 GO!"
    expect(result).toEqual "Whoa, chill out!"

  it "only number", ->
    result = bob.hey "1, 2, 3"
    expect(result).toEqual "Whatever."

  it "shouting with special characters", ->
    result = bob.hey "ZOMG THE %^*@#$(*^ ZOMBIES ARE COMING!!11!!1!"
    expect(result).toEqual "Whoa, chill out!"

  it "shouting with no exclamation mark", ->
    result = bob.hey "I HATE YOU"
    expect(result).toEqual "Whoa, chill out!"

  it "statement containing question mark", ->
    result = bob.hey "Ending with a ? means a question."
    expect(result).toEqual "Whatever."

  it "prattling on", ->
    result = bob.hey "Wait! Hang on.  Are you going to be OK?"
    expect(result).toEqual "Sure."

  it "silence", ->
    result = bob.hey ""
    expect(result).toEqual "Fine. Be that way!"

  it "prolonged silence", ->
    result = bob.hey "   "
    expect(result).toEqual "Fine. Be that way!"
