class Garden

  STUDENTS = %w(Alice Bob Charlie David Eve Fred Ginny
                Harriet Ileana Joseph Kincaid Larry)

  PLANTS = { 'G' => :grass,
             'C' => :clover,
             'R' => :radishes,
             'V' => :violets }

  attr_reader :diagram, :students, :plant_ary
  def initialize(diagram, students = STUDENTS)
    @diagram = diagram
    @plant_ary = diagram.split("\n").map {|x| x.split(//)}
    @students = students.sort
    setup_student_methods
  end

  def setup_student_methods
    students.each_with_index do |student, index|
      self.class.send(:define_method, student.downcase) do
        plants_for(index).map { |x| PLANTS[x] }
      end
    end
  end

  def plants_for(index)
    [ plant_ary.first[index * 2],
      plant_ary.first[index * 2 + 1],
      plant_ary.last[index * 2],
      plant_ary.last[index * 2 + 1] ]
  end

end
