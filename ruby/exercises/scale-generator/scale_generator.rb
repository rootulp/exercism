# Scale
class Scale
  ASCENDING_INTERVALS = %w[m M A].freeze
  CHROMATIC_SCALE = %w[C C# D D# E F F# G G# A A# B].freeze
  FLAT_CHROMATIC_SCALE = %w[C Db D Eb E F Gb G Ab A Bb B].freeze
  FLAT_KEYS = %w[F Bb Eb Ab Db Gb d g c f bb eb].freeze

  attr_reader :tonic, :scale_name, :pattern
  def initialize(tonic, scale_name, pattern = false)
    @tonic = tonic
    @scale_name = scale_name
    @pattern = pattern
  end

  def name
    "#{tonic.capitalize} #{scale_name}"
  end

  def pitches
    return reorder_chromatic_scale unless pattern
    last_index = 0
    pattern.each_char.with_object([]) do |char, arr|
      arr << reorder_chromatic_scale[last_index]
      last_index += ASCENDING_INTERVALS.index(char) + 1
    end
  end

  def reorder_chromatic_scale
    return chromatic_scale if tonic.capitalize == 'C'
    index = chromatic_scale.index(tonic.capitalize)
    chromatic_scale[index..-1] + chromatic_scale[0..index - 1]
  end

  def chromatic_scale
    FLAT_KEYS.include?(tonic) ? FLAT_CHROMATIC_SCALE : CHROMATIC_SCALE
  end
end
