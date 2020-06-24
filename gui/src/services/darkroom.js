import RestClient from "./restClient";

export default class Darkroom {
  constructor(options) {
    const opts = options || {};

    this.options = opts;
    this.client = new RestClient(opts);
  }

  buildUrl(path) {
    return this.client.buildUrl(path);
  }

  /**
   * Info
   */
  getInfo() {
    return this.client.get("/");
  }

  getStatus() {
    return this.client.status("/");
  }

  /**
   * Darkrooms
   */
  // get a list of all darkrooms
  getAllDarkrooms(params) {
    return this.client.get("/api/darkrooms", params);
  }
  createDarkroom(params) {
    return this.client.post("/api/darkrooms", params);
  }
}
