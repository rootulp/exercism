# SecretHandshake
class SecretHandshake
  attr_reader :decimal
  def initialize(decimal)
    @decimal = decimal
  end

  def commands
    return [] unless decimal.is_a?(Integer)
    handshake
  end

  def binary
    format('%05d', decimal.to_s(2))
  end

  private

  def handshake
    results = []
    results << 'wink'            if binary.reverse[0] == '1'
    results << 'double blink'    if binary.reverse[1] == '1'
    results << 'close your eyes' if binary.reverse[2] == '1'
    results << 'jump'            if binary.reverse[3] == '1'
    results.reverse!             if binary.reverse[4] == '1'
    results
  end
end
