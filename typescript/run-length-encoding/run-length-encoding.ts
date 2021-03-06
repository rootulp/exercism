import {last, isUndefined, startsWith, concat, slice, repeat} from "lodash"

export default class RunLengthEncoding {
    static encode(data: string): string {
        const chunks: string[] = RunLengthEncoding.splitIntoChunks(data)
        return chunks.map(RunLengthEncoding.encodeChunk).join('')
    }

    static decode(encodedData: string): string {
        const encodedChunks: string[] = RunLengthEncoding.splitIntoEncodedChunks(encodedData)
        return encodedChunks.map(RunLengthEncoding.decodeChunk).join('')
    }

    static splitIntoChunks(data: string): string[] {
        return data.split("").reduce((chunks: string[], currentCharacter: string) => {
            const lastChunkOfConsecutiveCharacters: string | undefined = last(chunks)

            if (!isUndefined(lastChunkOfConsecutiveCharacters) && startsWith(lastChunkOfConsecutiveCharacters, currentCharacter)) {
                return concat(slice(chunks, 0, -1), last(chunks) + currentCharacter)
            } else {
                return concat(chunks, currentCharacter)
            }
        }, [])
    }

    static splitIntoEncodedChunks(encodedData: string): string[] {
        const digitsFollowedByData = new RegExp(/(\d*\D{1})/g)
        const encodedChunk = encodedData.match(digitsFollowedByData)
        return encodedChunk ? encodedChunk : []
    }

    static decodeChunk(encodedChunk: string): string {
        if (encodedChunk.length === 1) {
            return encodedChunk
        } else {
            const numOccurences = parseInt(encodedChunk.slice(0, -1), 10)
            const data = encodedChunk.slice(-1)
            return repeat(data, numOccurences)
        }
    }

    static encodeChunk(chunk: string): string {
        if (chunk.length === 1) {
            return chunk
        } else {
            return `${chunk.length}${chunk.charAt(0)}`
        }
    }
}
