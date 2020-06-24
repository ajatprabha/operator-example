import MockAdapter from "axios-mock-adapter";

export default class Mock {
  constructor(axios) {
    const mockDelay = 0;

    this.mock = new MockAdapter(axios, { delayResponse: mockDelay });
    this.mock.injectMocks = () => {
      return this.mock;
    };
  }

  setupPluginMocks() {
    this.mock
      .injectMocks() // additional mocks added from RestClient
      .onAny()
      .passThrough();
  }

  setupMockEndpoints() {
    console.warn(
      "%c ✨You are mocking api requests.",
      "background: gray; color: white; display: block; padding: 0.25rem;"
    );

    this.mock
      // .onGet("/api/darkrooms")
      // .reply(200, {
      //   items: [
      //     {
      //       name: "darkroom-sample",
      //       version: "0.1.0",
      //       source: {
      //         type: "WebFolder"
      //       },
      //       domains: [
      //         "students.darkroom.gojek.io"
      //       ]
      //     }
      //   ]
      // })
      .onAny()
      .passThrough();
  }
}
