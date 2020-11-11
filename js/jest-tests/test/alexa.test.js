// helloAlexa returns good morning
// helloAlexa, on Monday returns good morning, have a great start to the week
// helloAlexa, on Friday returns good morning, hope you have great plan for the weekend

function helloAlexa(output, day) {
  output.say("Good Morning");
  if (day == "Mon") output.say("have a great start to the week");
  if (day == "Fri") output.say("hope you have great plan for the weekend");
}

describe("Alexa", () => {
  it("says good morning", () => {
    const output = { say: jest.fn() };
    helloAlexa(output, "Sun");
    expect(output.say).toHaveBeenLastCalledWith("Good Morning");
  });

  it("on monday it says good morning, have a great start to the week", () => {
    const output = { say: jest.fn() };
    const day = { say: jest.fn() };

    helloAlexa(output, "Mon");
    expect(output.say).toBeCalledWith("Good Morning");
    expect(output.say).toHaveBeenLastCalledWith(
      "have a great start to the week"
    );
  });

  it("on friday it says good morning, hope you have great plan for the weekend", () => {
    const output = { say: jest.fn() };

    helloAlexa(output, "Fri");
    expect(output.say).toBeCalledWith("Good Morning");
    expect(output.say).toHaveBeenLastCalledWith(
      "hope you have great plan for the weekend"
    );
  });

  //toHaveBeenLastCalledWithexpect(drink).toHaveBeenNthCalledWith(1, 'lemon');  expect(drink).toHaveBeenNthCalledWith(2, 'octopus');
});
