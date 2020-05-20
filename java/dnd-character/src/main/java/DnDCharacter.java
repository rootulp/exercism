import java.util.Arrays;

class DnDCharacter {

    int ability() {
		int[] rolls = { roll(), roll(), roll(), roll() };
		Arrays.sort(rolls);
		int[] topThreeRolls = Arrays.copyOf(rolls, 3);
		return Arrays.stream(topThreeRolls).sum();
    }

    int modifier(int constitution) {
		return (int) Math.floor((constitution - 10) / 2.0);
    }

    int getStrength() {
        throw new UnsupportedOperationException("Delete this statement and write your own implementation.");
    }

    int getDexterity() {
        throw new UnsupportedOperationException("Delete this statement and write your own implementation.");
    }

    int getConstitution() {
        throw new UnsupportedOperationException("Delete this statement and write your own implementation.");
    }

    int getIntelligence() {
        throw new UnsupportedOperationException("Delete this statement and write your own implementation.");
    }

    int getWisdom() {
        throw new UnsupportedOperationException("Delete this statement and write your own implementation.");
    }

    int getCharisma() {
        throw new UnsupportedOperationException("Delete this statement and write your own implementation.");
    }

    int getHitpoints() {
        throw new UnsupportedOperationException("Delete this statement and write your own implementation.");
	}

	/**
	 *
	 * @return a random number between 1 and 6 inclusive.
	 */
	private int roll() {
		return (int) Math.floor(Math.random() * 6);
	}

}
