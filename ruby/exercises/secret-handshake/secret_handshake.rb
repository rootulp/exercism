# SecretHandshake
class SecretHandshake
  COMMANDS = ['wink', 'double blink', 'close your eyes', 'jump'].freeze
  attr_reader :decimal
  def initialize(decimal)
    @decimal = decimal
  end

  def commands
    return [] unless decimal.is_a?(Integer)
    binary.reverse[4] == '1' ? handshake.reverse : handshake
  end

  def binary
    format('%05d', decimal.to_s(2))
  end

  private

  def handshake
    COMMANDS.select.with_index do |_, index|
      binary.reverse[index] == '1'
    end
  end
end
