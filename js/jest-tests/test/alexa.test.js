// helloAlexa returns good morning
// helloAlexa, on Monday returns good morning, have a great start to the week
// helloAlexa, on Friday returns good morning, hope you have great plan for the weekend

function helloAlexa(output) {
  output.say("Good morning");
}

describe("Alexa", () => {
  it("says good morning", () => {
    helloAlexa({ say: (text) => console.log(text) });
  });
});
