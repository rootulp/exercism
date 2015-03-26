class WordProblem

  OPERATORS = {
    'plus' => '+',
    'minus' => '-',
    'multiplied' => '*',
    'divided' => '/'
  }

  attr_reader :question
  def initialize(question)
    @question = question.chomp('?')
  end

  def answer
    raise ArgumentError if operator_stack.empty? || number_stack.empty?
    operators = operator_stack
    numbers = number_stack

    while operators.any?
      op = operators.shift
      num1, num2 = numbers.shift(2)
      result = evaluate(op, num1, num2)
      numbers.unshift(result)
    end

    numbers.pop
  end

  private

  def words
    question.split(' ')
  end

  def is_operator?(word)
    OPERATORS.include?(word)
  end

  def is_number?(word)
    true if Float(word) rescue false # taken from StackOverflow
  end

  def operator_stack
    words.keep_if { |word| is_operator?(word) }
  end

  def number_stack
    words.keep_if { |word| is_number?(word) }.map(&:to_i)
  end

  def evaluate(op, num1, num2)
    num1.send(OPERATORS[op], num2)
  end

end

