import org.junit.jupiter.api.Test;

import java.io.PrintStream;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertTrue;
import static org.mockito.Mockito.mock;
import static org.mockito.Mockito.verify;

public class ThingTest {


    /**
     * 1. If we put echo it output ohce
     * 2. If we input a palendrome it reverse and outputs ¡Bonita palabra! ana for example ,
     * input ana would produce ¡Bonita palabra!
     * 3. If we input stop , ohce will produce *Adios < your name >*
     * 4. If given an input with your name , it should greet with your name
     */


    @Test
    void itOutputsInputInReverse() {
        OcheReader oche = new OcheReader("echo", System.out);

        assertEquals("ohce",
                oche.getResponse());

        final PrintStream mock = mock(PrintStream.class);

        oche = new OcheReader("tom", mock);

        assertEquals("mot",
                oche.getResponse());

        verify(mock).println("mot");
    }

    @Test
    void itGreetsWithYourName() {
        final PrintStream mock = mock(PrintStream.class);

        oche = new OcheReader("Tom", mock);
        verify(mock).println("¡Buenas noches Tom");
    }

    @Test
    void it_should_not_fail() {
        assertTrue(true);
    }
}
