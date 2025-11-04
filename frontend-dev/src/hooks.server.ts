import { dev, building, browser } from "$app/environment";

if (!building) {
  if (!browser && dev && import.meta.env.DEV) {
    // Node version 25で既知の回避不能エラーが出るのでバージョン注意
    const { server } = await import("./mocks/server");

    server.listen({ onUnhandledRequest: "warn" });
    console.log("[MSW] server mocking enabled");
  }
}
