class SecretHandshake

  attr_reader :commands
  def initialize(decimal)
    if decimal.is_a?(Integer)
      binary = sprintf('%05d', decimal.to_s(2))
      @commands = build_handshake(binary)
    else
      @commands = []
    end
  end

  private

  def build_handshake(binary)
    results = []
    results << 'wink'            if binary.reverse[0] == "1"
    results << 'double blink'    if binary.reverse[1] == "1"
    results << 'close your eyes' if binary.reverse[2] == "1"
    results << 'jump'            if binary.reverse[3] == "1"
    results.reverse!             if binary.reverse[4] == "1"
    results
  end
end
