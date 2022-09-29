# Rna Transcription

Welcome to Rna Transcription on Exercism's Rust Track.
If you need help running the tests or submitting your code, check out `HELP.md`.

## Instructions

Given a DNA strand, return its RNA complement (per RNA transcription).

Both DNA and RNA strands are a sequence of nucleotides.

The four nucleotides found in DNA are adenine (**A**), cytosine (**C**),
guanine (**G**) and thymine (**T**).

The four nucleotides found in RNA are adenine (**A**), cytosine (**C**),
guanine (**G**) and uracil (**U**).

Given a DNA strand, its transcribed RNA strand is formed by replacing
each nucleotide with its complement:

* `G` -> `C`
* `C` -> `G`
* `T` -> `A`
* `A` -> `U`

By using private fields in structs with public `new` functions returning
`Option` or `Result` (as here with `DNA::new` & `RNA::new`), we can guarantee
that the internal representation of `DNA` is correct. Because every valid DNA
string has a valid RNA string, we don't need to return a `Result`/`Option` from
`into_rna`.

This explains the type signatures you will see in the tests.

## Source

### Created by

- @EduardoBautista

### Contributed to by

- @andrewclarkson
- @AndrewKvalheim
- @ashleygwilliams
- @benreyn
- @coriolinus
- @cwhakes
- @EduardoBautista
- @efx
- @ErikSchierboom
- @etrepum
- @IanWhitney
- @kytrinyx
- @lutostag
- @mkantor
- @nfiles
- @petertseng
- @pminten
- @rofrol
- @rpottsoh
- @samcday
- @shingtaklam1324
- @stringparser
- @TheDarkula
- @xakon
- @ZapAnton

### Based on

Hyperphysics - http://hyperphysics.phy-astr.gsu.edu/hbase/Organic/transcription.html