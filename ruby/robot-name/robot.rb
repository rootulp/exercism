class Robot
  attr_reader :name

  def initialize
    @name = generate_name
  end

  def reset
    @name = generate_name
  end

  def generate_name
    prefix = (0...2).map { ('A'..'Z').to_a[rand(26)] }.join
    suffix = (0...3).map { ('1'..'9').to_a[rand(9)] }.join
    prefix + suffix
  end

end