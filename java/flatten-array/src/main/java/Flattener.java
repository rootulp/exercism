import java.util.Collection;
import java.util.List;
import java.util.Objects;
import java.util.stream.Collectors;
import java.util.stream.Stream;

public class Flattener {

    public List<?> flatten(List<?> nestedList) {
        return flattenStream(nestedList.stream())
                .collect(Collectors.toList());
    }

    private Stream<?> flattenStream(Stream<?> stream) {
        return stream.filter(Objects::nonNull)
                .flatMap(item -> item instanceof Collection ?
                    flattenStream(((Collection<?>) item).stream()) :
                    Stream.of(item)
                );
    }
}
