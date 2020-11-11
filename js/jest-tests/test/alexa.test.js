// helloAlexa returns good morning
// helloAlexa, on Monday returns good morning, have a great start to the week
// helloAlexa, on Friday returns good morning, hope you have great plan for the weekend

function helloAlexa(output) {
  output.say("Good Morning");
  output.say("have a great start to the week");
}

describe("Alexa", () => {
  it("says good morning", () => {

    const output = { say:  jest.fn() };
    helloAlexa(output);
    expect(output.say).toHaveBeenLastCalledWith("Good Morning");
  });

  it("on monday it says good morning, have a great start to the week", () => { 
    const output = { say:  jest.fn() };
    helloAlexa(output);
    expect(output.say).toBeCalledWith("Good Morning");
    expect(output.say).toBeCalledWith("have a great start to the week");
  });

  //toHaveBeenLastCalledWithexpect(drink).toHaveBeenNthCalledWith(1, 'lemon');  expect(drink).toHaveBeenNthCalledWith(2, 'octopus');
});
