import { setupWorker } from "msw/browser";
import { handlers } from "./handlers";

// ブラウザ環境で動作するService Workerを設定します
export const worker = setupWorker(...handlers);
