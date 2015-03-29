class Garden

  DEFAULT_STUDENTS = %w(Alice Bob Charlie David Eve Fred Ginny
                        Harriet Ileana Joseph Kincaid Larry)

  PLANTS = { 'G' => :grass,
             'C' => :clover,
             'R' => :radishes,
             'V' => :violets }

  attr_reader :diagram, :diagram_clean, :students
  def initialize(diagram, students = DEFAULT_STUDENTS)
    @diagram = diagram
    @diagram_clean = diagram.split("\n").map {|x| x.split(//)}
    @students = students.sort
    setup_student_methods
  end

  def setup_student_methods
    students.each_with_index do |student, index|
      instance_eval "def #{student.downcase}; #{plants_for(index)}; end"
    end
  end
  
  def plants_for(index)
    symbols_for(index).map { |x| PLANTS[x] }
  end

  def symbols_for(index)
    diagram_clean[0].values_at(index * 2, index * 2 + 1) +
    diagram_clean[1].values_at(index * 2, index * 2 + 1)
  end

end
