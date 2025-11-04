import { dev, building } from "$app/environment";

if (!building) {
  if (dev) {
    const { worker } = await import("./mocks/browser");

    await worker.start({ onUnhandledRequest: "bypass" });
  }
}
