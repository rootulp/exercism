create or replace package ut_hamming#
is
  procedure run;
end ut_hamming#;
/

create or replace package body ut_hamming#
is

  procedure test (
    i_descn                                       varchar2
   ,i_exp                                         pls_integer
   ,i_act                                         pls_integer
  )
  is
  begin
    if i_exp = i_act then
      dbms_output.put_line('SUCCESS: ' || i_descn);
    else
      dbms_output.put_line(
           'FAILURE: '   || i_descn
        || ': expected ' || nvl('' || i_exp, 'null')
        || ', but got '  || nvl('' || i_act, 'null')
        || '!'
      );
    end if;
  end test;

  procedure run
  is
  begin
    test('test_no_difference_between_identical_strands'                    , 0, hamming#.distance(i_first => 'A'           , i_second => 'A'           ));
    --test('test_complete_hamming_distance_of_for_single_nucleotide_strand', 1, hamming#.distance(i_first => 'A'           , i_second => 'G'           ));
    --test('test_complete_hamming_distance_of_for_small_strand'            , 2, hamming#.distance(i_first => 'AG'          , i_second => 'CT'          ));
    --test('test_small_hamming_distance'                                   , 1, hamming#.distance(i_first => 'AG'          , i_second => 'AT'          ));
    --test('test_small_hamming_distance_in_longer_strand'                  , 1, hamming#.distance(i_first => 'GGACG'       , i_second => 'GGTCG'       ));
    --test('test_nonunique_characters_within_first_strand'                 , 1, hamming#.distance(i_first => 'AGA'         , i_second => 'AGG'         ));
    --test('test_nonunique_characters_within_second_strand'                , 1, hamming#.distance(i_first => 'AGG'         , i_second => 'AGA'         ));
    --test('test_large_hamming_distance'                                   , 4, hamming#.distance(i_first => 'GATACA'      , i_second => 'GCATAA'      ));
    --test('test_hamming_distance_in_very_long_strand'                     , 9, hamming#.distance(i_first => 'GGACGGATTCTG', i_second => 'AGGACGGATTCT'));
  exception
    when others then
      dbms_output.put_line('Test execution failed.');
      dbms_output.put_line(sqlerrm);
  end run;

end ut_hamming#;
/

begin
  ut_hamming#.run;
end;
/
