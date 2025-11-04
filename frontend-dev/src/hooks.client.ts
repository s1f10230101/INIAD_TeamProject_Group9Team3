import { building, dev } from "$app/environment";

if (!building && dev) {
  const { worker } = await import("./mocks/browser");

  await worker.start({ onUnhandledRequest: "bypass" });
}
