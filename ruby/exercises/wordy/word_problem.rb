# Word Problem
class WordProblem
  OPERATORS = {
    'plus' => '+',
    'minus' => '-',
    'multiplied' => '*',
    'divided' => '/'
  }.freeze
  attr_reader :question
  def initialize(question)
    @question = question.chomp('?')
  end

  def answer
    fail ArgumentError if empty_stacks?
    calculate
  end

  private

  def calculate
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

  def evaluate(op, num1, num2)
    num1.send(op, num2)
  end

  def empty_stacks?
    operator_stack.empty? || number_stack.empty?
  end

  def words
    question.split(' ')
  end

  def operator_stack
    words.keep_if { |word| operator?(word) }.map { |word| to_op(word) }
  end

  def number_stack
    words.keep_if { |word| number?(word) }.map(&:to_i)
  end

  def operator?(word)
    OPERATORS.include?(word)
  end

  def to_op(op)
    OPERATORS[op]
  end

  def number?(word)
    Integer(word)
  rescue
    false
  end
end
