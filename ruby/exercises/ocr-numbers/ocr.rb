# frozen_string_literal: true

# OCR
class OCR
  GARBLE = '?'.freeze
  RAW_NUMS = <<-RAW_NUMS.chomp
 _     _  _     _  _  _  _  _
| |  | _| _||_||_ |_   ||_||_|
|_|  ||_  _|  | _||_|  ||_| _|

  RAW_NUMS

  attr_reader :input, :nums
  def initialize(input)
    @input = input
    @nums = split_blocks(RAW_NUMS)
  end

  def convert
    split_chunks(input).map do |chunk|
      split_blocks(chunk).map do |block|
        translate_block(block)
      end.join
    end.join(',')
  end

  private

  def translate_block(block)
    nums.include?(block) ? nums.index(block).to_s : GARBLE
  end

  def split_chunks(chunks)
    chunks.split("\n\n")
  end

  # split_blocks is too complex - consider splitting this up
  def split_blocks(blocks)
    blocks.split("\n")
          .map { |block| block.scan(/.{1,3}/) }
          .map { |row| row.map { |val| val.ljust(3) } }
          .map { |row| row.empty? ? ['   '] : row } # for 1 and 4
          .transpose
          .map { |x| x.join("\n") }
  end
end
