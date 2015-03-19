class OCR

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
    result = String.new

    split_chunks(input).each do |blocks|
      blocks.each do |block|
        result << convert_block(block)
      end
      result << ","
    end
    result.chomp(",")
  end

  private

  def convert_block(block)
    if nums.include?(block)
      nums.index(block).to_s
    else
      "?"
    end
  end

  def split_chunks(chunks)
    chunks.split("\n\n").map do |blocks|
      split_blocks(blocks)
    end
  end

  def split_blocks(blocks)
    blocks.split("\n")
          .map {|x| x.scan(/.{1,3}/)}
          .map {|x| x.map {|y| y.rstrip}}
          .transpose
          .map {|x| x.join("\n") << "\n"}
  end
end
