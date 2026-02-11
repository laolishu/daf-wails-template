import { GetSystemInfo } from "wailsjs/go/backend/App";
import type { SystemInfo } from "../types/systemInfo";

export async function fetchSystemInfo(): Promise<SystemInfo> {
  return GetSystemInfo();
}
