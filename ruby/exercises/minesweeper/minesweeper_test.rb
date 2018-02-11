require_relative 'minesweeper'

class MinesweeperTest < Minitest::Test
  def test_transform1
    inp = ['+------+', '| *  * |', '|  *   |', '|    * |', '|   * *|',
           '| *  * |', '|      |', '+------+']
    out = ['+------+', '|1*22*1|', '|12*322|', '| 123*2|', '|112*4*|',
           '|1*22*2|', '|111111|', '+------+']
    assert_equal out, Board.transform(inp)
  end

  def test_transform2
    inp = ['+-----+', '| * * |', '|     |', '|   * |', '|  * *|',
           '| * * |', '+-----+']
    out = ['+-----+', '|1*2*1|', '|11322|', '| 12*2|', '|12*4*|',
           '|1*3*2|', '+-----+']
    assert_equal out, Board.transform(inp)
  end

  def test_transform3
    inp = ['+-----+', '| * * |', '+-----+']
    out = ['+-----+', '|1*2*1|', '+-----+']
    assert_equal out, Board.transform(inp)
  end

  def test_transform4
    inp = ['+-+', '|*|', '| |', '|*|', '| |', '| |', '+-+']
    out = ['+-+', '|*|', '|2|', '|*|', '|1|', '| |', '+-+']
    assert_equal out, Board.transform(inp)
  end

  def test_transform5
    inp = ['+-+', '|*|', '+-+']
    out = ['+-+', '|*|', '+-+']
    assert_equal out, Board.transform(inp)
  end

  def test_transform6
    inp = ['+--+', '|**|', '|**|', '+--+']
    out = ['+--+', '|**|', '|**|', '+--+']
    assert_equal out, Board.transform(inp)
  end

  def test_transform7
    inp = ['+--+', '|**|', '|**|', '+--+']
    out = ['+--+', '|**|', '|**|', '+--+']
    assert_equal out, Board.transform(inp)
  end

  def test_transform8
    inp = ['+---+', '|***|', '|* *|', '|***|', '+---+']
    out = ['+---+', '|***|', '|*8*|', '|***|', '+---+']
    assert_equal out, Board.transform(inp)
  end

  def test_transform9
    inp = ['+-----+', '|     |', '|   * |', '|     |', '|     |',
           '| *   |', '+-----+']
    out = ['+-----+', '|  111|', '|  1*1|', '|  111|', '|111  |',
           '|1*1  |', '+-----+']
    assert_equal out, Board.transform(inp)
  end

  def test_different_len
    inp = ['+-+', '| |', '|*  |', '|  |', '+-+']
    assert_raises(ValueError) do
      Board.transform(inp)
    end
  end

  def test_faulty_border
    inp = ['+-----+', '*   * |', '+-- --+']
    assert_raises(ValueError) do
      Board.transform(inp)
    end
  end

  def test_invalid_char
    inp = ['+-----+', '|X  * |', '+-----+']
    assert_raises(ValueError) do
      Board.transform(inp)
    end
  end
end
