class OCR

    RAW_NUMS = <<-RAW_NUMS.chomp
 _     _  _     _  _  _  _  _
| |  | _| _||_||_ |_   ||_||_|
|_|  ||_  _|  | _||_|  ||_| _|

    RAW_NUMS

    NUMS_ARY = RAW_NUMS.split("\n")
                       .map {|x| x.scan(/.{1,3}/)}
                       .map {|x| x.map {|y| y.rstrip}}
                       .transpose
    NUMS = NUMS_ARY.map {|x| x.join("\n") << "\n"}

  attr_reader :text
  def initialize(text)
    @text = text
  end

  def convert
    NUMS.index(text).to_s
  end

end
