interface IRunLengthEncoding {
    encode(data: string): string
    decode(encodedData: string): string

}

export default class RunLengthEncoding implements IRunLengthEncoding {
    encode(data: string): string {
        throw new Error("Method not implemented.")
    }

    decode(encodedData: string): string {
        throw new Error("Method not implemented.")
    }

}
