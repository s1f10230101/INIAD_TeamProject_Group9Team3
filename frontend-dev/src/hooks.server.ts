import { building, browser, dev } from "$app/environment";

if (!building) {
  if (!browser && dev) {
    // Node version 25で既知の回避不能エラーが出るのでバージョン注意
    const { server } = await import("./mocks/server");

    server.listen({ onUnhandledRequest: "warn" });
    console.log("[MSW] server mocking enabled");
  }
}
