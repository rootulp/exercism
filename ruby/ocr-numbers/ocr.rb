# OCR
class OCR
  RAW_NUMS = <<-RAW_NUMS.chomp
 _     _  _     _  _  _  _  _
| |  | _| _||_||_ |_   ||_||_|
|_|  ||_  _|  | _||_|  ||_| _|

  RAW_NUMS
  NUMS = split_blocks(RAW_NUMS)

  attr_reader :input
  def initialize(input)
    @input = input
  end

  def convert
    split_chunks(input).map do |blocks|
      blocks.map do |block|
        convert_block(block)
      end.join(',')
    end.join.chomp(',')
  end

  private

  def convert_block(block)
    NUMS.include?(block) ? NUMS.index(block).to_s : '?'
  end

  def split_chunks(chunks)
    chunks.split("\n\n").map { |blocks| split_blocks(blocks) }
  end

  def split_blocks(blocks)
    blocks.split("\n")
          .map { |x| x.scan(/.{1,3}/) }
          .map { |x| x.map(&:rstrip) }
          .transpose
          .map { |x| x.join("\n") << "\n" }
  end
end
