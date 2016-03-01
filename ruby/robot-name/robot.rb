# Robot
class Robot
  attr_reader :name
  def initialize
    reset
  end

  def reset
    @name = generate_name
  end

  private

  def generate_name
    prefix + suffix
  end

  def prefix
    (0...2).map { ('A'..'Z').to_a[rand(26)] }.join
  end

  def suffix
    (0...3).map { ('1'..'9').to_a[rand(9)] }.join
  end
end
