import {last, isUndefined, startsWith, concat, slice} from "lodash"

export default class RunLengthEncoding {
    static encode(data: string): string {
        const chunks = RunLengthEncoding.splitIntoConsecutiveCharacters(data)
        return chunks.map((chunk) => RunLengthEncoding.encodeChunk(chunk)).join('')
    }

    static decode(encodedData: string): string {
        return encodedData
    }

    static splitIntoConsecutiveCharacters(data: string): string[] {
        return data.split("").reduce((chunks: string[], currentCharacter: string) => {
            const lastChunkOfConsecutiveCharacters: string | undefined = last(chunks)

            if (!isUndefined(lastChunkOfConsecutiveCharacters) && startsWith(lastChunkOfConsecutiveCharacters, currentCharacter)) {
                return concat(slice(chunks, 0, -1), last(chunks) + currentCharacter)
            } else {
                return concat(chunks, currentCharacter)
            }
        }, [])
    }

    static encodeChunk(chunk: string): string {
        if (chunk.length === 1) {
            return chunk
        } else {
            return `${chunk.length}${chunk.charAt(0)}`
        }

    }
}
