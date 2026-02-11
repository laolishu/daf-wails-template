import { GetConfigSummary } from "wailsjs/go/backend/App";
import type { ConfigSummary } from "../types/configSummary";

export async function fetchConfigSummary(): Promise<ConfigSummary> {
  return GetConfigSummary();
}
