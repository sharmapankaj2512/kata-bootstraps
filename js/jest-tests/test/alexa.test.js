// helloAlexa returns good morning
// helloAlexa, on Monday returns good morning, have a great start to the week
// helloAlexa, on Friday returns good morning, hope you have great plan for the weekend

function helloAlexa(output) {
  output.say("Good Morning");
}

describe("Alexa", () => {
  it("says good morning", () => {

    const output = { say:  jest.fn() };
    helloAlexa(output);
    //expect(output.say).toBeCalled();
    expect(output.say).toBeCalledWith("Good Morning");

  });
});
