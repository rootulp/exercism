import java.util.Arrays;

class DnDCharacter {

	public static int generateAbility() {
		int[] rolls = { roll(), roll(), roll(), roll() };
		Arrays.sort(rolls);
		int[] topThreeRolls = Arrays.copyOf(rolls, 3);
		return Arrays.stream(topThreeRolls).sum();
	}

    public static int calculateModifier(int constitution) {
		return (int) Math.floor((constitution - 10) / 2.0);
    }

	/**
	 *
	 * @return a random number between 1 and 6 inclusive.
	 */
	private static int roll() {
		return (int) Math.floor(Math.random() * 6) ;
	}

	private final int strength;
	private final int dexterity;
	private final int constituion;
	private final int intelligence;
	private final int wisdom;
	private final int charisma;
	private final int hitpoints;

	public DnDCharacter() {
		this.strength = generateAbility();
		this.dexterity = generateAbility();
		this.constituion = generateAbility();
		this.intelligence = generateAbility();
		this.wisdom = generateAbility();
		this.charisma = generateAbility();
		this.hitpoints = 10 + this.modifier(this.constituion);
	}

	/**
	 * This method is strictly exported because unit tests for this class expect an instance method.
	 * Please use the static method `generateAbility` instead of this method.
	 */
    int ability() {
		return generateAbility();
    }

	/**
	 * This method is strictly exported because unit tests for this class expect an instance method.
	 * Please use the static method `calculateModifier` instead of this method.
	 */
    int modifier(int constitution) {
		return calculateModifier(constitution);
    }

    int getStrength() {
		return this.strength;
    }

    int getDexterity() {
		return this.dexterity;
    }

    int getConstitution() {
		return this.constituion;
    }

    int getIntelligence() {
		return this.intelligence;
    }

    int getWisdom() {
		return this.wisdom;
    }

    int getCharisma() {
		return this.charisma;
    }

    int getHitpoints() {
		return this.hitpoints;
	}

}
