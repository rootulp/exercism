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
    result = ""
    split_blocks(input).each do |block|
      if nums.include?(block)
        result << nums.index(block).to_s
      else
        result << "?"
      end
    end
    result
  end

  def split_blocks(text)
    text.split("\n")
        .map {|x| x.scan(/.{1,3}/)}
        .map {|x| x.map {|y| y.rstrip}}
        .transpose
        .map {|x| x.join("\n") << "\n"}
  end
end
