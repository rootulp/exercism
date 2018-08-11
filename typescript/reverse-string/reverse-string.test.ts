import ReverseString from './reverse-string'

describe('Reverse String', () => {
    it('an empty string', () => {
        const expected = ''
        expect(ReverseString.reverse('')).toEqual(expected)
    })

    it('a word', () => {
        const expected = 'tobor'
        expect(ReverseString.reverse('robot')).toEqual(expected)
    })

    it('a capitalized word', () => {
        const expected = 'nemaR'
        expect(ReverseString.reverse('Ramen')).toEqual(expected)
    })

    it('a sentence with punctuation', () => {
        const expected = `!yrgnuh m'I`
        expect(ReverseString.reverse(`I'm hungry!`)).toEqual(expected)
    })

    it('a palindrome', () => {
        const expected = 'racecar'
        expect(ReverseString.reverse('racecar')).toEqual(expected)
    })
})
