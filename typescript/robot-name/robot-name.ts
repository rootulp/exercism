export default class Robot {
  private _name: string;
  private static existingNames = new Set();
  constructor() {
    this._name = this.generateName();
    Robot.existingNames.add(this._name);
  }

  public get name(): string {
    return this._name;
  }

  public resetName(): void {
    const name = this.generateName();
    if (Robot.existingNames.has(name)) {
      this.resetName();
    } else {
      this._name = name;
      Robot.existingNames.add(this._name);
    }
  }

  public static releaseNames(): void {
    Robot.existingNames = new Set();
  }

  private generateName(): string {
    return [
      this.randomLetter(),
      this.randomLetter(),
      this.randomNumber(),
      this.randomNumber(),
      this.randomNumber(),
    ].join("")
  }

  private randomLetter(): string {
    const ALPHABET = "ABCDEFGHIJKLMNOPQRSTUVWXYZ";
    return ALPHABET.charAt(Math.floor(Math.random() * 26));
  }

  private randomNumber(): string {
    return Math.floor(Math.random() * 10).toString();
  }
}
