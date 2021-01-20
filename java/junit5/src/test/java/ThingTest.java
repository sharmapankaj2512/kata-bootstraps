import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

public class ThingTest {

    @Test
    void fail() {
        Thing thing = new Thing();
        String value = thing.callForAction();
        assertEquals("Food", value);
    }

    /*
    1. If we put echo it output ohce
    2. If we input a palendrome it reverse and outputs ¡Bonita palabra! ana for example ,
       input ana would produce ¡Bonita palabra!
    3. If we input stop , ohce will produce *Adios < your name >*
    4. If given an input with your name , it should greet with your name

     */


    @Test
    void should_greet_with_name(){

    }

    @Test
    void it_should_not_fail() {
        assertTrue(true);
    }
}
